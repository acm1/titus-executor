package docker

import (
	"archive/tar"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Netflix/titus-executor/config"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const sshdConfig = `
# See the sshd_config(5) manpage for details

# What ports, IPs and protocols we listen for
Port 7522
# Use these options to restrict which interfaces/protocols sshd will bind to
#ListenAddress ::
#ListenAddress 0.0.0.0
Protocol 2
# HostKeys for protocol version 2
HostKey /titus/sshd/etc/ssh/ssh_host_rsa_key
HostKey /titus/sshd/etc/ssh/ssh_host_ecdsa_key
HostKey /titus/sshd/etc/ssh/ssh_host_ed25519_key

Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-512,hmac-sha2-256,umac-128@openssh.com

# Logging
SyslogFacility AUTH
LogLevel INFO

# Authentication:
LoginGraceTime 120
PermitRootLogin without-password
StrictModes yes

PubkeyAuthentication yes
AuthorizedKeysFile	/titus/ssh_keys/%u

TrustedUserCAKeys /titus/etc/ssh/trusted_user_ca_keys.pub
AuthorizedPrincipalsFile /titus/etc/ssh/authorized_principals_%u

# Don't read the user's ~/.rhosts and ~/.shosts files
IgnoreRhosts yes
# similar for protocol version 2
HostbasedAuthentication no
# Uncomment if you don't trust ~/.ssh/known_hosts for RhostsRSAAuthentication
#IgnoreUserKnownHosts yes

# To enable empty passwords, change to yes (NOT RECOMMENDED)
PermitEmptyPasswords no

# Change to yes to enable challenge-response passwords (beware issues with
# some PAM modules and threads)
ChallengeResponseAuthentication no

# Change to no to disable tunnelled clear text passwords
PasswordAuthentication no

# Kerberos options
#KerberosAuthentication no
#KerberosGetAFSToken no
#KerberosOrLocalPasswd yes
#KerberosTicketCleanup yes

# GSSAPI options
#GSSAPIAuthentication no
#GSSAPICleanupCredentials yes

X11Forwarding no
X11DisplayOffset 10
PrintMotd no
PrintLastLog no
TCPKeepAlive yes
#UseLogin no

#MaxStartups 10:30:60
#Banner /etc/issue.net

# Allow client to pass locale environment variables
AcceptEnv LANG LC_*

Subsystem sftp /titus/sshd/usr/lib64/misc/sftp-server

PidFile /run/sshd.pid
`

func AddContainerSSHDConfig(taskID string, appName string, iamProfileARN string, tw *tar.Writer, cfg config.Config) error {
	if iamProfileARN == "" {
		return errors.New("Could not get IAM role from container")
	}

	iamProfile, err := arn.Parse(iamProfileARN)
	if err != nil {
		return err
	}

	caData, err := ioutil.ReadFile(cfg.ContainerSSHDCAFile)
	if err != nil {
		return err
	}
	logrus.Info("Going to make sshdconfig with data")
	return addContainerSSHDConfigWithData(taskID, appName, tw, cfg, caData, iamProfile.AccountID)
}

func addContainerSSHDConfigWithData(taskID string, appName string, tw *tar.Writer, cfg config.Config, caData []byte, accountIDs ...string) error {
	sshConfigBytes := []byte(sshdConfig)
	err := tw.WriteHeader(&tar.Header{
		Name: "titus/etc/ssh/sshd_config",
		Mode: 0644,
		Size: int64(len(sshConfigBytes)),
	})
	if err != nil {
		return err
	}
	_, err = tw.Write(sshConfigBytes)
	if err != nil {
		return err
	}

	err = tw.WriteHeader(&tar.Header{
		Name: "titus/etc/ssh/trusted_user_ca_keys.pub",
		Mode: 0644,
		Size: int64(len(caData)),
	})
	if err != nil {
		return err
	}
	_, err = tw.Write(caData)
	if err != nil {
		return err
	}

	// The format that is used for SSH Users is:
	// $(unix username):$(app name):$(aws account id):$(task id)

	users := append(cfg.ContainerSSHDUsers, appName)
	for _, username := range users {
		lines := []string{}
		for _, accountID := range accountIDs {
			lines = append(
				lines,
				fmt.Sprintf("%s:%s:%s:%s", username, appName, accountID, taskID), // key scoped to username, appname, account ID, and task ID
				fmt.Sprintf("%s:%s:%s", appName, accountID, taskID),              // key has access to any username for this given app in this given account, with this task ID
				fmt.Sprintf("%s:%s", appName, accountID),                         // key has access to any username for this given app in this given account
				taskID,                                                           // key has access to any username on this task ID
				fmt.Sprintf("%s:%s", username, taskID),                           // key has access to this given username on this task ID
			)
		}
		line := []byte(strings.Join(lines, "\n"))
		err = tw.WriteHeader(&tar.Header{
			Name: fmt.Sprintf("titus/etc/ssh/authorized_principals_%s", username),
			Mode: 0644,
			Size: int64(len(line)),
		})
		if err != nil {
			return err
		}
		_, err = tw.Write(line)
		if err != nil {
			return err
		}
	}
	return nil
}

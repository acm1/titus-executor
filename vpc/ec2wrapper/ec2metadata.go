package ec2wrapper

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/service/ec2"
	set "github.com/deckarep/golang-set"
	"github.com/sirupsen/logrus"
)

// EC2MetadataClientWrapper wraps the EC2 library and provides some helper functions
type EC2MetadataClientWrapper struct {
	ec2metadata *ec2metadata.EC2Metadata
	logger      *logrus.Entry
}

// NetworkInterface represents an ENI, or a similar thing
type NetworkInterface interface {
	GetDeviceNumber() int
	GetInterfaceID() string
	GetSubnetID() string
	GetMAC() string
	GetSecurityGroupIds() map[string]struct{}
	GetIPv4Addresses() []string
	Refresh() error
	FreeIPv4Addresses(context.Context, client.ConfigProvider, []string) error
	// TODO: Add IPv6 addresses
}

// EC2NetworkInterface encapsulates information about network interfaces
type EC2NetworkInterface struct {
	mdc *EC2MetadataClientWrapper
	// deviceNumber is the AWS / EC2 device index, it should correlate to eth${DEVICEIDX}
	deviceNumber int
	// interfaceID is the ENI ID
	interfaceID string
	// subnetID is the Id of the subnet which the interface is in
	subnetID string
	// mac is the mac address in fully expanded form, i.e. 00:00:00:00:00:00
	mac string
	// securityGroupIds is the set of SGs (IDs) on the interface
	securityGroupIds map[string]struct{}
	// We don't use type net.IP here, because the same IP can end up being duplicated as a key
	// even though they're equal
	// ipv4Addresses is the set of currently assigned IPs -- The primary IPv4 address is the first in this list
	ipv4Addresses []string
}

// NewEC2MetadataClientWrapper creates a new ec2metadata wrapper instance
func NewEC2MetadataClientWrapper(session client.ConfigProvider, logger *logrus.Entry) *EC2MetadataClientWrapper {
	return &EC2MetadataClientWrapper{
		ec2metadata: ec2metadata.New(session),
		logger:      logger,
	}
}

// PrimaryInterfaceMac returns the mac of the primary interface
func (mdc *EC2MetadataClientWrapper) PrimaryInterfaceMac() (string, error) {
	val, err := mdc.getMetadata("mac")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(val), nil
}

// InstanceID returns the i- of the instance
func (mdc *EC2MetadataClientWrapper) InstanceID() (string, error) {
	val, err := mdc.getMetadata("instance-id")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(val), nil

}

// AvailabilityZone returns the qualified availability zone of the instance
func (mdc *EC2MetadataClientWrapper) AvailabilityZone() (string, error) {
	val, err := mdc.getMetadata("placement/availability-zone")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(val), nil
}

// Interfaces returns a map of mac addresses to interfaces
func (mdc *EC2MetadataClientWrapper) Interfaces() (map[string]NetworkInterface, error) {
	ret := make(map[string]NetworkInterface)
	val, err := mdc.getMetadata("network/interfaces/macs/")
	if err != nil {
		return ret, err
	}

	for _, mac := range strings.Split(strings.TrimSpace(val), "\n") {
		actualMac := strings.TrimRight(mac, "/")
		ret[actualMac], err = mdc.GetInterface(actualMac)
		if err != nil {
			return ret, err
		}
	}

	return ret, nil
}

// GetInterface fetches a populated interface by mac address
func (mdc *EC2MetadataClientWrapper) GetInterface(mac string) (NetworkInterface, error) {
	ret := &EC2NetworkInterface{
		mac: mac,
		mdc: mdc,
	}

	devNumberString, err := mdc.getDataForInterface(mac, "device-number")
	if err != nil {
		return ret, err
	}

	ret.deviceNumber, err = strconv.Atoi(devNumberString)
	if err != nil {
		return ret, err
	}

	ret.interfaceID, err = mdc.getDataForInterface(mac, "interface-id")
	if err != nil {
		return ret, err
	}

	ret.subnetID, err = mdc.getDataForInterface(mac, "subnet-id")
	if err != nil {
		return ret, err
	}

	return ret, ret.Refresh()
}

func (mdc *EC2MetadataClientWrapper) getDataForInterface(mac, data string) (string, error) {
	path := fmt.Sprintf("network/interfaces/macs/%s/%s", mac, data)
	str, err := mdc.getMetadata(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), err
}

func (mdc *EC2MetadataClientWrapper) getMetadata(path string) (string, error) {
	var val string
	var err error
	for i := 0; i < 10; i++ {
		val, err = mdc.ec2metadata.GetMetadata(path)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "429") {
			// Sleep a minimum of 10 milliseconds, and up to 100 ms
			jitter := time.Millisecond * time.Duration(rand.Intn(90)+10)
			time.Sleep(jitter)
		} else {
			// "Fatal error"
			break
		}
	}
	if err != nil {
		mdc.logger.WithField("path", path).Warning("Error while fetching metadata: ", err)
	}
	return val, err
}

// Refresh updates the security groups. and local IPv4 addresses for an interface
func (ni *EC2NetworkInterface) Refresh() error {
	ni.securityGroupIds = make(map[string]struct{})

	securityGroupIdsString, err := ni.mdc.getDataForInterface(ni.mac, "security-group-ids")
	if err != nil {
		return err
	}
	for _, sgID := range strings.Split(securityGroupIdsString, "\n") {
		ni.securityGroupIds[sgID] = struct{}{}
	}

	localIPv4s, err := ni.mdc.getDataForInterface(ni.mac, "local-ipv4s")
	if err != nil {
		return err
	}
	ni.ipv4Addresses = strings.Split(localIPv4s, "\n")
	for idx, addr := range ni.ipv4Addresses {
		ni.ipv4Addresses[idx] = strings.Trim(strings.TrimSpace(addr), "\x00")
	}

	return nil
}

// GetDeviceNumber get the AWS / EC2 device index, it should correlate to eth${DEVICEIDX}
func (ni *EC2NetworkInterface) GetDeviceNumber() int {
	return ni.deviceNumber
}

// GetInterfaceID returns the ENI ID
func (ni *EC2NetworkInterface) GetInterfaceID() string {
	return ni.interfaceID
}

// GetSubnetID gets the Id of the subnet which the interface is in
func (ni *EC2NetworkInterface) GetSubnetID() string {
	return ni.subnetID
}

// GetMAC is the mac address in fully expanded form, i.e. 00:00:00:00:00:00
func (ni *EC2NetworkInterface) GetMAC() string {
	if len(ni.mac) != len("00:00:00:00:00:00") {
		panic("Corrupted mac address")
	}
	return ni.mac
}

// GetSecurityGroupIds is the set of SGs (IDs) on the interface
func (ni *EC2NetworkInterface) GetSecurityGroupIds() map[string]struct{} {
	return ni.securityGroupIds
}

// GetIPv4Addresses is the set of currently assigned IPs -- The primary IPv4 address is the first in this list
func (ni *EC2NetworkInterface) GetIPv4Addresses() []string {
	return ni.ipv4Addresses
}

// FreeIPv4Addresses calls the EC2 / AWS API to free IP addresses
func (ni *EC2NetworkInterface) FreeIPv4Addresses(ctx context.Context, session client.ConfigProvider, deallocationList []string) error {

	unassignPrivateIPAddressesInput := &ec2.UnassignPrivateIpAddressesInput{
		PrivateIpAddresses: aws.StringSlice(deallocationList),
		NetworkInterfaceId: aws.String(ni.GetInterfaceID()),
	}

	_, err := ec2.New(session).UnassignPrivateIpAddressesWithContext(ctx, unassignPrivateIPAddressesInput)
	return err
}

// GetLockPath returns the path you should use for locks on this interface
func GetLockPath(ni NetworkInterface) string {
	return filepath.Join("interfaces", ni.GetMAC())
}

// GetIPv4AddressesAsSet returns a copy of the IPv4Addresses from this network interface as a set
func GetIPv4AddressesAsSet(ni NetworkInterface) set.Set {
	s := set.NewThreadUnsafeSet()
	for _, ip := range ni.GetIPv4Addresses() {
		s.Add(ip)
	}
	return s
}

// +build !windows

package dmidecode

import (
	"errors"
	"time"
)

type ReqType uint8

const (
	TypeBIOSInfo ReqType = iota
	TypeSystem
	TypeBaseBoard
	TypeChassis
	TypeProcessor
	TypeMemoryController
	TypeMemoryModule
	TypeCache
	TypePortConnector
	TypeSystemSlots
	TypeOnBoardDevices
	TypeOEMStrings
	TypeSystemConfigurationOptions
	TypeBIOSLanguage
	TypeGroupAssociations
	TypeSystemEventLog
	TypePhysicalMemoryArray
	TypeMemoryDevice
	Type32BitMemoryError
	TypeMemoryArrayMappedAddress
	TypeMemoryDeviceMappedAddress
	TypeBuiltInPointingDevice
	TypePortableBattery
	TypeSystemReset
	TypeHardwareSecurity
	TypeSystemPowerControls
	TypeVoltageProbe
	TypeCoolingDevice
	TypeTemperatureProbe
	TypeElectricalCurrentProbe
	TypeOOBRemoteAccess
	TypeBootIntegrityServices
	TypeSystemBoot
	Type64BitMemoryError
	TypeManagementDevice
	TypeManagementDeviceComponent
	TypeManagementDeviceThresholdData
	TypeMemoryChannel
	TypeIPMIDevice
	TypePowerSupply
	TypeAdditionalInformation
	TypeOnBoardDevice
	TypeVendorRangeBegin = 128
)

var (
	ErrPtrToSlice        = errors.New("dmidecode: parameter must be pointer to slice")
	ErrInvalidEntityType = errors.New("dmidecode: invalid entity type")
	ErrNotFount          = errors.New("dmidecode: not found")
)

type objType interface {
	Type() ReqType
}

type ReqBiosInfo struct {
	Vendor          string    `dmidecode:"Vendor"`
	Version         string    `dmidecode:"Version"`
	Address         string    `dmidecode:"Address"`
	BIOSRevision    string    `dmidecode:"BIOS Revision"`
	ReleaseDate     time.Time `dmidecode:"Release Date"`
	RuntimeSize     string    `dmidecode:"RuntimeSize"`
	RomSize         string    `dmidecode:"ROM Size"`
	Characteristics []string  `dmidecode:"Characteristics"`
}

type ReqSystem struct {
	Manufacturer string `dmidecode:"Manufacturer"`
	ProductName  string `dmidecode:"Product Name"`
	Version      string `dmidecode:"Version"`
	SerialNumber string `dmidecode:"Serial Number"`
	UUID         string `dmidecode:"UUID"`
	Family       string `dmidecode:"Family"`
}

type ReqBaseBoard struct {
}

type ReqChassis struct {
}

type ReqProcessor struct {
}

type ReqMemoryController struct {
}

type ReqMemoryModule struct {
}

type ReqCache struct {
}

type ReqPortConnector struct {
}

type ReqSystemSlots struct {
}

type ReqOnBoardDevices struct {
}

type ReqOEMStrings struct {
}

type ReqSystemConfigurationOptions struct {
}

type ReqBIOSLanguage struct {
}

type ReqGroupAssociations struct {
}

type ReqSystemEventLog struct {
}

type ReqPhysicalMemoryArray struct {
	Location               string `dmidecode:"Location"`
	Use                    string `dmidecode:"Use"`
	ErrorCorrection        string `dmidecode:"Error Correction"`
	MaximumCapacity        string `dmidecode:"Maximum Capacity"`
	ErrorInformationHandle string `dmidecode:"Error Information Handle"`
	NumberOfDevices        int    `dmidecode:"Number Of Devices"`
}

type ReqMemoryDevice struct {
}

type Req32BitMemoryError struct {
}

type ReqMemoryArrayMappedAddress struct {
}

type ReqMemoryDeviceMappedAddress struct {
}

type ReqBuiltInPointingDevice struct {
}

type ReqPortableBattery struct {
}

type ReqSystemReset struct {
}

type ReqHardwareSecurity struct {
}

type ReqSystemPowerControls struct {
}

type ReqVoltageProbe struct {
}

type ReqCoolingDevice struct {
}

type ReqTemperatureProbe struct {
}

type ReqElectricalCurrentProbe struct {
}

type ReqOOBRemoteAccess struct {
}

type ReqBootIntegrityServices struct {
}

type ReqSystemBoot struct {
}

type Req64BitMemoryError struct {
}

type ReqManagementDevice struct {
}

type ReqManagementDeviceComponent struct {
}

type ReqManagementDeviceThresholdData struct {
}

type ReqMemoryChannel struct {
}

type ReqIPMIDevice struct {
}

type ReqPowerSupply struct {
}

type ReqAdditionalInformation struct {
}

type ReqOnBoardDevice struct {
}

type ReqOemSpecificType struct {
	HeaderAndData []byte   `dmidecode:"Header and Data"`
	Strings       []string `dmidecode:"Strings"`
}

// not that much time spent on this part actually. Multiline selection is the trick

var _ objType = (*ReqBiosInfo)(nil)
var _ objType = (*ReqSystem)(nil)
var _ objType = (*ReqBaseBoard)(nil)
var _ objType = (*ReqChassis)(nil)
var _ objType = (*ReqProcessor)(nil)
var _ objType = (*ReqMemoryController)(nil)
var _ objType = (*ReqMemoryModule)(nil)
var _ objType = (*ReqCache)(nil)
var _ objType = (*ReqPortConnector)(nil)
var _ objType = (*ReqSystemSlots)(nil)
var _ objType = (*ReqOnBoardDevices)(nil)
var _ objType = (*ReqOEMStrings)(nil)
var _ objType = (*ReqSystemConfigurationOptions)(nil)
var _ objType = (*ReqBIOSLanguage)(nil)
var _ objType = (*ReqGroupAssociations)(nil)
var _ objType = (*ReqSystemEventLog)(nil)
var _ objType = (*ReqPhysicalMemoryArray)(nil)
var _ objType = (*ReqMemoryDevice)(nil)
var _ objType = (*Req32BitMemoryError)(nil)
var _ objType = (*ReqMemoryArrayMappedAddress)(nil)
var _ objType = (*ReqMemoryDeviceMappedAddress)(nil)
var _ objType = (*ReqBuiltInPointingDevice)(nil)
var _ objType = (*ReqPortableBattery)(nil)
var _ objType = (*ReqSystemReset)(nil)
var _ objType = (*ReqHardwareSecurity)(nil)
var _ objType = (*ReqSystemPowerControls)(nil)
var _ objType = (*ReqVoltageProbe)(nil)
var _ objType = (*ReqCoolingDevice)(nil)
var _ objType = (*ReqTemperatureProbe)(nil)
var _ objType = (*ReqElectricalCurrentProbe)(nil)
var _ objType = (*ReqOOBRemoteAccess)(nil)
var _ objType = (*ReqBootIntegrityServices)(nil)
var _ objType = (*ReqSystemBoot)(nil)
var _ objType = (*Req64BitMemoryError)(nil)
var _ objType = (*ReqManagementDevice)(nil)
var _ objType = (*ReqManagementDeviceComponent)(nil)
var _ objType = (*ReqManagementDeviceThresholdData)(nil)
var _ objType = (*ReqMemoryChannel)(nil)
var _ objType = (*ReqIPMIDevice)(nil)
var _ objType = (*ReqPowerSupply)(nil)
var _ objType = (*ReqAdditionalInformation)(nil)
var _ objType = (*ReqOnBoardDevice)(nil)
var _ objType = (*ReqOemSpecificType)(nil)

func newReqBiosInfo() interface{}                      { return &ReqBiosInfo{} }
func newReqSystem() interface{}                        { return &ReqSystem{} }
func newReqBaseBoard() interface{}                     { return &ReqBaseBoard{} }
func newReqChassis() interface{}                       { return &ReqChassis{} }
func newReqProcessor() interface{}                     { return &ReqProcessor{} }
func newReqMemoryController() interface{}              { return &ReqMemoryController{} }
func newReqMemoryModule() interface{}                  { return &ReqMemoryModule{} }
func newReqCache() interface{}                         { return &ReqCache{} }
func newReqPortConnector() interface{}                 { return &ReqPortConnector{} }
func newReqSystemSlots() interface{}                   { return &ReqSystemSlots{} }
func newReqOnBoardDevices() interface{}                { return &ReqOnBoardDevices{} }
func newReqOEMStrings() interface{}                    { return &ReqOEMStrings{} }
func newReqSystemConfigurationOptions() interface{}    { return &ReqSystemConfigurationOptions{} }
func newReqBIOSLanguage() interface{}                  { return &ReqBIOSLanguage{} }
func newReqGroupAssociations() interface{}             { return &ReqGroupAssociations{} }
func newReqSystemEventLog() interface{}                { return &ReqSystemEventLog{} }
func newReqPhysicalMemoryArray() interface{}           { return &ReqPhysicalMemoryArray{} }
func newReqMemoryDevice() interface{}                  { return &ReqMemoryDevice{} }
func newReq32BitMemoryError() interface{}              { return &Req32BitMemoryError{} }
func newReqMemoryArrayMappedAddress() interface{}      { return &ReqMemoryArrayMappedAddress{} }
func newReqMemoryDeviceMappedAddress() interface{}     { return &ReqMemoryDeviceMappedAddress{} }
func newReqBuiltInPointingDevice() interface{}         { return &ReqBuiltInPointingDevice{} }
func newReqPortableBattery() interface{}               { return &ReqPortableBattery{} }
func newReqSystemReset() interface{}                   { return &ReqSystemReset{} }
func newReqHardwareSecurity() interface{}              { return &ReqHardwareSecurity{} }
func newReqSystemPowerControls() interface{}           { return &ReqSystemPowerControls{} }
func newReqVoltageProbe() interface{}                  { return &ReqVoltageProbe{} }
func newReqCoolingDevice() interface{}                 { return &ReqCoolingDevice{} }
func newReqTemperatureProbe() interface{}              { return &ReqTemperatureProbe{} }
func newReqElectricalCurrentProbe() interface{}        { return &ReqElectricalCurrentProbe{} }
func newReqOOBRemoteAccess() interface{}               { return &ReqOOBRemoteAccess{} }
func newReqBootIntegrityServices() interface{}         { return &ReqBootIntegrityServices{} }
func newReqSystemBoot() interface{}                    { return &ReqSystemBoot{} }
func newReq64BitMemoryError() interface{}              { return &Req64BitMemoryError{} }
func newReqManagementDevice() interface{}              { return &ReqManagementDevice{} }
func newReqManagementDeviceComponent() interface{}     { return &ReqManagementDeviceComponent{} }
func newReqManagementDeviceThresholdData() interface{} { return &ReqManagementDeviceThresholdData{} }
func newReqMemoryChannel() interface{}                 { return &ReqMemoryChannel{} }
func newReqIPMIDevice() interface{}                    { return &ReqIPMIDevice{} }
func newReqPowerSupply() interface{}                   { return &ReqPowerSupply{} }
func newReqAdditionalInformation() interface{}         { return &ReqAdditionalInformation{} }
func newReqOnBoardDevice() interface{}                 { return &ReqOnBoardDevice{} }
func newReqOemSpecificType() interface{}               { return &ReqOemSpecificType{} }

func (req *ReqBiosInfo) Type() ReqType                      { return TypeBIOSInfo }
func (req *ReqSystem) Type() ReqType                        { return TypeSystem }
func (req *ReqBaseBoard) Type() ReqType                     { return TypeBaseBoard }
func (req *ReqChassis) Type() ReqType                       { return TypeChassis }
func (req *ReqProcessor) Type() ReqType                     { return TypeProcessor }
func (req *ReqMemoryController) Type() ReqType              { return TypeMemoryController }
func (req *ReqMemoryModule) Type() ReqType                  { return TypeMemoryModule }
func (req *ReqCache) Type() ReqType                         { return TypeCache }
func (req *ReqPortConnector) Type() ReqType                 { return TypePortConnector }
func (req *ReqSystemSlots) Type() ReqType                   { return TypeSystemSlots }
func (req *ReqOnBoardDevices) Type() ReqType                { return TypeOnBoardDevices }
func (req *ReqOEMStrings) Type() ReqType                    { return TypeOEMStrings }
func (req *ReqSystemConfigurationOptions) Type() ReqType    { return TypeSystemConfigurationOptions }
func (req *ReqBIOSLanguage) Type() ReqType                  { return TypeBIOSLanguage }
func (req *ReqGroupAssociations) Type() ReqType             { return TypeGroupAssociations }
func (req *ReqSystemEventLog) Type() ReqType                { return TypeSystemEventLog }
func (req *ReqPhysicalMemoryArray) Type() ReqType           { return TypePhysicalMemoryArray }
func (req *ReqMemoryDevice) Type() ReqType                  { return TypeMemoryDevice }
func (req *Req32BitMemoryError) Type() ReqType              { return Type32BitMemoryError }
func (req *ReqMemoryArrayMappedAddress) Type() ReqType      { return TypeMemoryArrayMappedAddress }
func (req *ReqMemoryDeviceMappedAddress) Type() ReqType     { return TypeMemoryDeviceMappedAddress }
func (req *ReqBuiltInPointingDevice) Type() ReqType         { return TypeBuiltInPointingDevice }
func (req *ReqPortableBattery) Type() ReqType               { return TypePortableBattery }
func (req *ReqSystemReset) Type() ReqType                   { return TypeSystemReset }
func (req *ReqHardwareSecurity) Type() ReqType              { return TypeHardwareSecurity }
func (req *ReqSystemPowerControls) Type() ReqType           { return TypeSystemPowerControls }
func (req *ReqVoltageProbe) Type() ReqType                  { return TypeVoltageProbe }
func (req *ReqCoolingDevice) Type() ReqType                 { return TypeCoolingDevice }
func (req *ReqTemperatureProbe) Type() ReqType              { return TypeTemperatureProbe }
func (req *ReqElectricalCurrentProbe) Type() ReqType        { return TypeElectricalCurrentProbe }
func (req *ReqOOBRemoteAccess) Type() ReqType               { return TypeOOBRemoteAccess }
func (req *ReqBootIntegrityServices) Type() ReqType         { return TypeBootIntegrityServices }
func (req *ReqSystemBoot) Type() ReqType                    { return TypeSystemBoot }
func (req *Req64BitMemoryError) Type() ReqType              { return Type64BitMemoryError }
func (req *ReqManagementDevice) Type() ReqType              { return TypeManagementDevice }
func (req *ReqManagementDeviceComponent) Type() ReqType     { return TypeManagementDeviceComponent }
func (req *ReqManagementDeviceThresholdData) Type() ReqType { return TypeManagementDeviceThresholdData }
func (req *ReqMemoryChannel) Type() ReqType                 { return TypeMemoryChannel }
func (req *ReqIPMIDevice) Type() ReqType                    { return TypeIPMIDevice }
func (req *ReqPowerSupply) Type() ReqType                   { return TypePowerSupply }
func (req *ReqAdditionalInformation) Type() ReqType         { return TypeAdditionalInformation }
func (req *ReqOnBoardDevice) Type() ReqType                 { return TypeOnBoardDevice }
func (req *ReqOemSpecificType) Type() ReqType               { return TypeVendorRangeBegin }

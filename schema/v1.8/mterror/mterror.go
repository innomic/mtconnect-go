// Code generated by xgen. DO NOT EDIT.

package mterror

// MTConnectError is The root node for MTConnect
type MTConnectError *MTConnectErrorType

// SenderType is The sender of the message
type SenderType string

// CreationTimeType is The date and time the document was created
type CreationTimeType string

// SequenceType is A sequence number
type SequenceType int

// TestIndicatorType is A debugging flag for testing.
type TestIndicatorType bool

// InstanceIdType is The instance number of the agent, used for fault tolerance
type InstanceIdType int

// BufferSizeType is The size of the agents buffer
type BufferSizeType int

// TimestampType is The time the sample was reported
type TimestampType string

// OccurrenceTimeType is The time a sample occurred
type OccurrenceTimeType string

// VersionType is A version number
type VersionType string

// NameType is A short name for any element
type NameType string

// UuidType is A universally unique id that uniquely identifies the element for
//         it's entire life
type UuidType string

// SerialNumberAttrType is A serial number for a piece of equipment
type SerialNumberAttrType string

// ItemSourceType is The measurement source
type ItemSourceType string

// RateType is A sample rate in milliseconds per sample
type RateType float32

// ComponentIdType is The id of the component (maps to the id from probe)
type ComponentIdType string

// IDType is An identifier
type IDType string

// SignificantDigitsValueType is The number significant digits
type SignificantDigitsValueType int

// CompositionIdType is The item's reference to the Device model composition
type CompositionIdType string

// DurationTimeType is A length of time in seconds
type DurationTimeType float32

// RemovedType is A flag indicating the item has been removed
type RemovedType bool

// KeyType is The key for adata set
type KeyType string

// DeviceModelChangeTimeType is A timestamp in 8601 format of the last update of the Device information
//         for any device
type DeviceModelChangeTimeType string

// AssetIdType is The unique id of the asset
type AssetIdType string

// AssetAttrTypeType is An asset type
type AssetAttrTypeType string

// AssetBufferSizeType is The maximum number of assets
type AssetBufferSizeType int

// AssetCountAttrType is The number of assets
type AssetCountAttrType int

// FloatListValueType is Common floating point sample value
type FloatListValueType []float32

// ThreeSpaceValueType is A three dimensional value 'X Y Z' or 'A B C'
type ThreeSpaceValueType *FloatListValueType

// DescriptionTextType is A description
type DescriptionTextType string

// DataItemEnumExtType is Extended tyoe for The types of measurements available
type DataItemEnumExtType string

// DataItemEnumEnum is A user variable
type DataItemEnumEnum string

// DataItemEnumType is The types of measurements available
type DataItemEnumType struct {
	DataItemEnumExtType string
	DataItemEnumEnum    string
}

// DataItemSubEnumExtType is Extended tyoe for The sub-types for a measurement
type DataItemSubEnumExtType string

// DataItemSubEnumEnum is Reported or measured value of amount included in the {{term(part)}}.
type DataItemSubEnumEnum string

// DataItemSubEnumType is The sub-types for a measurement
type DataItemSubEnumType struct {
	DataItemSubEnumExtType string
	DataItemSubEnumEnum    string
}

// DataItemStatisticsExtType is Extended tyoe for Statistical operations on data
type DataItemStatisticsExtType string

// DataItemStatisticsEnum is Statistical Standard Deviation value calculated for the data item
//             during the calculation period.
type DataItemStatisticsEnum string

// DataItemStatisticsType is Statistical operations on data
type DataItemStatisticsType struct {
	DataItemStatisticsExtType string
	DataItemStatisticsEnum    string
}

// UnitsExtType is Extended tyoe for The units supported
type UnitsExtType string

// UnitsEnum is A 3D Unit Vector. Space delimited list of three floating point
//             numbers.
type UnitsEnum string

// UnitsType is The units supported
type UnitsType struct {
	UnitsExtType string
	UnitsEnum    string
}

// NativeUnitsExtType is Extended tyoe for The units supported for the source equipment that can
//         be converted into MTC Units.
type NativeUnitsExtType string

// NativeUnitsEnum is Pascal per minute.
type NativeUnitsEnum string

// NativeUnitsType is The units supported for the source equipment that can be converted into
//         MTC Units.
type NativeUnitsType struct {
	NativeUnitsExtType string
	NativeUnitsEnum    string
}

// CoordinateSystemEnumType is The coordinate system that represents the working area for a
//             particular workpiece whose origin is shifted within the `MACHINE`
//             coordinate system. If the `WORK` coordinates are not currently
//             defined in the piece of equipment, the `MACHINE` coordinates will be
//             used.
type CoordinateSystemEnumType string

// DataItemResetValueExtType is Extended tyoe for The reset intervals
type DataItemResetValueExtType string

// DataItemResetValueEnum is The value of the {{term(Data Item)}} is to be reset at the end of a
//             7-day period.
type DataItemResetValueEnum string

// DataItemResetValueType is The reset intervals
type DataItemResetValueType struct {
	DataItemResetValueExtType string
	DataItemResetValueEnum    string
}

// HeaderAttributesType is A timestamp in 8601 format of the last update of the Device
//           information for any device
type HeaderAttributesType struct {
	VersionAttr               string `xml:"version,attr"`
	CreationTimeAttr          string `xml:"creationTime,attr"`
	TestIndicatorAttr         bool   `xml:"testIndicator,attr,omitempty"`
	InstanceIdAttr            int    `xml:"instanceId,attr"`
	SenderAttr                string `xml:"sender,attr"`
	DeviceModelChangeTimeAttr string `xml:"deviceModelChangeTime,attr"`
}

// HeaderType is Supplemental data usually placed at the beginning of a
//         {{term(Document)}}.
type HeaderType struct {
	VersionAttr       string `xml:"version,attr"`
	CreationTimeAttr  string `xml:"creationTime,attr"`
	TestIndicatorAttr bool   `xml:"testIndicator,attr,omitempty"`
	InstanceIdAttr    int    `xml:"instanceId,attr"`
	SenderAttr        string `xml:"sender,attr"`
	BufferSizeAttr    int    `xml:"bufferSize,attr"`
	Value             string `xml:",chardata"`
}

// MTConnectErrorType is A collection of errors
type MTConnectErrorType struct {
	Header *HeaderType `xml:"Header"`
	Error  *ErrorType  `xml:"Error"`
	Errors *ErrorsType `xml:"Errors"`
}

// ErrorsType is An error result
type ErrorsType struct {
	Error []*ErrorType `xml:"Error"`
}

// ErrorDescriptionType is The text description of the error
type ErrorDescriptionType string

// ErrorCodeType is The asset ID cannot be found
type ErrorCodeType string

// ErrorType is An error description
type ErrorType struct {
	ErrorCodeAttr string `xml:"errorCode,attr"`
	Value         string `xml:",chardata"`
}

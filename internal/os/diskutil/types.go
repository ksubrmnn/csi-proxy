package diskutil

/*
typedef struct _STORAGE_PROPERTY_QUERY {
	STORAGE_PROPERTY_ID PropertyId;
	STORAGE_QUERY_TYPE  QueryType;
	BYTE                AdditionalParameters[1];
  } STORAGE_PROPERTY_QUERY, *PSTORAGE_PROPERTY_QUERY;
*/

type StoragePropertyID uint32

const (
	StorageDeviceProperty StoragePropertyID = iota
	StorageAdapterProperty
	StorageDeviceIDProperty
	StorageDeviceUniqueIDProperty
	StorageDeviceWriteCacheProperty
	StorageMiniportProperty
	StorageAccessAlignmentProperty
	StorageDeviceSeekPenaltyProperty
	StorageDeviceTrimProperty
	StorageDeviceWriteAggregationProperty
	StorageDeviceDeviceTelemetryProperty
	StorageDeviceLBProvisioningProperty
	StorageDevicePowerProperty
	StorageDeviceCopyOffloadProperty
	StorageDeviceResiliencyProperty
	StorageDeviceMediumProductType
	StorageAdapterRpmbProperty
	StorageAdapterCryptoProperty
	StorageDeviceIoCapabilityProperty
	StorageAdapterProtocolSpecificProperty
	StorageDeviceProtocolSpecificProperty
	StorageAdapterTemperatureProperty
	StorageDeviceTemperatureProperty
	StorageAdapterPhysicalTopologyProperty
	StorageDevicePhysicalTopologyProperty
	StorageDeviceAttributesProperty
	StorageDeviceManagementStatus
	StorageAdapterSerialNumberProperty
	StorageDeviceLocationProperty
	StorageDeviceNumaProperty
	StorageDeviceZonedDeviceProperty
	StorageDeviceUnsafeShutdownCount
	StorageDeviceEnduranceProperty
)

type StorageQueryType uint32

const (
	PropertyStandardQuery StorageQueryType = iota
	PropertyExistsQuery
	PropertyMaskQuery
	PropertyQueryMaxDefined
)

type StoragePropertyQuery struct {
	PropertyID           StoragePropertyID
	QueryType            StorageQueryType
	AdditionalParameters byte
}

/*
BOOL DeviceIoControl(
  HANDLE       hDevice,
  DWORD        dwIoControlCode,
  LPVOID       lpInBuffer,
  DWORD        nInBufferSize,
  LPVOID       lpOutBuffer,
  DWORD        nOutBufferSize,
  LPDWORD      lpBytesReturned,
  LPOVERLAPPED lpOverlapped
);

typedef struct _STORAGE_DEVICE_NUMBER {
  DEVICE_TYPE DeviceType;
  DWORD       DeviceNumber;
  DWORD       PartitionNumber;
} STORAGE_DEVICE_NUMBER, *PSTORAGE_DEVICE_NUMBER;
*/

type StorageDeviceNumber struct {
	DeviceType      DeviceType
	DeviceNumber    uint32
	PartitionNumber uint32
}
type DeviceType uint16

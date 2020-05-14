// +build windows

package disk

import (
	"syscall"
	"unsafe"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
)

const (
	IOCTL_STORAGE_GET_DEVICE_NUMBER = 0x2D1080
)

/*
DWORD GetDiskNumber(HANDLE Disk, PLONG Number) {
	ULONG bytes;
	DWORD err = ERROR_SUCCESS;
	STORAGE_DEVICE_NUMBER devNum;
	*Number = -1;

	if (!DeviceIoControl(Disk, IOCTL_STORAGE_GET_DEVICE_NUMBER, NULL, 0, &devNum, sizeof(STORAGE_DEVICE_NUMBER), &bytes, FALSE)) {
		err = GetLastError();
		printf("IOCTL_STORAGE_GET_DEVICE_NUMBER failed: %d\n", err);
		return err;
	}
	*Number = devNum.DeviceNumber;
	return err;
}
*/
func GetDiskNumber(disk syscall.Handle, number *uint32) error {
	var bytes uint32
	devNum := StorageDeviceNumber{}

	buflen := uint32(unsafe.Sizeof(devNum.DeviceType)) + uint32(unsafe.Sizeof(devNum.DeviceNumber)) + uint32(unsafe.Sizeof(devNum.PartitionNumber))
	err := syscall.DeviceIoControl(disk, IOCTL_STORAGE_GET_DEVICE_NUMBER, nil, 0, (*byte)(unsafe.Pointer(&devNum)), buflen, &bytes, nil)
	*number = devNum.DeviceNumber
	return err
}

/*
DWORD DiskHasPage83Id(HANDLE Disk, PCHAR MatchId, ULONG MatchLen, PBOOL Found) {
	STORAGE_PROPERTY_QUERY qry;
	PSTORAGE_DEVICE_ID_DESCRIPTOR pDevIdDesc = NULL;
	PSTORAGE_IDENTIFIER pId = NULL;
	ULONG buffer_sz = 4*1024;
	ULONG sz = 0;
	ULONG m, n;
	DWORD err = ERROR_SUCCESS;

	*Found = FALSE;
	pDevIdDesc = (PSTORAGE_DEVICE_ID_DESCRIPTOR) malloc(buffer_sz);
	if (pDevIdDesc == NULL) {
		printf("Error allocating memory to get the query storage descriptors \n");
		err = ERROR_NOT_ENOUGH_MEMORY;
		goto EXIT;
	}

	qry.QueryType = PropertyStandardQuery;
	qry.PropertyId = StorageDeviceIdProperty;

	if (!DeviceIoControl(Disk, IOCTL_STORAGE_QUERY_PROPERTY, &qry, sizeof(STORAGE_PROPERTY_QUERY), pDevIdDesc, buffer_sz, &sz, NULL)) {
		err = GetLastError();
		printf("IOCTL_STORAGE_QUERY_PROPERTY failed: %d \n", err);
		goto EXIT;
	}

	pId = (PSTORAGE_IDENTIFIER) pDevIdDesc->Identifiers;

	for (n = 0; n < pDevIdDesc->NumberOfIdentifiers; n++) {
		if ((pId->CodeSet == StorageIdCodeSetAscii) && (pId->Association == StorageIdAssocDevice)) {
			if (MatchLen > pId->IdentifierSize) {
				continue;
			}
			for (m = 0; m < (pId->IdentifierSize - MatchLen + 1); m++) {
				if (memcmp(MatchId, pId->Identifier + m, MatchLen) == 0) {
					*Found = TRUE;
					goto EXIT;
				}
			}
		}
		pId = (PSTORAGE_IDENTIFIER)((ULONG_PTR)pId + pId->NextOffset);
	}

EXIT:
	if (pDevIdDesc != NULL) {
		free(pDevIdDesc);
	}
	return err;
}
*/

func DiskHasPage83Id(disk syscall.Handle, matchID *char, matchLen uint64, found *bool) error {
	var query StoragePropertyQuery

}

// This file was autogenerated using mkcode -- registry.go
// DO NOT EDIT.

package ntdll

import "unsafe"

type KeyInformationClass uint32

const (
	KeyBasicInformation          KeyInformationClass = 0
	KeyNodeInformation                               = 1
	KeyFullInformation                               = 2
	KeyNameInformation                               = 3
	KeyCachedInformation                             = 4
	KeyFlagsInformation                              = 5
	KeyVirtualizationInformation                     = 6
	KeyHandleTagsInformation                         = 7
	MaxKeyInfoClass                                  = 8
)

type KeyValueInformationClass uint32

const (
	KeyValueBasicInformation          KeyValueInformationClass = 0
	KeyValueFullInformation                                    = 1
	KeyValuePartialInformation                                 = 2
	KeyValueFullInformationAlign64                             = 3
	KeyValuePartialInformationAlign64                          = 4
	MaxKeyValueInfoClass                                       = 5
)

var (
	procNtCreateKey         = modntdll.NewProc("NtCreateKey")
	procNtDeleteKey         = modntdll.NewProc("NtDeleteKey")
	procNtEnumerateKey      = modntdll.NewProc("NtEnumerateKey")
	procNtEnumerateValueKey = modntdll.NewProc("NtEnumerateValueKey")
	procNtFlushKey          = modntdll.NewProc("NtFlushKey")
	procNtNotifyChangeKey   = modntdll.NewProc("NtNotifyChangeKey")
	procNtOpenKey           = modntdll.NewProc("NtOpenKey")
	procNtQueryValueKey     = modntdll.NewProc("NtQueryValueKey")
	procNtSetValueKey       = modntdll.NewProc("NtSetValueKey")
)

func NtCreateKey(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes, TitleIndex uint32, Class *UnicodeString, CreateOptions uint32, Disposition *uint32) NtStatus {
	r0, _, _ := procNtCreateKey.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)), uintptr(TitleIndex), uintptr(unsafe.Pointer(Class)), uintptr(CreateOptions), uintptr(unsafe.Pointer(Disposition)))
	return NtStatus(r0)
}

func NtDeleteKey(KeyHandle Handle) NtStatus {
	r0, _, _ := procNtDeleteKey.Call(uintptr(KeyHandle))
	return NtStatus(r0)
}

func NtEnumerateKey(KeyHandle Handle, Index uint32, KeyInformationClass KeyInformationClass, KeyInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtEnumerateKey.Call(uintptr(KeyHandle), uintptr(Index), uintptr(KeyInformationClass), uintptr(unsafe.Pointer(KeyInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtEnumerateValueKey(KeyHandle Handle, Index uint32, KeyValueInformationClass KeyValueInformationClass, KeyValueInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtEnumerateValueKey.Call(uintptr(KeyHandle), uintptr(Index), uintptr(KeyValueInformationClass), uintptr(unsafe.Pointer(KeyValueInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtFlushKey(KeyHandle Handle) NtStatus {
	r0, _, _ := procNtFlushKey.Call(uintptr(KeyHandle))
	return NtStatus(r0)
}

func NtNotifyChangeKey(KeyHandle Handle, Event Handle, ApcRoutine *IoApcRoutine, ApcContext *byte, IoStatusBlock *IoStatusBlock, CompletionFilter uint32, WatchTree bool, Buffer *byte, BufferSize uint32, Asynchronous bool) NtStatus {
	r0, _, _ := procNtNotifyChangeKey.Call(uintptr(KeyHandle), uintptr(Event), uintptr(unsafe.Pointer(ApcRoutine)), uintptr(unsafe.Pointer(ApcContext)), uintptr(unsafe.Pointer(IoStatusBlock)), uintptr(CompletionFilter), fromBool(WatchTree), uintptr(unsafe.Pointer(Buffer)), uintptr(BufferSize), fromBool(Asynchronous))
	return NtStatus(r0)
}

func NtOpenKey(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes) NtStatus {
	r0, _, _ := procNtOpenKey.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

func NtQueryValueKey(KeyHandle Handle, ValueName *UnicodeString, KeyValueInformationClass KeyValueInformationClass, KeyValueInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtQueryValueKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(ValueName)), uintptr(KeyValueInformationClass), uintptr(unsafe.Pointer(KeyValueInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtSetValueKey(KeyHandle Handle, ValueName *UnicodeString, TitleIndex uint32, Type uint32, Data *byte, DataSize uint32) NtStatus {
	r0, _, _ := procNtSetValueKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(ValueName)), uintptr(TitleIndex), uintptr(Type), uintptr(unsafe.Pointer(Data)), uintptr(DataSize))
	return NtStatus(r0)
}

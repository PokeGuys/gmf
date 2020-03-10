package gmf

/*

#cgo pkg-config: libavutil

#include "libavutil/dict.h"

*/
import "C"
import "fmt"

type AVDictionary struct {
	avDictionary *C.struct_AVDictionary
}

type AVDictionaryEntry C.AVDictionaryEntry

func NewAVDictionary() *AVDictionary {
	return &AVDictionary{
		avDictionary: &C.struct_AVDictionary{},
	}
}

func (cp *AVDictionary) Free() {
	C.av_dict_free(&cp.avDictionary)
}

func (cp *AVDictionary) GetCount() int {
	return int(C.av_dict_count(cp.avDictionary))
}

func (cp *AVDictionary) Get(key string) string {
	element := C.av_dict_get(cp.avDictionary, C.CString(key), nil, C.AV_DICT_IGNORE_SUFFIX)
	if element == nil {
		return ""
	}

	return C.GoString(element.value)
}

func (cp *AVDictionary) Set(key, value string) error {
	ret := C.av_dict_set(&cp.avDictionary, C.CString(key), C.CString(value), C.int(0))
	if ret < 0 {
		return fmt.Errorf("failed to set dictionary %s value - %d", key, value)
	}

	return nil
}

func (cp *AVDictionary) SetInt(key string, value int) error {
	ret := C.av_dict_set_int(&cp.avDictionary, C.CString(key), C.long(value), C.int(0))
	if ret < 0 {
		return fmt.Errorf("failed to set dictionary %s value - %d", key, value)
	}

	return nil
}

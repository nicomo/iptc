// Package iptc exposes Golang bindings for libiptcdata.
package iptc

/*
#cgo LDFLAGS: -liptcdata

#include <stdlib.h>
#include <libiptcdata/iptc-data.h>

IptcDataSet *get_iptc_dataset(IptcData *iptcData, unsigned int i);
*/
import "C"

import (
        "fmt"
        "unsafe"
        "reflect"

        "github.com/mitchellh/mapstructure"
        "github.com/zidizei/iptc/data"
)

// MetaData contains the IPTC metadata of a file.
type MetaData struct {
        File       string
        Entries    map[string]data.Entry
}

// Decode the MetaData entries into the specified struct t.
func (d MetaData) Decode(t interface{}) error {
        decoded := make(map[string]interface{})

        v := reflect.Indirect(reflect.ValueOf(t))

	for i := 0; i < v.NumField(); i++ {
                k := v.Type().Field(i).Name
                v, ok := d.Entries[k]

                if ok == false {
                        // We could not find an element named k inside
                        // the Entries map, which means the file does not have that
                        // particular IPTC metadata entry available.
                        continue
                }

                switch reflect.TypeOf(v).String() {
                case "*data.ListEntry":
                        list, err := d.ListTag(k)
                        if err != nil {
                                continue
                        }

                        decoded[k] = list

                case "*data.NumberEntry":
                        num, err := d.NumberTag(k)
                        if err != nil {
                                continue
                        }

                        decoded[k] = num

                default:
                        str, err := d.StringTag(k)
                        if err != nil {
                                continue
                        }

                        decoded[k] = str

                }
	}

        return mapstructure.Decode(decoded, t)
}

// StringTag returns the specified tag's value as a string.
func (d MetaData) StringTag(tag string) (string, error) {
        raw, ok := d.Entries[tag]
        if ok == false {
                return "", fmt.Errorf(ErrIptcTagNotFound.Error(), tag, d.File)
        }

        value, ok := raw.Value().(string)
        if ok == false {
                return "", fmt.Errorf(ErrIptcTagType.Error(), tag, d.File)
        }

        return value, nil
}

// ListTag returns the specified tag's value as a string.
func (d MetaData) ListTag(tag string) ([]string, error) {
        raw, ok := d.Entries[tag]
        if ok == false {
                return nil, fmt.Errorf(ErrIptcTagNotFound.Error(), tag, d.File)
        }

        value, ok := raw.Value().([]string)
        if ok == false {
                return nil, fmt.Errorf(ErrIptcTagType.Error(), tag, d.File)
        }

        return value, nil
}

// NumberTag returns the specified tag's value as a string.
func (d MetaData) NumberTag(tag string) (int, error) {
        raw, ok := d.Entries[tag]
        if ok == false {
                return 0, fmt.Errorf(ErrIptcTagNotFound.Error(), tag, d.File)
        }

        value, ok := raw.Value().(int)
        if ok == false {
                return 0, fmt.Errorf(ErrIptcTagType.Error(), tag, d.File)
        }

        return value, nil
}

// Open the given path and attempts to return any IPTC data read.
func Open(file string) (MetaData, error) {
        cfile      := C.CString(file)
        iptcData   := C.iptc_data_new_from_jpeg(cfile)
        parsedData := MetaData { File: file }

        C.free(unsafe.Pointer(cfile))

        if iptcData == nil {
                return parsedData, fmt.Errorf(ErrNoIptcData.Error(), file)
        }

        defer func() {
                C.iptc_data_unref(iptcData)
        }()

        parsedData.Entries = parseIptcData(iptcData)

        return parsedData, nil
}

// Load the IPTC data from the specified file into the specified struct
func Load(file string, t interface{}) error {
        metadata, err := Open(file)
        if err != nil {
                return err
        }

        return metadata.Decode(t)
}

// Parses an IPTC data blob generating a map of records and tags to
// string values.
func parseIptcData(iptcData *C.IptcData) (map[string]data.Entry) {
        parsed := make(map[string]data.Entry)

        for i := C.uint(0); i < iptcData.count; i++ {
                dataSet := C.get_iptc_dataset(iptcData, i)

                recordTags := data.IptcRecords[int(dataSet.record)]
                dataEntry  := recordTags[int(dataSet.tag)]

                if _, err := parsed[dataEntry.GetName()]; err == false {
                        parsed[dataEntry.GetName()] = dataEntry.New()
                }

                var recordValue string

                switch C.iptc_dataset_get_format(dataSet) {
                case C.IPTC_FORMAT_BYTE, C.IPTC_FORMAT_SHORT, C.IPTC_FORMAT_LONG:
                        recordValue = fmt.Sprintf("%d", C.iptc_dataset_get_value(dataSet))

                case C.IPTC_FORMAT_BINARY:
                        value := make([]C.char, 256)
                        C.iptc_dataset_get_as_str(dataSet, &value[0], C.uint(len(value)))

                        recordValue = fmt.Sprintf("%c", value[:(dataSet.size*3)-1])

                default:
                        value        := make([]C.uchar, 256)
                        actualLength := C.iptc_dataset_get_data(dataSet, &value[0], C.uint(len(value)))

                        recordValue = fmt.Sprintf("%s", value[:actualLength-1])
                }

                parsed[dataEntry.GetName()].AppendRaw(recordValue)
        }

        return parsed
}

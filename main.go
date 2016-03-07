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
        "errors"
        "fmt"
        "unsafe"
)

// Record contains the IPTC records of a file.
type Record map[string]DataEntry

// DataEntry contains the parsed IPTC record data.
type DataEntry struct {
        ID    int
        Name  string
        Data  map[string][]string
}

// ErrNoIptcData errors indicate files without any IPTC data.
var ErrNoIptcData = errors.New("no IPTC data in file %s")

// IptcRecords maps IPTC record ID types to their name.
var IptcRecords = map[int]string {
        1:	"IPTCEnvelope",
        2:      "IPTCApplication",
}

// IptcEnvelopeTags maps IPTC envelope tag IDs to their tag's name.
// See: http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
var IptcEnvelopeTags = map[int]string {
        0:	"EnvelopeRecordVersion",
        5:	"Destination",
        20:	"FileFormat",
        22:	"FileVersion",
        30:	"ServiceIdentifier",
        40:	"EnvelopeNumber",
        50:	"ProductID",
        60:	"EnvelopePriority",
        70:	"DateSent",
        80:	"TimeSent",
        90:	"CodedCharacterSet",
        100:	"UniqueObjectName",
        120:	"ARMIdentifier",
        122:	"ARMVersion",
}

// IptcApplicationTags maps IPTC application tag IDs to their tag's name.
// See: http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
var IptcApplicationTags = map[int]string {
        0:      "ApplicationRecordVersion",
        3:      "ObjectTypeReference",
        4:      "ObjectAttributeReference",
        5:      "ObjectName",
        7:      "EditStatus",
        8:      "EditorialUpdate",
        10:     "Urgency",
        12:     "SubjectReference",
        15:     "Category",
        20:     "SupplementalCategories",
        22:     "FixtureIdentifier",
        25:     "Keywords",
        26:     "ContentLocationCode",
        27:     "ContentLocationName",
        30:     "ReleaseDate",
        35:     "ReleaseTime",
        37:     "ExpirationDate",
        38:     "ExpirationTime",
        40:     "SpecialInstructions",
        42:     "ActionAdvised",
        45:     "ReferenceService",
        47:     "ReferenceDate",
        50:     "ReferenceNumber",
        55:     "DateCreated",
        60:     "TimeCreated",
        62:     "DigitalCreationDate",
        63:     "DigitalCreationTime",
        65:     "OriginatingProgram",
        70:	"ProgramVersion",
        75:	"ObjectCycle",
        80:	"By-line",
        85:	"By-lineTitle",
        90:	"City",
        92:	"Sub-location",
        95:	"Province-State",
        100:	"Country-PrimaryLocationCode",
        101:	"Country-PrimaryLocationName",
        103:	"OriginalTransmissionReference",
        105:	"Headline",
        110:	"Credit",
        115:	"Source",
        116:	"CopyrightNotice",
        118:	"Contact",
        120:	"Caption-Abstract",
        121:	"LocalCaption",
        122:	"Writer-Editor",
        125:	"RasterizedCaption",
        130:	"ImageType",
        131:	"ImageOrientation",
        135:	"LanguageIdentifier",
        150:	"AudioType",
        151:	"AudioSamplingRate",
        152:	"AudioSamplingResolution",
        153:	"AudioDuration",
        154:	"AudioOutcue",
        184:	"JobID",
        185:	"MasterDocumentID",
        186:	"ShortDocumentID",
        187:	"UniqueDocumentID",
        188:	"OwnerID",
        200:	"ObjectPreviewFileFormat",
        201:	"ObjectPreviewFileVersion",
        202:	"ObjectPreviewData",
        221:	"Prefs",
        225:	"ClassifyState",
        228:	"SimilarityIndex",
        230:	"DocumentNotes",
        231:	"DocumentHistory",
        232:	"ExifCameraInfo",
        255:	"CatalogSets",
}

// Open the given path and attempts to return any IPTC data read.
func Open(file string) (Record, error) {
        cfile    := C.CString(file)
        iptcData := C.iptc_data_new_from_jpeg(cfile)

        C.free(unsafe.Pointer(cfile))

        if iptcData == nil {
                return Record{}, fmt.Errorf(ErrNoIptcData.Error(), file)
        }

        defer func() {
                C.iptc_data_unref(iptcData)
        }()

        return parseIptcData(iptcData)
}


// Parses an IPTC data blob generating a map of records and tags to
// string values.
func parseIptcData(iptcData *C.IptcData) (Record, error) {
        parsed := Record{}

        for i := C.uint(0); i < iptcData.count; i++ {
                dataSet := C.get_iptc_dataset(iptcData, i)

                recordID   := IptcRecords[int(dataSet.record)]
                recordName := IptcApplicationTags[int(dataSet.tag)]

                if _, err := parsed[recordID]; err == false {
                        parsed[recordID] = DataEntry {
                                ID: int(dataSet.record),
                                Name: recordID,
                                Data: make(map[string][]string),
                        }
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

                parsed[recordID].Data[recordName] = append(parsed[recordID].Data[recordName], recordValue)
        }

        return parsed, nil
}

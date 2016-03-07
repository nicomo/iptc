package iptc

import "github.com/zidizei/iptc/data"


// TagsTable are maps of IPTC tag IDs with their tag names.
type TagsTable map[int]data.Entry


// IptcRecords maps IPTC record ID types to their name.
var IptcRecords = map[int]TagsTable {
        1:	IptcEnvelopeTags,
        2:      IptcApplicationTags,
}

// IptcEnvelopeTags maps IPTC envelope tag IDs to their tag's name.
// See: http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
var IptcEnvelopeTags = TagsTable {
        0:	&data.NumberEntry{ ID: 0,   Name: "EnvelopeRecordVersion"},
        5:	&data.StringEntry{ ID: 5,   Name: "Destination" },
        20:	&data.NumberEntry{ ID: 20,  Name: "FileFormat"},
        22:	&data.NumberEntry{ ID: 22,  Name: "FileVersion"},
        30:	&data.StringEntry{ ID: 30,  Name: "ServiceIdentifier" },
        40:	&data.NumberEntry{ ID: 40,  Name: "EnvelopeNumber"},
        50:	&data.StringEntry{ ID: 50,  Name: "ProductID" },
        60:	&data.NumberEntry{ ID: 60,  Name: "EnvelopePriority"},
        70:	&data.NumberEntry{ ID: 70,  Name: "DateSent"},
        80:	&data.StringEntry{ ID: 80,  Name: "TimeSent" },
        90:	&data.StringEntry{ ID: 90,  Name: "CodedCharacterSet" },
        100:	&data.StringEntry{ ID: 100, Name: "UniqueObjectName" },
        120:	&data.NumberEntry{ ID: 120, Name: "ARMIdentifier"},
        122:	&data.NumberEntry{ ID: 122, Name: "ARMVersion"},
}

// IptcApplicationTags maps IPTC application tag IDs to their tag's name.
// See: http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
var IptcApplicationTags = TagsTable {
        0:      &data.NumberEntry{ ID: 0, Name: "ApplicationRecordVersion"},
        3:      &data.StringEntry{ ID: 3, Name: "ObjectTypeReference"},
        4:      &data.StringEntry{ ID: 4, Name: "ObjectAttributeReference"},
        5:      &data.StringEntry{ ID: 5, Name: "ObjectName"},
        7:      &data.StringEntry{ ID: 7, Name: "EditStatus"},
        8:      &data.NumberEntry{ ID: 8, Name: "EditorialUpdate"},
        10:     &data.NumberEntry{ ID: 10, Name: "Urgency"},
        12:     &data.StringEntry{ ID: 12, Name: "SubjectReference"},
        15:     &data.StringEntry{ ID: 15, Name: "Category"},
        20:     &data.StringEntry{ ID: 20, Name: "SupplementalCategories"},
        22:     &data.StringEntry{ ID: 22, Name: "FixtureIdentifier"},
        25:     &data.ListEntry{ ID: 25, Name: "Keywords"},
        26:     &data.StringEntry{ ID: 26, Name: "ContentLocationCode"},
        27:     &data.StringEntry{ ID: 27, Name: "ContentLocationName"},
        30:     &data.NumberEntry{ ID: 30, Name: "ReleaseDate"},
        35:     &data.StringEntry{ ID: 35, Name: "ReleaseTime"},
        37:     &data.NumberEntry{ ID: 37, Name: "ExpirationDate"},
        38:     &data.StringEntry{ ID: 38, Name: "ExpirationTime"},
        40:     &data.StringEntry{ ID: 40, Name: "SpecialInstructions"},
        42:     &data.NumberEntry{ ID: 42, Name: "ActionAdvised"},
        45:     &data.StringEntry{ ID: 45, Name: "ReferenceService"},
        47:     &data.NumberEntry{ ID: 47, Name: "ReferenceDate"},
        50:     &data.NumberEntry{ ID: 50, Name: "ReferenceNumber"},
        55:     &data.NumberEntry{ ID: 55, Name: "DateCreated"},
        60:     &data.StringEntry{ ID: 60, Name: "TimeCreated"},
        62:     &data.NumberEntry{ ID: 62, Name: "DigitalCreationDate"},
        63:     &data.StringEntry{ ID: 63, Name: "DigitalCreationTime"},
        65:     &data.StringEntry{ ID: 65, Name: "OriginatingProgram"},
        70:	&data.StringEntry{ ID: 70, Name: "ProgramVersion"},
        75:	&data.StringEntry{ ID: 75, Name: "ObjectCycle"},
        80:	&data.StringEntry{ ID: 80, Name: "By-line"},
        85:	&data.StringEntry{ ID: 85, Name: "By-lineTitle"},
        90:	&data.StringEntry{ ID: 90, Name: "City"},
        92:	&data.StringEntry{ ID: 92, Name: "Sub-location"},
        95:	&data.StringEntry{ ID: 95, Name: "Province-State"},
        100:	&data.StringEntry{ ID: 100, Name: "Country-PrimaryLocationCode"},
        101:	&data.StringEntry{ ID: 101, Name: "Country-PrimaryLocationName"},
        103:	&data.StringEntry{ ID: 103, Name: "OriginalTransmissionReference"},
        105:	&data.StringEntry{ ID: 105, Name: "Headline"},
        110:	&data.StringEntry{ ID: 110, Name: "Credit"},
        115:	&data.StringEntry{ ID: 115, Name: "Source"},
        116:	&data.StringEntry{ ID: 116, Name: "CopyrightNotice"},
        118:	&data.StringEntry{ ID: 118, Name: "Contact"},
        120:	&data.StringEntry{ ID: 120, Name: "Caption-Abstract"},
        121:	&data.StringEntry{ ID: 121, Name: "LocalCaption"},
        122:	&data.StringEntry{ ID: 122, Name: "Writer-Editor"},
        125:	&data.StringEntry{ ID: 125, Name: "RasterizedCaption"},
        130:	&data.StringEntry{ ID: 130, Name: "ImageType"},
        131:	&data.ImageOrientationEntry{ ID: 131, Name: "ImageOrientation"},
        135:	&data.StringEntry{ ID: 135, Name: "LanguageIdentifier"},
        150:	&data.AudioTypeEntry{ ID: 150, Name: "AudioType"},
        151:	&data.NumberEntry{ ID: 151, Name: "AudioSamplingRate"},
        152:	&data.NumberEntry{ ID: 152, Name: "AudioSamplingResolution"},
        153:	&data.NumberEntry{ ID: 153, Name: "AudioDuration"},
        154:	&data.StringEntry{ ID: 154, Name: "AudioOutcue"},
        184:	&data.StringEntry{ ID: 184, Name: "JobID"},
        185:	&data.StringEntry{ ID: 185, Name: "MasterDocumentID"},
        186:	&data.StringEntry{ ID: 186, Name: "ShortDocumentID"},
        187:	&data.StringEntry{ ID: 187, Name: "UniqueDocumentID"},
        188:	&data.StringEntry{ ID: 188, Name: "OwnerID"},
        200:	&data.FileFormatEntry{ ID: 200, Name: "ObjectPreviewFileFormat"},
        201:	&data.NumberEntry{ ID: 201, Name: "ObjectPreviewFileVersion"},
        202:	&data.NumberEntry{ ID: 202, Name: "ObjectPreviewData"},
        221:	&data.StringEntry{ ID: 221, Name: "Prefs"},
        225:	&data.StringEntry{ ID: 225, Name: "ClassifyState"},
        228:	&data.StringEntry{ ID: 228, Name: "SimilarityIndex"},
        230:	&data.StringEntry{ ID: 230, Name: "DocumentNotes"},
        231:	&data.StringEntry{ ID: 231, Name: "DocumentHistory"},
        232:	&data.StringEntry{ ID: 232, Name: "ExifCameraInfo"},
        255:	&data.StringEntry{ ID: 255, Name: "CatalogSets"},
}

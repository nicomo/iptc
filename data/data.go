package data

// Entry interfaces define a IPTC tag's value, its name and its id.
type Entry interface {
        SetID(id int)
        SetName(name string)
        SetRaw(raw []string)
        AppendRaw(raw string)
        New()                   Entry
        Value()                 interface{}
        GetID()                 int
        GetName()               string
}

// TagsTable are maps of IPTC tag IDs with their tag names.
type TagsTable map[int]Entry


// IptcRecords maps IPTC record ID types to their name.
var IptcRecords = map[int]TagsTable {
        1:	IptcEnvelopeTags,
        2:      IptcApplicationTags,
}

// IptcEnvelopeTags maps IPTC envelope tag IDs to their tag's name.
// See: http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
var IptcEnvelopeTags = TagsTable {
        0:	&NumberEntry{ ID: 0,   Name: "EnvelopeRecordVersion"},
        5:	&StringEntry{ ID: 5,   Name: "Destination" },
        20:	&NumberEntry{ ID: 20,  Name: "FileFormat"},
        22:	&NumberEntry{ ID: 22,  Name: "FileVersion"},
        30:	&StringEntry{ ID: 30,  Name: "ServiceIdentifier" },
        40:	&NumberEntry{ ID: 40,  Name: "EnvelopeNumber"},
        50:	&StringEntry{ ID: 50,  Name: "ProductID" },
        60:	&NumberEntry{ ID: 60,  Name: "EnvelopePriority"},
        70:	&NumberEntry{ ID: 70,  Name: "DateSent"},
        80:	&StringEntry{ ID: 80,  Name: "TimeSent" },
        90:	&StringEntry{ ID: 90,  Name: "CodedCharacterSet" },
        100:	&StringEntry{ ID: 100, Name: "UniqueObjectName" },
        120:	&NumberEntry{ ID: 120, Name: "ARMIdentifier"},
        122:	&NumberEntry{ ID: 122, Name: "ARMVersion"},
}

// IptcApplicationTags maps IPTC application tag IDs to their tag's name.
// See: http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
var IptcApplicationTags = TagsTable {
        0:      &NumberEntry{ ID: 0, Name: "ApplicationRecordVersion"},
        3:      &StringEntry{ ID: 3, Name: "ObjectTypeReference"},
        4:      &StringEntry{ ID: 4, Name: "ObjectAttributeReference"},
        5:      &StringEntry{ ID: 5, Name: "ObjectName"},
        7:      &StringEntry{ ID: 7, Name: "EditStatus"},
        8:      &NumberEntry{ ID: 8, Name: "EditorialUpdate"},
        10:     &NumberEntry{ ID: 10, Name: "Urgency"},
        12:     &StringEntry{ ID: 12, Name: "SubjectReference"},
        15:     &StringEntry{ ID: 15, Name: "Category"},
        20:     &StringEntry{ ID: 20, Name: "SupplementalCategories"},
        22:     &StringEntry{ ID: 22, Name: "FixtureIdentifier"},
        25:     &ListEntry{ ID: 25, Name: "Keywords"},
        26:     &StringEntry{ ID: 26, Name: "ContentLocationCode"},
        27:     &StringEntry{ ID: 27, Name: "ContentLocationName"},
        30:     &NumberEntry{ ID: 30, Name: "ReleaseDate"},
        35:     &StringEntry{ ID: 35, Name: "ReleaseTime"},
        37:     &NumberEntry{ ID: 37, Name: "ExpirationDate"},
        38:     &StringEntry{ ID: 38, Name: "ExpirationTime"},
        40:     &StringEntry{ ID: 40, Name: "SpecialInstructions"},
        42:     &NumberEntry{ ID: 42, Name: "ActionAdvised"},
        45:     &StringEntry{ ID: 45, Name: "ReferenceService"},
        47:     &NumberEntry{ ID: 47, Name: "ReferenceDate"},
        50:     &NumberEntry{ ID: 50, Name: "ReferenceNumber"},
        55:     &NumberEntry{ ID: 55, Name: "DateCreated"},
        60:     &StringEntry{ ID: 60, Name: "TimeCreated"},
        62:     &NumberEntry{ ID: 62, Name: "DigitalCreationDate"},
        63:     &StringEntry{ ID: 63, Name: "DigitalCreationTime"},
        65:     &StringEntry{ ID: 65, Name: "OriginatingProgram"},
        70:	&StringEntry{ ID: 70, Name: "ProgramVersion"},
        75:	&StringEntry{ ID: 75, Name: "ObjectCycle"},
        80:	&StringEntry{ ID: 80, Name: "By-line"},
        85:	&StringEntry{ ID: 85, Name: "By-lineTitle"},
        90:	&StringEntry{ ID: 90, Name: "City"},
        92:	&StringEntry{ ID: 92, Name: "Sub-location"},
        95:	&StringEntry{ ID: 95, Name: "Province-State"},
        100:	&StringEntry{ ID: 100, Name: "Country-PrimaryLocationCode"},
        101:	&StringEntry{ ID: 101, Name: "Country-PrimaryLocationName"},
        103:	&StringEntry{ ID: 103, Name: "OriginalTransmissionReference"},
        105:	&StringEntry{ ID: 105, Name: "Headline"},
        110:	&StringEntry{ ID: 110, Name: "Credit"},
        115:	&StringEntry{ ID: 115, Name: "Source"},
        116:	&StringEntry{ ID: 116, Name: "CopyrightNotice"},
        118:	&StringEntry{ ID: 118, Name: "Contact"},
        120:	&StringEntry{ ID: 120, Name: "Caption-Abstract"},
        121:	&StringEntry{ ID: 121, Name: "LocalCaption"},
        122:	&StringEntry{ ID: 122, Name: "Writer-Editor"},
        125:	&StringEntry{ ID: 125, Name: "RasterizedCaption"},
        130:	&StringEntry{ ID: 130, Name: "ImageType"},
        131:	&ImageOrientationEntry{ ID: 131, Name: "ImageOrientation"},
        135:	&StringEntry{ ID: 135, Name: "LanguageIdentifier"},
        150:	&AudioTypeEntry{ ID: 150, Name: "AudioType"},
        151:	&NumberEntry{ ID: 151, Name: "AudioSamplingRate"},
        152:	&NumberEntry{ ID: 152, Name: "AudioSamplingResolution"},
        153:	&NumberEntry{ ID: 153, Name: "AudioDuration"},
        154:	&StringEntry{ ID: 154, Name: "AudioOutcue"},
        184:	&StringEntry{ ID: 184, Name: "JobID"},
        185:	&StringEntry{ ID: 185, Name: "MasterDocumentID"},
        186:	&StringEntry{ ID: 186, Name: "ShortDocumentID"},
        187:	&StringEntry{ ID: 187, Name: "UniqueDocumentID"},
        188:	&StringEntry{ ID: 188, Name: "OwnerID"},
        200:	&FileFormatEntry{ ID: 200, Name: "ObjectPreviewFileFormat"},
        201:	&NumberEntry{ ID: 201, Name: "ObjectPreviewFileVersion"},
        202:	&NumberEntry{ ID: 202, Name: "ObjectPreviewData"},
        221:	&StringEntry{ ID: 221, Name: "Prefs"},
        225:	&StringEntry{ ID: 225, Name: "ClassifyState"},
        228:	&StringEntry{ ID: 228, Name: "SimilarityIndex"},
        230:	&StringEntry{ ID: 230, Name: "DocumentNotes"},
        231:	&StringEntry{ ID: 231, Name: "DocumentHistory"},
        232:	&StringEntry{ ID: 232, Name: "ExifCameraInfo"},
        255:	&StringEntry{ ID: 255, Name: "CatalogSets"},
}

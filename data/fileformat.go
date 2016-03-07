package data

// FileFormatEntry indicates the object preview format of a file.
type FileFormatEntry struct {
        ID     int
        Name   string
        Raw  []string
}

var formats = map[string]string {
        "0":      "No ObjectData",
        "1":      "IPTC-NAA Digital Newsphoto Parameter Record",
        "2":      "IPTC7901 Recommended Message Format",
        "3":      "Tagged Image File Format (Adobe/Aldus Image data)",
        "4":      "Illustrator (Adobe Graphics data)",
        "5":      "AppleSingle (Apple Computer Inc)",
        "6":      "NAA 89-3 (ANPA 1312)",
        "7":      "MacBinary II",
        "8":      "IPTC Unstructured Character Oriented File Format (UCOFF)",
        "9":      "United Press International ANPA 1312 variant",
        "10":     "United Press International Down-Load Message",
        "11":     "JPEG File Interchange (JFIF)",
        "12":     "Photo-CD Image-Pac (Eastman Kodak)",
        "13":     "Bit Mapped Graphics File [.BMP] (Microsoft)",
        "14":     "Digital Audio File [.WAV] (Microsoft & Creative Labs)",
        "15":     "Audio plus Moving Video [.AVI] (Microsoft)",
        "16":     "PC DOS/Windows Executable Files [.COM][.EXE]",
        "17":     "Compressed Binary File [.ZIP] (PKWare Inc)",
        "18":     "Audio Interchange File Format AIFF (Apple Computer Inc)",
        "19":     "RIFF Wave (Microsoft Corporation)",
        "20":     "Freehand (Macromedia/Aldus)",
        "21":     "Hypertext Markup Language [.HTML] (The Internet Society)",
        "22":     "MPEG 2 Audio Layer 2 (Musicom), ISO/IEC",
        "23":     "MPEG 2 Audio Layer 3, ISO/IEC",
        "24":     "Portable Document File [.PDF] Adobe",
        "25":     "News Industry Text Format (NITF)",
        "26":     "Tape Archive [.TAR]",
        "27":     "Tidningarnas Telegrambyra NITF version (TTNITF DTD)",
        "28":     "Ritzaus Bureau NITF version (RBNITF DTD)",
        "29":     "Corel Draw [.CDR",
}

// New creates a new copy of a FileFormatEntry.
func (d *FileFormatEntry) New() Entry {
        new := &FileFormatEntry{}
        *new = *d

        return new
}

// SetID changes the tag ID of a FileFormatEntry.
func (d *FileFormatEntry) SetID(id int) {
        d.ID = id
}

// SetName changes the tag name of a FileFormatEntry.
func (d *FileFormatEntry) SetName(name string) {
        d.Name = name
}

// Value returns the FileFormatEntry's value as a single string.
func (d *FileFormatEntry) Value() interface{} {
        value, ok := formats[d.Raw[0]]
        if ok == false {
                return d.Raw[0]
        }

        return value
}

// SetRaw sets the raw data value of a FileFormatEntry.
func (d *FileFormatEntry) SetRaw(raw []string) {
        d.Raw = raw
}

// AppendRaw appends a raw string value to FileFormatEntry.
func (d *FileFormatEntry) AppendRaw(raw string) {
        d.Raw = append(d.Raw, raw)
}

// GetName returns the FileFormatEntry's name.
func (d *FileFormatEntry) GetName() string {
        return d.Name
}

// GetID returns the FileFormatEntry's id.
func (d *FileFormatEntry) GetID() int {
        return d.ID
}

func (d *FileFormatEntry) String() string {
        return d.Value().(string)
}

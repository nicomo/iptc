package data

// ImageOrientationEntry indicates the type of an image file.
type ImageOrientationEntry struct {
        ID     int
        Name   string
        Raw  []string
}

var orientations = map[string]string {
        "L": "Landscape",
        "P": "Portrait",
        "S": "Square",
}

// New creates a new copy of a ImageOrientationEntry.
func (d *ImageOrientationEntry) New() Entry {
        new := &ImageOrientationEntry{}
        *new = *d

        return new
}


// SetID changes the tag ID of a ImageOrientationEntry.
func (d *ImageOrientationEntry) SetID(id int) {
        d.ID = id
}

// SetName changes the tag name of a ImageOrientationEntry.
func (d *ImageOrientationEntry) SetName(name string) {
        d.Name = name
}

// Value returns the ImageOrientationEntry's value as a single string.
func (d *ImageOrientationEntry) Value() interface{} {
        value, ok := orientations[d.Raw[0]]
        if ok == false {
                return d.Raw[0]
        }

        return value
}

// SetRaw sets the raw data value of a ImageOrientationEntry.
func (d *ImageOrientationEntry) SetRaw(raw []string) {
        d.Raw = raw
}

// AppendRaw appends a raw string value to ImageOrientationEntry.
func (d *ImageOrientationEntry) AppendRaw(raw string) {
        d.Raw = append(d.Raw, raw)
}

// GetName returns the ImageOrientationEntry's name.
func (d *ImageOrientationEntry) GetName() string {
        return d.Name
}

// GetID returns the ImageOrientationEntry's id.
func (d *ImageOrientationEntry) GetID() int {
        return d.ID
}

package data

import "strings"

// StringEntry contains the parsed value of an IPTC tag.
type StringEntry struct {
        ID     int
        Name   string
        Raw  []string
}

// New creates a new copy of a StringEntry.
func (d *StringEntry) New() Entry {
        new := &StringEntry{}
        *new = *d

        return new
}

// SetID changes the tag ID of a StringEntry.
func (d *StringEntry) SetID(id int) {
        d.ID = id
}

// SetName changes the tag name of a StringEntry.
func (d *StringEntry) SetName(name string) {
        d.Name = name
}

// Value returns the StringEntry's value as a single string.
func (d *StringEntry) Value() interface{} {
        return strings.Join(d.Raw, " ")
}

// SetRaw sets the raw data value of a StringEntry.
func (d *StringEntry) SetRaw(raw []string) {
        d.Raw = raw
}

// AppendRaw appends a raw string value to StringEntry.
func (d *StringEntry) AppendRaw(raw string) {
        d.Raw = append(d.Raw, raw)
}

// GetName returns the StringEntry's name.
func (d *StringEntry) GetName() string {
        return d.Name
}

// GetID returns the StringEntry's id.
func (d *StringEntry) GetID() int {
        return d.ID
}

func (d *StringEntry) String() string {
        return strings.Join(d.Raw, " ")
}

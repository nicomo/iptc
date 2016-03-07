package data

import "strings"

// ListEntry contains the parsed value of an IPTC tag.
type ListEntry struct {
        ID     int
        Name   string
        Raw  []string
}

// New creates a new copy of a ListEntry.
func (d *ListEntry) New() Entry {
        new := &ListEntry{}
        *new = *d

        return new
}

// SetID changes the tag ID of a ListEntry.
func (d *ListEntry) SetID(id int) {
        d.ID = id
}

// SetName changes the tag name of a ListEntry.
func (d *ListEntry) SetName(name string) {
        d.Name = name
}

// Value returns the ListEntry's value as a single string.
func (d *ListEntry) Value() interface{} {
        return d.Raw
}

// SetRaw sets the raw data value of a ListEntry.
func (d *ListEntry) SetRaw(raw []string) {
        d.Raw = raw
}

// AppendRaw appends a raw string value to ListEntry.
func (d *ListEntry) AppendRaw(raw string) {
        d.Raw = append(d.Raw, raw)
}

// GetName returns the ListEntry's name.
func (d *ListEntry) GetName() string {
        return d.Name
}

// GetID returns the ListEntry's id.
func (d *ListEntry) GetID() int {
        return d.ID
}

func (d *ListEntry) String() string {
        return strings.Join(d.Value().([]string), " ")
}

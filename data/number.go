package data

import "strconv"

// NumberEntry contains the parsed value of an IPTC tag.
type NumberEntry struct {
        ID     int
        Name   string
        Raw  []string
}

// New creates a new copy of a NumberEntry.
func (d *NumberEntry) New() Entry {
        new := &NumberEntry{}
        *new = *d

        return new
}

// SetID changes the tag ID of a NumberEntry.
func (d *NumberEntry) SetID(id int) {
        d.ID = id
}

// SetName changes the tag name of a NumberEntry.
func (d *NumberEntry) SetName(name string) {
        d.Name = name
}

// Value returns the NumberEntry's value as a single string.
func (d *NumberEntry) Value() interface{} {
        value, err := strconv.Atoi(d.Raw[0])
        if err != nil {
                value = 0
        }

        return value
}

// SetRaw sets the raw data value of a NumberEntry.
func (d *NumberEntry) SetRaw(raw []string) {
        d.Raw = raw
}

// AppendRaw appends a raw string value to NumberEntry.
func (d *NumberEntry) AppendRaw(raw string) {
        d.Raw = append(d.Raw, raw)
}

// GetName returns the NumberEntry's name.
func (d *NumberEntry) GetName() string {
        return d.Name
}

// GetID returns the NumberEntry's id.
func (d *NumberEntry) GetID() int {
        return d.ID
}

package data

// AudioTypeEntry indicates the type of audio of a file.
type AudioTypeEntry struct {
        ID     int
        Name   string
        Raw  []string
}

var types = map[string]string {
        "0T": "Text Only",
        "1A": "Mono Actuality",
        "1C": "Mono Question and Answer Session",
        "1M": "Mono Music",
        "1Q": "Mono Response to a Question",
        "1R": "Mono Raw Sound",
        "1S": "Mono Scener",
        "1V": "Mono Voicer",
        "1W": "Mono Wrap",
        "2A": "Stereo Actuality",
        "2C": "Stereo Question and Answer Session",
        "2M": "Stereo Music",
        "2Q": "Stereo Response to a Question",
        "2R": "Stereo Raw Sound",
        "2S": "Stereo Scener",
        "2V": "Stereo Voicer",
        "2W": "Stereo Wra",
}

// New creates a new copy of a AudioTypeEntry.
func (d *AudioTypeEntry) New() Entry {
        new := &AudioTypeEntry{}
        *new = *d

        return new
}

// SetID changes the tag ID of a AudioTypeEntry.
func (d *AudioTypeEntry) SetID(id int) {
        d.ID = id
}

// SetName changes the tag name of a AudioTypeEntry.
func (d *AudioTypeEntry) SetName(name string) {
        d.Name = name
}

// Value returns the AudioTypeEntry's value as a single string.
func (d *AudioTypeEntry) Value() interface{} {
        value, ok := types[d.Raw[0]]
        if ok == false {
                return d.Raw[0]
        }

        return value
}

// SetRaw sets the raw data value of a AudioTypeEntry.
func (d *AudioTypeEntry) SetRaw(raw []string) {
        d.Raw = raw
}

// AppendRaw appends a raw string value to AudioTypeEntry.
func (d *AudioTypeEntry) AppendRaw(raw string) {
        d.Raw = append(d.Raw, raw)
}

// GetName returns the AudioTypeEntry's name.
func (d *AudioTypeEntry) GetName() string {
        return d.Name
}

// GetID returns the AudioTypeEntry's id.
func (d *AudioTypeEntry) GetID() int {
        return d.ID
}

func (d *AudioTypeEntry) String() string {
        return d.Value().(string)
}

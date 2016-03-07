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

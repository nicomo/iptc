package iptc

import "errors"

// ErrNoIptcData errors indicate files without any IPTC data.
var ErrNoIptcData = errors.New("no IPTC data in file %s")

// ErrIptcTagNotFound errors indicate that the specified tag name does not exist for a file.
var ErrIptcTagNotFound = errors.New("no IPTC data tag '%s' found in file %s")

// ErrIptcTagType errors indicate that the wrong Tag() function was used to access its value.
var ErrIptcTagType = errors.New("wrong IPTC tag type spcified for '%s' in file %s")

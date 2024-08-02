package rfc6929

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"layeh.com/radius"
)

const (
	VendorSpecificAttribute byte = 26
)

func NewExtendedVendorSpecific(vendorId uint32, typ byte, extType radius.Type, attr radius.Attribute) (radius.Attribute, error) {
	if extType < 241 || extType > 246 {
		return nil, fmt.Errorf("Radius type %d is not an extension type", extType)
	}
	evsa := make(radius.Attribute, 6+len(attr))
	evsa[0] = VendorSpecificAttribute
	binary.BigEndian.PutUint32(evsa[1:5], vendorId)
	evsa[5] = typ
	copy(evsa[6:], attr)
	return evsa, nil
}

func GetExtendedVendorSpecific(extAttr radius.Attribute) (vendorID uint32, typ byte, attr radius.Attribute, err error) {
	if len(extAttr) < 7 {
		err = fmt.Errorf("Attribute is only %d bytes - must be at least 7", len(extAttr))
		return
	}
	vendorBuf := bytes.NewReader(extAttr[1:5])
	err = binary.Read(vendorBuf, binary.BigEndian, &vendorID)
	if extAttr[0] != VendorSpecificAttribute {
		err = fmt.Errorf("Radius type %d is not valid for a vendor-specific attribute", extAttr[0])
		return
	}
	typ = extAttr[5]
	attr = make(radius.Attribute, len(extAttr)-6)
	copy(attr, extAttr[6:])
	return
}

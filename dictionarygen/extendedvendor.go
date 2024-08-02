package dictionarygen

import (
	"io"
	"layeh.com/radius/dictionary"
)

func DictExtType(extType dictionary.AttributeType) string {
	switch extType {
	case dictionary.AttributeExtendedVSA1:
		return "rfc6929.ExtendedVendorSpecific1_Type"
	case dictionary.AttributeExtendedVSA2:
		return "rfc6929.ExtendedVendorSpecific2_Type"
	case dictionary.AttributeExtendedVSA3:
		return "rfc6929.ExtendedVendorSpecific3_Type"
	case dictionary.AttributeExtendedVSA4:
		return "rfc6929.ExtendedVendorSpecific4_Type"
	case dictionary.AttributeLongExtendedVSA5:
		return "rfc6929.ExtendedVendorSpecific5_Type"
	case dictionary.AttributeLongExtendedVSA6:
		return "rfc6929.ExtendedVendorSpecific6_Type"
	case dictionary.AttributeExtended1:
		return "rfc6929.ExtendedAttribute1_Type"
	case dictionary.AttributeExtended2:
		return "rfc6929.ExtendedAttribute2_Type"
	case dictionary.AttributeExtended3:
		return "rfc6929.ExtendedAttribute3_Type"
	case dictionary.AttributeExtended4:
		return "rfc6929.ExtendedAttribute4_Type"
	case dictionary.AttributeLongExtended5:
		return "rfc6929.ExtendedAttribute5_Type"
	case dictionary.AttributeLongExtended6:
		return "rfc6929.ExtendedAttribute6_Type"
	}
	return ""
}

func (g *Generator) genExtVendor(w io.Writer, vendor *dictionary.Vendor) {
	ident := identifier(vendor.Name)

	p(w)
	p(w, `func _`, ident, `_AddExtVendor(p *radius.Packet, typ byte, extType radius.Type, attr radius.Attribute) (err error) {`)
	p(w, `	evsa, err := rfc6929.NewExtendedVendorSpecific(_`, ident, `_VendorID, typ, extType, attr)`)
	p(w, `	if err != nil {`)
	p(w, `		return err`)
	p(w, `	}`)
	p(w, `	p.Add(extType, evsa)`)
	p(w, `	return`)
	p(w, `}`)
	p(w)

	p(w, `func _`, ident, `_GetsExtVendor(p *radius.Packet, typ byte, extType radius.Type) (values []radius.Attribute) {`)
	p(w, `	for _, avp := range p.Attributes {`)
	p(w, `		if avp.Type != extType {`)
	p(w, `			continue`)
	p(w, `		}`)
	p(w, `		vendorID, vsaType, attr, err := rfc6929.GetExtendedVendorSpecific(avp.Attribute)`)
	p(w, `		if err != nil || vendorID != _`, ident, `_VendorID || vsaType != typ {`)
	p(w, `			continue`)
	p(w, `		}`)
	p(w, `		values = append(values, attr)`)
	p(w, `	}`)
	p(w, `	return`)
	p(w, `}`)

	p(w)
	p(w, `func _`, ident, `_LookupExtVendor(p *radius.Packet, typ byte, extType radius.Type) (radius.Attribute, bool) {`)
	p(w, `	for _, avp := range p.Attributes {`)
	p(w, `		if avp.Type != extType {`)
	p(w, `			continue`)
	p(w, `		}`)
	p(w, `		vendorID, vsaType, attr, err := rfc6929.GetExtendedVendorSpecific(avp.Attribute)`)
	p(w, `		if err != nil || vendorID != _`, ident, `_VendorID || vsaType != typ {`)
	p(w, `			continue`)
	p(w, `		}`)
	p(w, `		return attr, true`)
	p(w, `	}`)
	p(w, `	return nil, false`)
	p(w, `}`)
	p(w)

	p(w, `func _`, ident, `_SetExtVendor(p *radius.Packet, typ byte, extType radius.Type, attr radius.Attribute) (err error) {`)
	p(w, `	_AlcatelIPD_DelExtVendor(p, typ, extType)`)
	p(w, `	return _AlcatelIPD_AddExtVendor(p, typ, extType, attr)`)
	p(w, `}`)
	p(w)

	p(w, `func _`, ident, `_DelExtVendor(p *radius.Packet, typ byte, extType radius.Type) {`)
	p(w, `	var deleted int`)
	p(w, `	for i, avp := range p.Attributes {`)
	p(w, `		if avp.Type != extType {`)
	p(w, `			continue`)
	p(w, `		}`)
	p(w, `		vendorID, vsaType, _, err := rfc6929.GetExtendedVendorSpecific(avp.Attribute)`)
	p(w, `		if err != nil || vendorID != _AlcatelIPD_VendorID || vsaType != typ {`)
	p(w, `			continue`)
	p(w, `		}`)
	p(w, `		p.Attributes = append(p.Attributes[:i-deleted], p.Attributes[i+1-deleted:]...)`)
	p(w, `		deleted++`)
	p(w, `	}`)
	p(w, `}`)
}

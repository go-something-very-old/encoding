package dat

import (
	"encoding/xml"

	bin "github.com/go-otserv/encoding/binary"
)

// Thing holds thing information: ID, type, attributes and sprites information
type Thing struct {
	XMLName      xml.Name       `xml:"thing"`
	ID           uint16         `xml:"id,attr"`
	Type         string         `xml:"type,attr"`
	Attributes   []Attribute    `xml:"Attribute,omitempty"`
	SpriteGroups []*SpriteGroup `xml:"-"`
}

// Some docstring to those constants
const (
	ITEM    = "item"
	OUTFIT  = "outfit"
	EFFECT  = "effect"
	MISSILE = "missile"
)

// NewThing creates new instance of Thing
func NewThing(id uint16, typ string) *Thing {
	return &Thing{ID: id, Type: typ}
}

// DeserializeThing parses .dat file and creates new Thing instance
func DeserializeThing(id uint16, typ string, datfh bin.Reader) (*Thing, error) {
	var err error
	thing := NewThing(id, typ)
	if err = thing.deserializeAttributes(datfh); err != nil {
		return thing, err
	}
	if err = thing.deserializeSpritesInfo(datfh); err != nil {
		return thing, err
	}
	return thing, nil
}

func (thing *Thing) deserializeAttributes(datfh bin.Reader) error {
	var err error
	var attrOp uint8
	var attr Attribute

	thing.Attributes = make([]Attribute, 0, 5)

	for attrOp != 255 {
		if attrOp, err = datfh.UInt8(); err != nil {
			return err
		}
		if attr, err = deserializeAttribute(attrOp, datfh); err != nil {
			return err
		}
		thing.Attributes = append(thing.Attributes, attr)
	}
	return nil
}

func (thing *Thing) deserializeSpritesInfo(datfh bin.Reader) error {
	var err error
	var groupsCount uint8
	var sprGr *SpriteGroup

	if thing.Type == OUTFIT {
		if groupsCount, err = datfh.UInt8(); err != nil {
			return err
		}
	} else {
		groupsCount = 1
	}

	thing.SpriteGroups = make([]*SpriteGroup, 0, groupsCount)

	for group := 1; group <= int(groupsCount); group++ {
		if sprGr, err = deserializeSpriteGroup(thing.Type, datfh); err != nil {
			return err
		}
		sprGr.Group = group
		thing.SpriteGroups = append(thing.SpriteGroups, sprGr)
	}
	return nil
}

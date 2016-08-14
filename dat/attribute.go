package dat

import (
	"encoding/xml"
	"fmt"

	bin "github.com/go-otserv/encoding/binary"
)

// Attribute is common interface for all Item attributes
type Attribute interface {
	attrName() string
}

// AttributeBase is base struct for all Item attributes
type AttributeBase struct {
	XMLName xml.Name `xml:"attr"`
	Name    string   `xml:"name,attr"`
}

// Ground attribute
// OpCode: 0
type Ground struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr Ground) attrName() string { return attr.Name }

// NewGround creates new Ground attribute
func NewGround(val uint16) *Ground {
	return &Ground{AttributeBase{Name: "ground"}, val}
}

// GroundBorder attribute
// OpCode: 1
type GroundBorder struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr GroundBorder) attrName() string { return attr.Name }

// NewGroundBorder creates new GroundBorder attribute
func NewGroundBorder() *GroundBorder {
	return &GroundBorder{}
}

// OnBottom attribute
// OpCode: 2
type OnBottom struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr OnBottom) attrName() string { return attr.Name }

// NewOnBottom creates new OnBottom attribute
func NewOnBottom() *OnBottom {
	return &OnBottom{AttributeBase{Name: "onBottom"}}
}

// OnTop attribute
// OpCode: 3
type OnTop struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr OnTop) attrName() string { return attr.Name }

// NewOnTop creates new OnTop attribute
func NewOnTop() *OnTop {
	return &OnTop{AttributeBase{Name: "onTop"}}
}

// Container attribute
// OpCode: 4
type Container struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Container) attrName() string { return attr.Name }

// NewContainer creates new Container attribute
func NewContainer() *Container {
	return &Container{AttributeBase{Name: "container"}}
}

// Stackable attribute
// OpCode: 5
type Stackable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Stackable) attrName() string { return attr.Name }

// NewStackable creates new Stackable attribute
func NewStackable() *Stackable {
	return &Stackable{AttributeBase{Name: "stackable"}}
}

// ForceUse attribute
// OpCode: 6
type ForceUse struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr ForceUse) attrName() string { return attr.Name }

// NewForceUse creates new ForceUse attribute
func NewForceUse() *ForceUse {
	return &ForceUse{AttributeBase{Name: "forceUse"}}
}

// MultiUse attribute
// OpCode: 7
type MultiUse struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr MultiUse) attrName() string { return attr.Name }

// NewMultiUse creates new MultiUse attribute
func NewMultiUse() *MultiUse {
	return &MultiUse{AttributeBase{Name: "multiUse"}}
}

// Writable attribute
// OpCode: 8
type Writable struct {
	AttributeBase
	TextLen uint16 `xml:"textLen,attr"`
}

// attrName implements Attribute interface
func (attr Writable) attrName() string { return attr.Name }

// NewWritable creates new Writable attribute
func NewWritable(val uint16) *Writable {
	return &Writable{AttributeBase{Name: "writable"}, val}
}

// WritableOnce attribute
// OpCode: 9
type WritableOnce struct {
	AttributeBase
	TextLen uint16 `xml:"textLen,attr"`
}

// attrName implements Attribute interface
func (attr WritableOnce) attrName() string { return attr.Name }

// NewWritableOnce creates new WritableOnce attribute
func NewWritableOnce(val uint16) *WritableOnce {
	return &WritableOnce{AttributeBase{Name: "writableOnce"}, val}
}

// FluidContainer attribute
// OpCode: 10
type FluidContainer struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr FluidContainer) attrName() string { return attr.Name }

// NewFluidContainer creates new FluidContainer attribute
func NewFluidContainer() *FluidContainer {
	return &FluidContainer{AttributeBase{Name: "fluidContainer"}}
}

// Splash attribute
// OpCode: 11
type Splash struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Splash) attrName() string { return attr.Name }

// NewSplash creates new Splash attribute
func NewSplash() *Splash {
	return &Splash{AttributeBase{Name: "splash"}}
}

// NotWalkable attribute
// OpCode: 12
type NotWalkable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr NotWalkable) attrName() string { return attr.Name }

// NewNotWalkable creates new NotWalkable attribute
func NewNotWalkable() *NotWalkable {
	return &NotWalkable{AttributeBase{Name: "notWalkable"}}
}

// NotMoveable attribute
// OpCode: 13
type NotMoveable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr NotMoveable) attrName() string { return attr.Name }

// NewNotMoveable creates new NotMoveable attribute
func NewNotMoveable() *NotMoveable {
	return &NotMoveable{AttributeBase{Name: "notMoveable"}}
}

// BlockProjectile attribute
// OpCode: 14
type BlockProjectile struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr BlockProjectile) attrName() string { return attr.Name }

// NewBlockProjectile creates new BlockProjectile attribute
func NewBlockProjectile() *BlockProjectile {
	return &BlockProjectile{AttributeBase{Name: "blockProjectile"}}
}

// NotPathable attribute
// OpCode: 15
type NotPathable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr NotPathable) attrName() string { return attr.Name }

// NewNotPathable creates new NotPathable attribute
func NewNotPathable() *NotPathable {
	return &NotPathable{AttributeBase{Name: "notPathable"}}
}

// NoMoveAnimation attribute
// OpCode: 16
type NoMoveAnimation struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr NoMoveAnimation) attrName() string { return attr.Name }

// NewNoMoveAnimation creates new NoMoveAnimation attribute
func NewNoMoveAnimation() *NoMoveAnimation {
	return &NoMoveAnimation{AttributeBase{Name: "noMoveAnimation"}}
}

// Pickupable attribute
// OpCode: 17
type Pickupable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Pickupable) attrName() string { return attr.Name }

// NewPickupable creates new Pickupable attribute
func NewPickupable() *Pickupable {
	return &Pickupable{AttributeBase{Name: "pickupable"}}
}

// Hangable attribute
// OpCode: 18
type Hangable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Hangable) attrName() string { return attr.Name }

// NewHangable creates new Hangable attribute
func NewHangable() *Hangable {
	return &Hangable{AttributeBase{Name: "hangable"}}
}

// HookSouth attribute
// OpCode: 19
type HookSouth struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr HookSouth) attrName() string { return attr.Name }

// NewHookSouth creates new HookSouth attribute
func NewHookSouth() *HookSouth {
	return &HookSouth{AttributeBase{Name: "hookSouth"}}
}

// HookEast attribute
// OpCode: 20
type HookEast struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr HookEast) attrName() string { return attr.Name }

// NewHookEast creates new HookEast attribute
func NewHookEast() *HookEast {
	return &HookEast{AttributeBase{Name: "hookEast"}}
}

// Rotateable attribute
// OpCode: 21
type Rotateable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Rotateable) attrName() string { return attr.Name }

// NewRotateable creates new Rotateable attribute
func NewRotateable() *Rotateable {
	return &Rotateable{AttributeBase{Name: "rotateable"}}
}

// Light attribute
// OpCode: 22
type Light struct {
	AttributeBase
	Intensity uint16 `xml:"intensity,attr"`
	Color     uint16 `xml:"color,attr"`
}

// attrName implements Attribute interface
func (attr Light) attrName() string { return attr.Name }

// NewLight creates new Light attribute
func NewLight(intensity, color uint16) *Light {
	return &Light{AttributeBase{Name: "light"}, intensity, color}
}

// DontHide attribute
// OpCode: 23
type DontHide struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr DontHide) attrName() string { return attr.Name }

// NewDontHide creates new DontHide attribute
func NewDontHide() *DontHide {
	return &DontHide{AttributeBase{Name: "dontHide"}}
}

// Translucent attribute
// OpCode: 24
type Translucent struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Translucent) attrName() string { return attr.Name }

// NewTranslucent creates new Translucent attribute
func NewTranslucent() *Translucent {
	return &Translucent{AttributeBase{Name: "translucent"}}
}

// Displacement attribute
// OpCode: 25
type Displacement struct {
	AttributeBase
	X uint16 `xml:"x,attr"`
	Y uint16 `xml:"y,attr"`
}

// attrName implements Attribute interface
func (attr Displacement) attrName() string { return attr.Name }

// NewDisplacement creates new Displacement attribute
func NewDisplacement(x, y uint16) *Displacement {
	return &Displacement{AttributeBase{Name: "displacement"}, x, y}
}

// Elevation attribute
// OpCode: 26
type Elevation struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr Elevation) attrName() string { return attr.Name }

// NewElevation creates new Elevation attribute
func NewElevation(val uint16) *Elevation {
	return &Elevation{AttributeBase{Name: "elevation"}, val}
}

// LyingCorpse attribute
// OpCode: 27
type LyingCorpse struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr LyingCorpse) attrName() string { return attr.Name }

// NewLyingCorpse creates new LyingCorpse attribute
func NewLyingCorpse() *LyingCorpse {
	return &LyingCorpse{AttributeBase{Name: "lyingCorpse"}}
}

// AnimateAlways attribute
// OpCode: 28
type AnimateAlways struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr AnimateAlways) attrName() string { return attr.Name }

// NewAnimateAlways creates new AnimateAlways attribute
func NewAnimateAlways() *AnimateAlways {
	return &AnimateAlways{AttributeBase{Name: "animateAlways"}}
}

// MinimapColor attribute
// OpCode: 29
type MinimapColor struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr MinimapColor) attrName() string { return attr.Name }

// NewMinimapColor creates new MinimapColor attribute
func NewMinimapColor(val uint16) *MinimapColor {
	return &MinimapColor{AttributeBase{Name: "minimapColor"}, val}
}

// LensHelp attribute
// OpCode: 30
type LensHelp struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr LensHelp) attrName() string { return attr.Name }

// NewLensHelp creates new LensHelp attribute
func NewLensHelp(val uint16) *LensHelp {
	return &LensHelp{AttributeBase{Name: "lensHelp"}, val}
}

// FullGround attribute
// OpCode: 31
type FullGround struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr FullGround) attrName() string { return attr.Name }

// NewFullGround creates new FullGround attribute
func NewFullGround() *FullGround {
	return &FullGround{AttributeBase{Name: "fullGround"}}
}

// Look attribute
// OpCode: 32
type Look struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Look) attrName() string { return attr.Name }

// NewLook creates new Look attribute
func NewLook() *Look {
	return &Look{AttributeBase{Name: "look"}}
}

// Cloth attribute
// OpCode: 33
// Slots:
// * 0 - both hands
// * 1 - head
// * 2 - necklace
// * 3 - backpack
// * 4 - armor
// * 5 - shield
// * 6 - weapon
// * 7 - legs
// * 8 - feet
// * 9 - ring
// * 10 - ammo
// * 11 - store inbox
type Cloth struct {
	AttributeBase
	Slot uint16 `xml:"slot,attr"`
}

// attrName implements Attribute interface
func (attr Cloth) attrName() string { return attr.Name }

// NewCloth creates new Cloth attribute
func NewCloth(slot uint16) *Cloth {
	return &Cloth{AttributeBase{Name: "cloth"}, slot}
}

// Market attribute
// OpCode: 34
type Market struct {
	AttributeBase
	Category         uint16 `xml:"category,attr"`
	TradeAs          uint16 `xml:"tradeAs,attr"`
	ShowAs           uint16 `xml:"showAs,attr"`
	ItemName         string `xml:"itemName,attr"`
	RestrictVocation uint16 `xml:"restrictVocation,attr"`
	RequiredLevel    uint16 `xml:"requiredLevel,attr"`
}

// attrName implements Attribute interface
func (attr Market) attrName() string { return attr.Name }

// NewMarket creates new Market attribute
func NewMarket(category, tradeAs, showAs uint16, itemName string,
	restrictVocation, requiredLevel uint16) *Market {
	return &Market{
		AttributeBase{Name: "market"},
		category,
		tradeAs,
		showAs,
		itemName,
		restrictVocation,
		requiredLevel,
	}
}

// Usable attribute
// OpCode: 35
type Usable struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr Usable) attrName() string { return attr.Name }

// NewUsable creates new Usable attribute
func NewUsable(val uint16) *Usable {
	return &Usable{AttributeBase{Name: "usable"}, val}
}

// Wrapable attribute
// OpCode: 36
type Wrapable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Wrapable) attrName() string { return attr.Name }

// NewWrapable creates new Wrapable attribute
func NewWrapable() *Wrapable {
	return &Wrapable{AttributeBase{Name: "wrapable"}}
}

// Unwrapable attribute
// OpCode: 37
type Unwrapable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Unwrapable) attrName() string { return attr.Name }

// NewUnwrapable creates new Unwrapable attribute
func NewUnwrapable() *Unwrapable {
	return &Unwrapable{AttributeBase{Name: "unwrapable"}}
}

// TopEffect attribute
// OpCode: 38
type TopEffect struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr TopEffect) attrName() string { return attr.Name }

// NewTopEffect creates new TopEffect attribute
func NewTopEffect() *TopEffect {
	return &TopEffect{AttributeBase{Name: "topEffect"}}
}

// Opacity attribute
// OpCode: 100
type Opacity struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr Opacity) attrName() string { return attr.Name }

// NewOpacity creates new Opacity attribute
func NewOpacity() *Opacity {
	return &Opacity{AttributeBase{Name: "opacity"}}
}

// NotPrewalkable attribute
// OpCode: 101
type NotPrewalkable struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr NotPrewalkable) attrName() string { return attr.Name }

// NewNotPrewalkable creates new NotPrewalkable attribute
func NewNotPrewalkable() *NotPrewalkable {
	return &NotPrewalkable{AttributeBase{Name: "notPrewalkable"}}
}

// FloorChange attribute
// OpCode: 252
type FloorChange struct {
	AttributeBase
}

// attrName implements Attribute interface
func (attr FloorChange) attrName() string { return attr.Name }

// NewFloorChange creates new FloorChange attribute
func NewFloorChange() *FloorChange {
	return &FloorChange{AttributeBase{Name: "floorChange"}}
}

// Deprecated attribute
// Opcode: 254
// It doesn't inherit from AttributeBase not to be XML serializable
type Deprecated struct{}

// attrName implements Attribute interface
func (attr Deprecated) attrName() string { return "Deprecated<254>" }

// NewDeprecated creates new Deprecated attribute
func NewDeprecated() *Deprecated {
	return &Deprecated{}
}

// DeserializeAttribute reads from `datfh` and creates proper attribute based
// on given `attrOp`
func deserializeAttribute(attrOp uint8, datfh bin.BinaryReader) (Attribute, error) {
	var err error
	var val, val2 uint16

	switch attrOp {
	case 0:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewGround(val), nil
	case 1:
		return NewGroundBorder(), nil
	case 2:
		return NewOnBottom(), nil
	case 3:
		return NewOnTop(), nil
	case 4:
		return NewContainer(), nil
	case 5:
		return NewStackable(), nil
	case 6:
		return NewForceUse(), nil
	case 7:
		return NewMultiUse(), nil
	case 8:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewWritable(val), nil
	case 9:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewWritableOnce(val), nil
	case 10:
		return NewFluidContainer(), nil
	case 11:
		return NewSplash(), nil
	case 12:
		return NewNotWalkable(), nil
	case 13:
		return NewNotMoveable(), nil
	case 14:
		return NewBlockProjectile(), nil
	case 15:
		return NewNotPathable(), nil
	case 16:
		return NewNoMoveAnimation(), nil
	case 17:
		return NewPickupable(), nil
	case 18:
		return NewHangable(), nil
	case 19:
		return NewHookSouth(), nil
	case 20:
		return NewHookEast(), nil
	case 21:
		return NewRotateable(), nil
	case 22:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		if val2, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewLight(val, val2), nil
	case 23:
		return NewDontHide(), nil
	case 24:
		return NewTranslucent(), nil
	case 25:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		if val2, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewDisplacement(val, val2), nil
	case 26:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewElevation(val), nil
	case 27:
		return NewLyingCorpse(), nil
	case 28:
		return NewAnimateAlways(), nil
	case 29:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewMinimapColor(val), nil
	case 30:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewLensHelp(val), nil
	case 31:
		return NewFullGround(), nil
	case 32:
		return NewLook(), nil
	case 33:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewCloth(val), nil
	case 34:
		var category, tradeAs, showAs, restrictVocation, requiredLevel uint16
		var itemName string

		if category, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		if tradeAs, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		if showAs, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		if itemName, err = datfh.String(); err != nil {
			return nil, err
		}
		if restrictVocation, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		if requiredLevel, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewMarket(
			category,
			tradeAs,
			showAs,
			itemName,
			restrictVocation,
			requiredLevel,
		), nil
	case 35:
		if val, err = datfh.UInt16(); err != nil {
			return nil, err
		}
		return NewUsable(val), nil
	case 36:
		return NewWrapable(), nil
	case 37:
		return NewUnwrapable(), nil
	case 38:
		return NewTopEffect(), nil
	case 100:
		return NewOpacity(), nil
	case 101:
		return NewNotPrewalkable(), nil
	case 252:
		return NewFloorChange(), nil
	case 254:
		return NewDeprecated(), nil
	case 255:
		break
	default:
		return nil, fmt.Errorf("Unknown attribute opcode: %d", attrOp)
	}
	return nil, nil
}

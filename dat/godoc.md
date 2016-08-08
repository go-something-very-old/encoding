# Package `dat`
## Overview

## Index

* Constants

* Types
  * [AnimateAlways](#AnimateAlways)
	 * [func NewAnimateAlways() *AnimateAlways](#NewAnimateAlways)

  * [AnimationPhase](#AnimationPhase)

  * [Attribute](#Attribute)

  * [AttributeBase](#AttributeBase)

  * [BlockProjectile](#BlockProjectile)
	 * [func NewBlockProjectile() *BlockProjectile](#NewBlockProjectile)

  * [Cloth](#Cloth)
	 * [func NewCloth(slot uint16) *Cloth](#NewCloth)

  * [Container](#Container)
	 * [func NewContainer() *Container](#NewContainer)

  * [Deprecated](#Deprecated)
	 * [func NewDeprecated() *Deprecated](#NewDeprecated)

  * [Displacement](#Displacement)
	 * [func NewDisplacement(x, y uint16) *Displacement](#NewDisplacement)

  * [DontHide](#DontHide)
	 * [func NewDontHide() *DontHide](#NewDontHide)

  * [Elevation](#Elevation)
	 * [func NewElevation(val uint16) *Elevation](#NewElevation)

  * [File](#File)
	 * [func Open(path string) (*File, error)](#Open)
	 * [func (datfh *File) AppendThing(thing *Thing)](#File-AppendThing)
	 * [func (datfh *File) Deserialize() error](#File-Deserialize)
	 * [func (datfh *File) DeserializeWithProgress(prChan chan<- int, errChan chan<- error, doneChan chan<- bool)](#File-DeserializeWithProgress)

  * [FloorChange](#FloorChange)
	 * [func NewFloorChange() *FloorChange](#NewFloorChange)

  * [FluidContainer](#FluidContainer)
	 * [func NewFluidContainer() *FluidContainer](#NewFluidContainer)

  * [ForceUse](#ForceUse)
	 * [func NewForceUse() *ForceUse](#NewForceUse)

  * [FullGround](#FullGround)
	 * [func NewFullGround() *FullGround](#NewFullGround)

  * [Ground](#Ground)
	 * [func NewGround(val uint16) *Ground](#NewGround)

  * [GroundBorder](#GroundBorder)
	 * [func NewGroundBorder() *GroundBorder](#NewGroundBorder)

  * [Hangable](#Hangable)
	 * [func NewHangable() *Hangable](#NewHangable)

  * [HookEast](#HookEast)
	 * [func NewHookEast() *HookEast](#NewHookEast)

  * [HookSouth](#HookSouth)
	 * [func NewHookSouth() *HookSouth](#NewHookSouth)

  * [LensHelp](#LensHelp)
	 * [func NewLensHelp(val uint16) *LensHelp](#NewLensHelp)

  * [Light](#Light)
	 * [func NewLight(intensity, color uint16) *Light](#NewLight)

  * [Look](#Look)
	 * [func NewLook() *Look](#NewLook)

  * [LyingCorpse](#LyingCorpse)
	 * [func NewLyingCorpse() *LyingCorpse](#NewLyingCorpse)

  * [Market](#Market)
	 * [func NewMarket(category, tradeAs, showAs uint16, itemName string,
	restrictVocation, requiredLevel uint16) *Market](#NewMarket)

  * [MinimapColor](#MinimapColor)
	 * [func NewMinimapColor(val uint16) *MinimapColor](#NewMinimapColor)

  * [MultiUse](#MultiUse)
	 * [func NewMultiUse() *MultiUse](#NewMultiUse)

  * [NoMoveAnimation](#NoMoveAnimation)
	 * [func NewNoMoveAnimation() *NoMoveAnimation](#NewNoMoveAnimation)

  * [NotMoveable](#NotMoveable)
	 * [func NewNotMoveable() *NotMoveable](#NewNotMoveable)

  * [NotPathable](#NotPathable)
	 * [func NewNotPathable() *NotPathable](#NewNotPathable)

  * [NotPrewalkable](#NotPrewalkable)
	 * [func NewNotPrewalkable() *NotPrewalkable](#NewNotPrewalkable)

  * [NotWalkable](#NotWalkable)
	 * [func NewNotWalkable() *NotWalkable](#NewNotWalkable)

  * [OnBottom](#OnBottom)
	 * [func NewOnBottom() *OnBottom](#NewOnBottom)

  * [OnTop](#OnTop)
	 * [func NewOnTop() *OnTop](#NewOnTop)

  * [Opacity](#Opacity)
	 * [func NewOpacity() *Opacity](#NewOpacity)

  * [Pickupable](#Pickupable)
	 * [func NewPickupable() *Pickupable](#NewPickupable)

  * [Rotateable](#Rotateable)
	 * [func NewRotateable() *Rotateable](#NewRotateable)

  * [Splash](#Splash)
	 * [func NewSplash() *Splash](#NewSplash)

  * [SpriteGroup](#SpriteGroup)

  * [Stackable](#Stackable)
	 * [func NewStackable() *Stackable](#NewStackable)

  * [Thing](#Thing)
	 * [func DeserializeThing(id uint16, typ string, datfh BinaryReader) (*Thing, error)](#DeserializeThing)
	 * [func NewThing(id uint16, typ string) *Thing](#NewThing)

  * [TopEffect](#TopEffect)
	 * [func NewTopEffect() *TopEffect](#NewTopEffect)

  * [Translucent](#Translucent)
	 * [func NewTranslucent() *Translucent](#NewTranslucent)

  * [Unwrapable](#Unwrapable)
	 * [func NewUnwrapable() *Unwrapable](#NewUnwrapable)

  * [Usable](#Usable)
	 * [func NewUsable(val uint16) *Usable](#NewUsable)

  * [Wrapable](#Wrapable)
	 * [func NewWrapable() *Wrapable](#NewWrapable)

  * [Writable](#Writable)
	 * [func NewWritable(val uint16) *Writable](#NewWritable)

  * [WritableOnce](#WritableOnce)
	 * [func NewWritableOnce(val uint16) *WritableOnce](#NewWritableOnce)

## Constants
```go
const (
	ITEM    = "item"
	OUTFIT  = "outfit"
	EFFECT  = "effect"
	MISSILE = "missile"
)
```
Some docstring to those constants

## Types

### type <a href="/blob/master/attribute.go#L421" name="AnimateAlways">AnimateAlways</a> [¶](#AnimateAlways)
```go
type AnimateAlways struct {
	AttributeBase
}
```
AnimateAlways attribute  
OpCode: 28  

#### func <a href="/blob/master/attribute.go#L429" name="NewAnimateAlways">NewAnimateAlways</a> [¶](#NewAnimateAlways)
```go
func NewAnimateAlways() *AnimateAlways
```
NewAnimateAlways creates new AnimateAlways attribute

### type <a href="/blob/master/spritegroup.go#L26" name="AnimationPhase">AnimationPhase</a> [¶](#AnimationPhase)
```go
type AnimationPhase struct {
	FrameA uint32
	FrameB uint32
}
```
AnimationPhase holds information about animation phase of SpriteGroup. One  
SpriteGroup might have many animation phases  

### type <a href="/blob/master/attribute.go#L9" name="Attribute">Attribute</a> [¶](#Attribute)
```go
type Attribute interface {
	attrName() string
}
```
Attribute is common interface for all Item attributes  

### type <a href="/blob/master/attribute.go#L14" name="AttributeBase">AttributeBase</a> [¶](#AttributeBase)
```go
type AttributeBase struct {
	XMLName xml.Name `xml:"attr"`
	Name    string   `xml:"name,attr"`
}
```
AttributeBase is base struct for all Item attributes  

### type <a href="/blob/master/attribute.go#L220" name="BlockProjectile">BlockProjectile</a> [¶](#BlockProjectile)
```go
type BlockProjectile struct {
	AttributeBase
}
```
BlockProjectile attribute  
OpCode: 14  

#### func <a href="/blob/master/attribute.go#L228" name="NewBlockProjectile">NewBlockProjectile</a> [¶](#NewBlockProjectile)
```go
func NewBlockProjectile() *BlockProjectile
```
NewBlockProjectile creates new BlockProjectile attribute

### type <a href="/blob/master/attribute.go#L506" name="Cloth">Cloth</a> [¶](#Cloth)
```go
type Cloth struct {
	AttributeBase
	Slot uint16 `xml:"slot,attr"`
}
```
Cloth attribute  
OpCode: 33  
Slots:  
* 0 - both hands  
* 1 - head  
* 2 - necklace  
* 3 - backpack  
* 4 - armor  
* 5 - shield  
* 6 - weapon  
* 7 - legs  
* 8 - feet  
* 9 - ring  
* 10 - ammo  
* 11 - store inbox  

#### func <a href="/blob/master/attribute.go#L515" name="NewCloth">NewCloth</a> [¶](#NewCloth)
```go
func NewCloth(slot uint16) *Cloth
```
NewCloth creates new Cloth attribute

### type <a href="/blob/master/attribute.go#L78" name="Container">Container</a> [¶](#Container)
```go
type Container struct {
	AttributeBase
}
```
Container attribute  
OpCode: 4  

#### func <a href="/blob/master/attribute.go#L86" name="NewContainer">NewContainer</a> [¶](#NewContainer)
```go
func NewContainer() *Container
```
NewContainer creates new Container attribute

### type <a href="/blob/master/attribute.go#L650" name="Deprecated">Deprecated</a> [¶](#Deprecated)
```go
type Deprecated struct{}
```
Deprecated attribute  
Opcode: 254  
It doesn't inherit from AttributeBase not to be XML serializable  

#### func <a href="/blob/master/attribute.go#L656" name="NewDeprecated">NewDeprecated</a> [¶](#NewDeprecated)
```go
func NewDeprecated() *Deprecated
```
NewDeprecated creates new Deprecated attribute

### type <a href="/blob/master/attribute.go#L376" name="Displacement">Displacement</a> [¶](#Displacement)
```go
type Displacement struct {
	AttributeBase
	X uint16 `xml:"x,attr"`
	Y uint16 `xml:"y,attr"`
}
```
Displacement attribute  
OpCode: 25  

#### func <a href="/blob/master/attribute.go#L386" name="NewDisplacement">NewDisplacement</a> [¶](#NewDisplacement)
```go
func NewDisplacement(x, y uint16) *Displacement
```
NewDisplacement creates new Displacement attribute

### type <a href="/blob/master/attribute.go#L348" name="DontHide">DontHide</a> [¶](#DontHide)
```go
type DontHide struct {
	AttributeBase
}
```
DontHide attribute  
OpCode: 23  

#### func <a href="/blob/master/attribute.go#L356" name="NewDontHide">NewDontHide</a> [¶](#NewDontHide)
```go
func NewDontHide() *DontHide
```
NewDontHide creates new DontHide attribute

### type <a href="/blob/master/attribute.go#L392" name="Elevation">Elevation</a> [¶](#Elevation)
```go
type Elevation struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}
```
Elevation attribute  
OpCode: 26  

#### func <a href="/blob/master/attribute.go#L401" name="NewElevation">NewElevation</a> [¶](#NewElevation)
```go
func NewElevation(val uint16) *Elevation
```
NewElevation creates new Elevation attribute

### type <a href="/blob/master/file.go#L4" name="File">File</a> [¶](#File)
```go
type File struct {
	BufferedBinaryFile
	Signature       uint32
	ContentRevision uint16
	Items           []*Thing
	Outfits         []*Thing
	Effects         []*Thing
	Missiles        []*Thing
	itemsCount      int
	outfitsCount    int
	effectsCount    int
	missilesCount   int
}
```
File is wrapper for reading .dat files  

#### func <a href="/blob/master/file.go#L19" name="Open">Open</a> [¶](#Open)
```go
func Open(path string) (*File, error)
```
Open opens given file for reading

#### func <a href="/blob/master/file.go#L136" name="File-AppendThing">AppendThing</a> [¶](#File-AppendThing)
```go
func (datfh *File) AppendThing(thing *Thing)
```
AppendThing appends given thing to proper list of (items|outfits|missiles|
effects) depending of thing Type

#### func <a href="/blob/master/file.go#L58" name="File-Deserialize">Deserialize</a> [¶](#File-Deserialize)
```go
func (datfh *File) Deserialize() error
```
Deserialize parses .dat file to extract things information

#### func <a href="/blob/master/file.go#L87" name="File-DeserializeWithProgress">DeserializeWithProgress</a> [¶](#File-DeserializeWithProgress)
```go
func (datfh *File) DeserializeWithProgress(prChan chan<- int, errChan chan<- error, doneChan chan<- bool)
```
DeserializeWithProgress does the same thing as Deserialize additionally
producing progress information to prChan channel, on finish true is
written to doneChan
As File is not thread safe, don't ever run this method in multiple
gorutines

### type <a href="/blob/master/attribute.go#L635" name="FloorChange">FloorChange</a> [¶](#FloorChange)
```go
type FloorChange struct {
	AttributeBase
}
```
FloorChange attribute  
OpCode: 252  

#### func <a href="/blob/master/attribute.go#L643" name="NewFloorChange">NewFloorChange</a> [¶](#NewFloorChange)
```go
func NewFloorChange() *FloorChange
```
NewFloorChange creates new FloorChange attribute

### type <a href="/blob/master/attribute.go#L164" name="FluidContainer">FluidContainer</a> [¶](#FluidContainer)
```go
type FluidContainer struct {
	AttributeBase
}
```
FluidContainer attribute  
OpCode: 10  

#### func <a href="/blob/master/attribute.go#L172" name="NewFluidContainer">NewFluidContainer</a> [¶](#NewFluidContainer)
```go
func NewFluidContainer() *FluidContainer
```
NewFluidContainer creates new FluidContainer attribute

### type <a href="/blob/master/attribute.go#L106" name="ForceUse">ForceUse</a> [¶](#ForceUse)
```go
type ForceUse struct {
	AttributeBase
}
```
ForceUse attribute  
OpCode: 6  

#### func <a href="/blob/master/attribute.go#L114" name="NewForceUse">NewForceUse</a> [¶](#NewForceUse)
```go
func NewForceUse() *ForceUse
```
NewForceUse creates new ForceUse attribute

### type <a href="/blob/master/attribute.go#L465" name="FullGround">FullGround</a> [¶](#FullGround)
```go
type FullGround struct {
	AttributeBase
}
```
FullGround attribute  
OpCode: 31  

#### func <a href="/blob/master/attribute.go#L473" name="NewFullGround">NewFullGround</a> [¶](#NewFullGround)
```go
func NewFullGround() *FullGround
```
NewFullGround creates new FullGround attribute

### type <a href="/blob/master/attribute.go#L21" name="Ground">Ground</a> [¶](#Ground)
```go
type Ground struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}
```
Ground attribute  
OpCode: 0  

#### func <a href="/blob/master/attribute.go#L30" name="NewGround">NewGround</a> [¶](#NewGround)
```go
func NewGround(val uint16) *Ground
```
NewGround creates new Ground attribute

### type <a href="/blob/master/attribute.go#L36" name="GroundBorder">GroundBorder</a> [¶](#GroundBorder)
```go
type GroundBorder struct {
	AttributeBase
}
```
GroundBorder attribute  
OpCode: 1  

#### func <a href="/blob/master/attribute.go#L44" name="NewGroundBorder">NewGroundBorder</a> [¶](#NewGroundBorder)
```go
func NewGroundBorder() *GroundBorder
```
NewGroundBorder creates new GroundBorder attribute

### type <a href="/blob/master/attribute.go#L276" name="Hangable">Hangable</a> [¶](#Hangable)
```go
type Hangable struct {
	AttributeBase
}
```
Hangable attribute  
OpCode: 18  

#### func <a href="/blob/master/attribute.go#L284" name="NewHangable">NewHangable</a> [¶](#NewHangable)
```go
func NewHangable() *Hangable
```
NewHangable creates new Hangable attribute

### type <a href="/blob/master/attribute.go#L304" name="HookEast">HookEast</a> [¶](#HookEast)
```go
type HookEast struct {
	AttributeBase
}
```
HookEast attribute  
OpCode: 20  

#### func <a href="/blob/master/attribute.go#L312" name="NewHookEast">NewHookEast</a> [¶](#NewHookEast)
```go
func NewHookEast() *HookEast
```
NewHookEast creates new HookEast attribute

### type <a href="/blob/master/attribute.go#L290" name="HookSouth">HookSouth</a> [¶](#HookSouth)
```go
type HookSouth struct {
	AttributeBase
}
```
HookSouth attribute  
OpCode: 19  

#### func <a href="/blob/master/attribute.go#L298" name="NewHookSouth">NewHookSouth</a> [¶](#NewHookSouth)
```go
func NewHookSouth() *HookSouth
```
NewHookSouth creates new HookSouth attribute

### type <a href="/blob/master/attribute.go#L450" name="LensHelp">LensHelp</a> [¶](#LensHelp)
```go
type LensHelp struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}
```
LensHelp attribute  
OpCode: 30  

#### func <a href="/blob/master/attribute.go#L459" name="NewLensHelp">NewLensHelp</a> [¶](#NewLensHelp)
```go
func NewLensHelp(val uint16) *LensHelp
```
NewLensHelp creates new LensHelp attribute

### type <a href="/blob/master/attribute.go#L332" name="Light">Light</a> [¶](#Light)
```go
type Light struct {
	AttributeBase
	Intensity uint16 `xml:"intensity,attr"`
	Color     uint16 `xml:"color,attr"`
}
```
Light attribute  
OpCode: 22  

#### func <a href="/blob/master/attribute.go#L342" name="NewLight">NewLight</a> [¶](#NewLight)
```go
func NewLight(intensity, color uint16) *Light
```
NewLight creates new Light attribute

### type <a href="/blob/master/attribute.go#L479" name="Look">Look</a> [¶](#Look)
```go
type Look struct {
	AttributeBase
}
```
Look attribute  
OpCode: 32  

#### func <a href="/blob/master/attribute.go#L487" name="NewLook">NewLook</a> [¶](#NewLook)
```go
func NewLook() *Look
```
NewLook creates new Look attribute

### type <a href="/blob/master/attribute.go#L407" name="LyingCorpse">LyingCorpse</a> [¶](#LyingCorpse)
```go
type LyingCorpse struct {
	AttributeBase
}
```
LyingCorpse attribute  
OpCode: 27  

#### func <a href="/blob/master/attribute.go#L415" name="NewLyingCorpse">NewLyingCorpse</a> [¶](#NewLyingCorpse)
```go
func NewLyingCorpse() *LyingCorpse
```
NewLyingCorpse creates new LyingCorpse attribute

### type <a href="/blob/master/attribute.go#L521" name="Market">Market</a> [¶](#Market)
```go
type Market struct {
	AttributeBase
	Category         uint16 `xml:"category,attr"`
	TradeAs          uint16 `xml:"tradeAs,attr"`
	ShowAs           uint16 `xml:"showAs,attr"`
	ItemName         string `xml:"itemName,attr"`
	RestrictVocation uint16 `xml:"restrictVocation,attr"`
	RequiredLevel    uint16 `xml:"requiredLevel,attr"`
}
```
Market attribute  
OpCode: 34  

#### func <a href="/blob/master/attribute.go#L535" name="NewMarket">NewMarket</a> [¶](#NewMarket)
```go
func NewMarket(category, tradeAs, showAs uint16, itemName string,
	restrictVocation, requiredLevel uint16) *Market
```
NewMarket creates new Market attribute

### type <a href="/blob/master/attribute.go#L435" name="MinimapColor">MinimapColor</a> [¶](#MinimapColor)
```go
type MinimapColor struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}
```
MinimapColor attribute  
OpCode: 29  

#### func <a href="/blob/master/attribute.go#L444" name="NewMinimapColor">NewMinimapColor</a> [¶](#NewMinimapColor)
```go
func NewMinimapColor(val uint16) *MinimapColor
```
NewMinimapColor creates new MinimapColor attribute

### type <a href="/blob/master/attribute.go#L120" name="MultiUse">MultiUse</a> [¶](#MultiUse)
```go
type MultiUse struct {
	AttributeBase
}
```
MultiUse attribute  
OpCode: 7  

#### func <a href="/blob/master/attribute.go#L128" name="NewMultiUse">NewMultiUse</a> [¶](#NewMultiUse)
```go
func NewMultiUse() *MultiUse
```
NewMultiUse creates new MultiUse attribute

### type <a href="/blob/master/attribute.go#L248" name="NoMoveAnimation">NoMoveAnimation</a> [¶](#NoMoveAnimation)
```go
type NoMoveAnimation struct {
	AttributeBase
}
```
NoMoveAnimation attribute  
OpCode: 16  

#### func <a href="/blob/master/attribute.go#L256" name="NewNoMoveAnimation">NewNoMoveAnimation</a> [¶](#NewNoMoveAnimation)
```go
func NewNoMoveAnimation() *NoMoveAnimation
```
NewNoMoveAnimation creates new NoMoveAnimation attribute

### type <a href="/blob/master/attribute.go#L206" name="NotMoveable">NotMoveable</a> [¶](#NotMoveable)
```go
type NotMoveable struct {
	AttributeBase
}
```
NotMoveable attribute  
OpCode: 13  

#### func <a href="/blob/master/attribute.go#L214" name="NewNotMoveable">NewNotMoveable</a> [¶](#NewNotMoveable)
```go
func NewNotMoveable() *NotMoveable
```
NewNotMoveable creates new NotMoveable attribute

### type <a href="/blob/master/attribute.go#L234" name="NotPathable">NotPathable</a> [¶](#NotPathable)
```go
type NotPathable struct {
	AttributeBase
}
```
NotPathable attribute  
OpCode: 15  

#### func <a href="/blob/master/attribute.go#L242" name="NewNotPathable">NewNotPathable</a> [¶](#NewNotPathable)
```go
func NewNotPathable() *NotPathable
```
NewNotPathable creates new NotPathable attribute

### type <a href="/blob/master/attribute.go#L621" name="NotPrewalkable">NotPrewalkable</a> [¶](#NotPrewalkable)
```go
type NotPrewalkable struct {
	AttributeBase
}
```
NotPrewalkable attribute  
OpCode: 101  

#### func <a href="/blob/master/attribute.go#L629" name="NewNotPrewalkable">NewNotPrewalkable</a> [¶](#NewNotPrewalkable)
```go
func NewNotPrewalkable() *NotPrewalkable
```
NewNotPrewalkable creates new NotPrewalkable attribute

### type <a href="/blob/master/attribute.go#L192" name="NotWalkable">NotWalkable</a> [¶](#NotWalkable)
```go
type NotWalkable struct {
	AttributeBase
}
```
NotWalkable attribute  
OpCode: 12  

#### func <a href="/blob/master/attribute.go#L200" name="NewNotWalkable">NewNotWalkable</a> [¶](#NewNotWalkable)
```go
func NewNotWalkable() *NotWalkable
```
NewNotWalkable creates new NotWalkable attribute

### type <a href="/blob/master/attribute.go#L50" name="OnBottom">OnBottom</a> [¶](#OnBottom)
```go
type OnBottom struct {
	AttributeBase
}
```
OnBottom attribute  
OpCode: 2  

#### func <a href="/blob/master/attribute.go#L58" name="NewOnBottom">NewOnBottom</a> [¶](#NewOnBottom)
```go
func NewOnBottom() *OnBottom
```
NewOnBottom creates new OnBottom attribute

### type <a href="/blob/master/attribute.go#L64" name="OnTop">OnTop</a> [¶](#OnTop)
```go
type OnTop struct {
	AttributeBase
}
```
OnTop attribute  
OpCode: 3  

#### func <a href="/blob/master/attribute.go#L72" name="NewOnTop">NewOnTop</a> [¶](#NewOnTop)
```go
func NewOnTop() *OnTop
```
NewOnTop creates new OnTop attribute

### type <a href="/blob/master/attribute.go#L607" name="Opacity">Opacity</a> [¶](#Opacity)
```go
type Opacity struct {
	AttributeBase
}
```
Opacity attribute  
OpCode: 100  

#### func <a href="/blob/master/attribute.go#L615" name="NewOpacity">NewOpacity</a> [¶](#NewOpacity)
```go
func NewOpacity() *Opacity
```
NewOpacity creates new Opacity attribute

### type <a href="/blob/master/attribute.go#L262" name="Pickupable">Pickupable</a> [¶](#Pickupable)
```go
type Pickupable struct {
	AttributeBase
}
```
Pickupable attribute  
OpCode: 17  

#### func <a href="/blob/master/attribute.go#L270" name="NewPickupable">NewPickupable</a> [¶](#NewPickupable)
```go
func NewPickupable() *Pickupable
```
NewPickupable creates new Pickupable attribute

### type <a href="/blob/master/attribute.go#L318" name="Rotateable">Rotateable</a> [¶](#Rotateable)
```go
type Rotateable struct {
	AttributeBase
}
```
Rotateable attribute  
OpCode: 21  

#### func <a href="/blob/master/attribute.go#L326" name="NewRotateable">NewRotateable</a> [¶](#NewRotateable)
```go
func NewRotateable() *Rotateable
```
NewRotateable creates new Rotateable attribute

### type <a href="/blob/master/attribute.go#L178" name="Splash">Splash</a> [¶](#Splash)
```go
type Splash struct {
	AttributeBase
}
```
Splash attribute  
OpCode: 11  

#### func <a href="/blob/master/attribute.go#L186" name="NewSplash">NewSplash</a> [¶](#NewSplash)
```go
func NewSplash() *Splash
```
NewSplash creates new Splash attribute

### type <a href="/blob/master/spritegroup.go#L7" name="SpriteGroup">SpriteGroup</a> [¶](#SpriteGroup)
```go
type SpriteGroup struct {
	Group           int
	FrameGroupType  uint8
	Width           uint8
	Height          uint8
	RealSize        uint8
	Layers          uint8
	PatternXNum     uint8
	PatternYNum     uint8
	PatternZNum     uint8
	Async           uint8
	StartPhase      uint8
	LoopCount       uint32
	AnimationPhases []*AnimationPhase
	Sprites         []uint32
}
```
SpriteGroup holds information about single group of sprites. One item might  
have one or more SpriteGroup assigned  

### type <a href="/blob/master/attribute.go#L92" name="Stackable">Stackable</a> [¶](#Stackable)
```go
type Stackable struct {
	AttributeBase
}
```
Stackable attribute  
OpCode: 5  

#### func <a href="/blob/master/attribute.go#L100" name="NewStackable">NewStackable</a> [¶](#NewStackable)
```go
func NewStackable() *Stackable
```
NewStackable creates new Stackable attribute

### type <a href="/blob/master/thing.go#L6" name="Thing">Thing</a> [¶](#Thing)
```go
type Thing struct {
	XMLName      xml.Name       `xml:"thing"`
	ID           uint16         `xml:"id,attr"`
	Type         string         `xml:"type,attr"`
	Attributes   []Attribute    `xml:"Attribute,omitempty"`
	SpriteGroups []*SpriteGroup `xml:"-"`
}
```
Thing holds thing information: ID, type, attributes and sprites information  

#### func <a href="/blob/master/thing.go#L28" name="DeserializeThing">DeserializeThing</a> [¶](#DeserializeThing)
```go
func DeserializeThing(id uint16, typ string, datfh BinaryReader) (*Thing, error)
```
DeserializeThing parses .dat file and creates new Thing instance

#### func <a href="/blob/master/thing.go#L23" name="NewThing">NewThing</a> [¶](#NewThing)
```go
func NewThing(id uint16, typ string) *Thing
```
NewThing creates new instance of Thing

### type <a href="/blob/master/attribute.go#L593" name="TopEffect">TopEffect</a> [¶](#TopEffect)
```go
type TopEffect struct {
	AttributeBase
}
```
TopEffect attribute  
OpCode: 38  

#### func <a href="/blob/master/attribute.go#L601" name="NewTopEffect">NewTopEffect</a> [¶](#NewTopEffect)
```go
func NewTopEffect() *TopEffect
```
NewTopEffect creates new TopEffect attribute

### type <a href="/blob/master/attribute.go#L362" name="Translucent">Translucent</a> [¶](#Translucent)
```go
type Translucent struct {
	AttributeBase
}
```
Translucent attribute  
OpCode: 24  

#### func <a href="/blob/master/attribute.go#L370" name="NewTranslucent">NewTranslucent</a> [¶](#NewTranslucent)
```go
func NewTranslucent() *Translucent
```
NewTranslucent creates new Translucent attribute

### type <a href="/blob/master/attribute.go#L579" name="Unwrapable">Unwrapable</a> [¶](#Unwrapable)
```go
type Unwrapable struct {
	AttributeBase
}
```
Unwrapable attribute  
OpCode: 37  

#### func <a href="/blob/master/attribute.go#L587" name="NewUnwrapable">NewUnwrapable</a> [¶](#NewUnwrapable)
```go
func NewUnwrapable() *Unwrapable
```
NewUnwrapable creates new Unwrapable attribute

### type <a href="/blob/master/attribute.go#L550" name="Usable">Usable</a> [¶](#Usable)
```go
type Usable struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}
```
Usable attribute  
OpCode: 35  

#### func <a href="/blob/master/attribute.go#L559" name="NewUsable">NewUsable</a> [¶](#NewUsable)
```go
func NewUsable(val uint16) *Usable
```
NewUsable creates new Usable attribute

### type <a href="/blob/master/attribute.go#L565" name="Wrapable">Wrapable</a> [¶](#Wrapable)
```go
type Wrapable struct {
	AttributeBase
}
```
Wrapable attribute  
OpCode: 36  

#### func <a href="/blob/master/attribute.go#L573" name="NewWrapable">NewWrapable</a> [¶](#NewWrapable)
```go
func NewWrapable() *Wrapable
```
NewWrapable creates new Wrapable attribute

### type <a href="/blob/master/attribute.go#L134" name="Writable">Writable</a> [¶](#Writable)
```go
type Writable struct {
	AttributeBase
	TextLen uint16 `xml:"textLen,attr"`
}
```
Writable attribute  
OpCode: 8  

#### func <a href="/blob/master/attribute.go#L143" name="NewWritable">NewWritable</a> [¶](#NewWritable)
```go
func NewWritable(val uint16) *Writable
```
NewWritable creates new Writable attribute

### type <a href="/blob/master/attribute.go#L149" name="WritableOnce">WritableOnce</a> [¶](#WritableOnce)
```go
type WritableOnce struct {
	AttributeBase
	TextLen uint16 `xml:"textLen,attr"`
}
```
WritableOnce attribute  
OpCode: 9  

#### func <a href="/blob/master/attribute.go#L158" name="NewWritableOnce">NewWritableOnce</a> [¶](#NewWritableOnce)
```go
func NewWritableOnce(val uint16) *WritableOnce
```
NewWritableOnce creates new WritableOnce attribute

***
_Last updated 8 Aug 2016_

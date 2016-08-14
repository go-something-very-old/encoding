package dat

import (
	"fmt"

	bin "github.com/go-otserv/encoding/binary"
)

// SpriteGroup holds information about single group of sprites. One item might
// have one or more SpriteGroup assigned
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

// AnimationPhase holds information about animation phase of SpriteGroup. One
// SpriteGroup might have many animation phases
type AnimationPhase struct {
	FrameA uint32
	FrameB uint32
}

func deserializeSpriteGroup(typ string, datfh bin.Reader) (*SpriteGroup, error) {
	var err error
	var sprID uint32
	var animPhases uint8

	sprGr := &SpriteGroup{}

	if typ == OUTFIT {
		if sprGr.FrameGroupType, err = datfh.UInt8(); err != nil {
			return nil, err
		}
	}

	if sprGr.Width, err = datfh.UInt8(); err != nil {
		return nil, err
	}
	if sprGr.Height, err = datfh.UInt8(); err != nil {
		return nil, err
	}
	if sprGr.Width > 1 || sprGr.Height > 1 {
		if sprGr.RealSize, err = datfh.UInt8(); err != nil {
			return nil, err
		}
	}

	if sprGr.Layers, err = datfh.UInt8(); err != nil {
		return nil, err
	}
	if sprGr.PatternXNum, err = datfh.UInt8(); err != nil {
		return nil, err
	}
	if sprGr.PatternYNum, err = datfh.UInt8(); err != nil {
		return nil, err
	}
	if sprGr.PatternZNum, err = datfh.UInt8(); err != nil {
		return nil, err
	}
	if animPhases, err = datfh.UInt8(); err != nil {
		return nil, err
	}

	sprGr.AnimationPhases = make([]*AnimationPhase, 0, animPhases-1)
	if animPhases > 1 {
		if sprGr.Async, err = datfh.UInt8(); err != nil {
			return nil, err
		}
		if sprGr.StartPhase, err = datfh.UInt8(); err != nil {
			return nil, err
		}
		if sprGr.LoopCount, err = datfh.UInt32(); err != nil {
			return nil, err
		}

		for phase := 0; phase < int(animPhases); phase++ {
			animPhase := &AnimationPhase{}
			if animPhase.FrameA, err = datfh.UInt32(); err != nil {
				return nil, err
			}
			if animPhase.FrameB, err = datfh.UInt32(); err != nil {
				return nil, err
			}
			sprGr.AnimationPhases = append(sprGr.AnimationPhases, animPhase)
		}
	}

	sprCount := int(sprGr.Width) * int(sprGr.Height) * int(sprGr.Layers) *
		int(sprGr.PatternXNum) * int(sprGr.PatternYNum) * int(sprGr.PatternZNum) *
		int(animPhases)
	if sprCount > 4096 {
		return nil, fmt.Errorf("sprites count for item > 4096")
	}

	for sprNum := 0; sprNum < sprCount; sprNum++ {
		if sprID, err = datfh.UInt32(); err != nil {
			return nil, err
		}
		sprGr.Sprites = append(sprGr.Sprites, sprID)
	}
	return sprGr, nil
}

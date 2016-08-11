package dat

import (
	bin "github.com/go-otserv/encoding/binary"
)

// File is wrapper for reading .dat files
type File struct {
	bin.BufferedBinaryFile
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

// Open opens given file for reading
func Open(path string) (*File, error) {
	var itemsCount, outfitsCount, effectsCount, missilesCount uint16

	buffh, err := OpenBufferedBinaryFile(path)
	if err != nil {
		return nil, err
	}
	datfh := &File{*buffh, 0, 0, nil, nil, nil, nil, 0, 0, 0, 0}

	if datfh.Signature, err = datfh.UInt32(); err != nil {
		return datfh, err
	}
	datfh.ContentRevision = uint16(datfh.Signature)
	if itemsCount, err = datfh.UInt16(); err != nil {
		return datfh, err
	}
	if outfitsCount, err = datfh.UInt16(); err != nil {
		return datfh, err
	}
	if effectsCount, err = datfh.UInt16(); err != nil {
		return datfh, err
	}
	if missilesCount, err = datfh.UInt16(); err != nil {
		return datfh, err
	}
	datfh.itemsCount = int(itemsCount + 1)
	datfh.outfitsCount = int(outfitsCount + 1)
	datfh.effectsCount = int(effectsCount + 1)
	datfh.missilesCount = int(missilesCount + 1)

	datfh.Items = make([]*Thing, 0, itemsCount+1)
	datfh.Outfits = make([]*Thing, 0, outfitsCount+1)
	datfh.Effects = make([]*Thing, 0, effectsCount+1)
	datfh.Missiles = make([]*Thing, 0, missilesCount+1)

	return datfh, nil
}

// Deserialize parses .dat file to extract things information
func (datfh *File) Deserialize() error {
	prChan := make(chan int)
	errChan := make(chan error)
	doneChan := make(chan bool)
	defer close(prChan)
	defer close(errChan)
	defer close(doneChan)

	go datfh.DeserializeWithProgress(prChan, errChan, doneChan)
	consumeProgress := func() error {
		for {
			select {
			case <-prChan:
				break
			case err := <-errChan:
				return err
			case <-doneChan:
				return nil
			}
		}
	}
	return consumeProgress()
}

// DeserializeWithProgress does the same thing as Deserialize additionally
// producing progress information to prChan channel, on finish true is
// written to doneChan
// As File is not thread safe, don't ever run this method in multiple
// gorutines
func (datfh *File) DeserializeWithProgress(prChan chan<- int, errChan chan<- error, doneChan chan<- bool) {
	var err error
	var thing *Thing

	typeNames := []string{ITEM, OUTFIT, EFFECT, MISSILE}
	typeCount := map[string]int{
		ITEM:    datfh.itemsCount,
		OUTFIT:  datfh.outfitsCount,
		EFFECT:  datfh.effectsCount,
		MISSILE: datfh.missilesCount,
	}
	typeToFirstID := map[string]int{
		ITEM:    100,
		OUTFIT:  1,
		EFFECT:  1,
		MISSILE: 1,
	}
	allCount := datfh.itemsCount + datfh.outfitsCount + datfh.effectsCount +
		datfh.missilesCount

	progress := func(n, total float64) int {
		return int(n / total * 100)
	}

	currentProgress := -1
	previousProgress := -1
	var commonID uint16 = 99
	for _, typ := range typeNames {
		firstID := typeToFirstID[typ]
		for itemCid := firstID; itemCid < typeCount[typ]; itemCid++ {
			commonID++
			if thing, err = DeserializeThing(commonID, typ, datfh); err != nil {
				errChan <- err
				return
			}
			datfh.AppendThing(thing)

			currentProgress = progress(float64(commonID), float64(allCount))
			if currentProgress != previousProgress {
				previousProgress = currentProgress
				prChan <- currentProgress
			}
		}
	}
	doneChan <- true
}

// AppendThing appends given thing to proper list of (items|outfits|missiles|
// effects) depending of thing Type
func (datfh *File) AppendThing(thing *Thing) {
	switch thing.Type {
	case ITEM:
		datfh.Items = append(datfh.Items, thing)
	case OUTFIT:
		datfh.Outfits = append(datfh.Outfits, thing)
	case EFFECT:
		datfh.Effects = append(datfh.Effects, thing)
	case MISSILE:
		datfh.Missiles = append(datfh.Missiles, thing)
	}
}

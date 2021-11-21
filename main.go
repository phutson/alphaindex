package alphaindex

import (
	"errors"
	"regexp"
	"strings"
)

type CellColumnIndex struct {
	cellIndex []byte
}

func DecRecursive(index []byte, workingPiece int) []byte {
	workingByte := index[workingPiece]
	if workingByte == 'A' {
		if workingPiece == 0 {
			index = index[1:]
			return index
		}
		index[workingPiece] = 'Z'
		index = DecRecursive(index, workingPiece-1)
		return index
	} else {
		index[workingPiece] = workingByte - 1
		return index
	}
}

func IncRecursive(index []byte, workingPiece int) []byte {

	if workingPiece < 0 {
		outIndex := make([]byte, 0, 1)
		outIndex = append(outIndex, 'A')
		index = append(outIndex, index...)
	} else {
		workingByte := index[workingPiece]
		if workingByte == 'Z' {
			index[workingPiece] = 'A'
			index = IncRecursive(index, workingPiece-1)
		} else {
			index[workingPiece] = workingByte + 1
		}
	}
	return index
}

func (cci *CellColumnIndex) Increment() {
	cellCount := len(cci.cellIndex)
	if cellCount == 0 {
		cci.cellIndex = append(cci.cellIndex, 'A')
		return
	}
	workingCell := cellCount - 1
	cci.cellIndex = IncRecursive(cci.cellIndex, workingCell)
}

func (cci *CellColumnIndex) Decriment() {
	cellCount := len(cci.cellIndex)
	if cellCount == 0 {

		return
	}
	workingCell := cellCount - 1
	cci.cellIndex = DecRecursive(cci.cellIndex, workingCell)
}

func (cci *CellColumnIndex) GetIndexString() string {
	return string(cci.cellIndex[:])
}

func NewIndex(initialIndex string) (*CellColumnIndex, error) {

	var cci CellColumnIndex
	if len(initialIndex) > 0 {
		lettersOnly := regexp.MustCompile(`^[a-zA-Z]*$`)
		if !lettersOnly.MatchString(initialIndex) {
			return nil, errors.New("invalid initial index")
		}
		cci.cellIndex = []byte(strings.ToUpper(initialIndex))
	} else {
		cci.cellIndex = make([]byte, 0, 5)
	}
	return &cci, nil
}

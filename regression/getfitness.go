package regression

import (
	"strings"

	"github.com/sjwhitworth/golearn/base"
)

func GetColumns(data *base.DenseInstances, bitString string, nonClass []base.Attribute, class []base.Attribute) *base.DenseInstances {
	bitSlice := strings.Split(bitString, "")

	attrs := make([]base.Attribute, 0)
	for i, a := range nonClass {
		if bitSlice[i] == "1" {
			attrs = append(attrs, a)
		}
	}

	attrs = append(attrs, class...)

	oldSpecs := make([]base.AttributeSpec, len(attrs))
	newSpecs := make([]base.AttributeSpec, len(attrs))

	newInst := base.NewDenseInstances()

	for i, a := range attrs {
		// Retrieve old AttributeSpec
		s, err := data.GetAttribute(a)
		if err != nil {
			panic(err)
		}
		oldSpecs[i] = s
		newSpecs[i] = newInst.AddAttribute(a)
	}

	for _, a := range class {
		newInst.AddClassAttribute(a)
	}

	// Allocate memory
	_, rows := data.Size()
	newInst.Extend(rows)

	// Copy each row from the old one to the new
	data.MapOverRows(oldSpecs, func(v [][]byte, r int) (bool, error) {
		for i, c := range v {
			newInst.Set(newSpecs[i], r, c)
		}
		return true, nil
	})

	return newInst
}

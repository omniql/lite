package store

import (
	"github.com/omniql/reflect"
	"io"
)



type Item struct {
	commons.Uint64Model
	SKU       string
	Condition Condition
}

type Description struct {
	commons.Uint16Model
	Title       string
	Description string
}

type ItemPatch struct {
	ID             uint64
	Title          *string
	SKU            *string
	UnitPrice      *int64
	LastVersion    int64
	ClientID       string
	ClientVersion  int64
	ClearVariation *bool
	//push variation
	Variations {
	opert: []Variation
}

	type Color struct {
	Name string
	Code string
}
}
func GenerateStore(container reflect.ApplicationContainer) (err error) {
	//create json and json patch for every resource
	resourceCount := container.LookupResources().ResourceCount()
	for i := 0; i < resourceCount; i++ {
		r, err := container.LookupResources().ResourceByKind(uint16(i))
		if err != nil {
			return
		}
		err = GenerateResourceModel(r)
		if err != nil {
			return
		}
	}
}

func GenerateResourceModel(container reflect.ResourceContainer, wr io.Writer) (err error) {
	//create model

}


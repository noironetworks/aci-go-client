package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanVDestClassName = "spanVDest"

type SpanVDest struct {
	BaseAttributes
    SpanVDestAttributes
}

type SpanVDestAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias        string `json:",omitempty"`
}


func NewSpanVDest(spanVDestRn, parentDn, description string, spanVDestAttr SpanVDestAttributes) *SpanVDest {
	dn := fmt.Sprintf("%s/%s", parentDn, spanVDestRn)
	return &SpanVDest{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         SpanVDestClassName,
			Rn:                spanVDestRn,
		},

		SpanVDestAttributes: spanVDestAttr,

	}
}

func (spanVDest *SpanVDest) ToMap() (map[string]string, error) {
	spanVDestMap, err := spanVDest.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(spanVDestMap, "annotation",spanVDest.Annotation)
	A(spanVDestMap, "nameAlias",spanVDest.NameAlias)
	return spanVDestMap, err
}

func SpanVDestFromContainerList(cont *container.Container, index int) *SpanVDest {

	SpanVDestCont := cont.S("imdata").Index(index).S(SpanVDestClassName, "attributes")
	return &SpanVDest {
		BaseAttributes{
			DistinguishedName: G(SpanVDestCont, "dn"),
			Description:       G(SpanVDestCont, "descr"),
			Status:            G(SpanVDestCont, "status"),
			ClassName:         SpanVDestClassName,
			Rn:                G(SpanVDestCont, "rn"),
		},

		SpanVDestAttributes{
			Annotation : G(SpanVDestCont, "annotation"),
			NameAlias : G(SpanVDestCont, "nameAlias"),
		},
	}
}

func SpanVDestFromContainer(cont *container.Container) *SpanVDest {

	return SpanVDestFromContainerList(cont, 0)
}

func SpanVDestListFromContainer(cont *container.Container) []*SpanVDest {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanVDest, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanVDestFromContainerList(cont, i)
	}

	return arr
}

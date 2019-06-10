package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanSpanLblClassName = "spanSpanLbl"

type SpanVSpanLbl struct {
	BaseAttributes
    SpanVSpanLblAttributes
}

type SpanVSpanLblAttributes struct {
    Annotation       string `json:",omitempty"`
	Name             string `json:",omitempty"`
    NameAlias        string `json:",omitempty"`
	Tag 	         string `json:",omitempty"`
}


func NewSpanVSpanLbl(spanVSpanLblRn, parentDn, description string, spanVSpanLblAttr SpanVSpanLblAttributes) *SpanVSpanLbl {
	dn := fmt.Sprintf("%s/%s", parentDn, spanVSpanLblRn)
	return &SpanVSpanLbl{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         SpanSpanLblClassName,
			Rn:                spanVSpanLblRn,
		},

		SpanVSpanLblAttributes: spanVSpanLblAttr,

	}
}

func (spanVSpanLbl *SpanVSpanLbl) ToMap() (map[string]string, error) {
	spanVSpanLblMap, err := spanVSpanLbl.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(spanVSpanLblMap, "annotation",spanVSpanLbl.Annotation)
	A(spanVSpanLblMap, "name",spanVSpanLbl.Name)
	A(spanVSpanLblMap, "nameAlias",spanVSpanLbl.NameAlias)
	A(spanVSpanLblMap, "tag",spanVSpanLbl.Tag)
	return spanVSpanLblMap, err
}

func SpanVSpanLblFromContainerList(cont *container.Container, index int) *SpanVSpanLbl {

	SpanVSpanLblCont := cont.S("imdata").Index(index).S(SpanSpanLblClassName, "attributes")
	return &SpanVSpanLbl {
		BaseAttributes{
			DistinguishedName: G(SpanVSpanLblCont, "dn"),
			Description:       G(SpanVSpanLblCont, "descr"),
			Status:            G(SpanVSpanLblCont, "status"),
			ClassName:         SpanSpanLblClassName,
			Rn:                G(SpanVSpanLblCont, "rn"),
		},

		SpanVSpanLblAttributes{
			Annotation : G(SpanVSpanLblCont, "annotation"),
			Name       : G(SpanVSpanLblCont, "name"),
			NameAlias  : G(SpanVSpanLblCont, "nameAlias"),
			Tag        : G(SpanVSpanLblCont, "tag"),
		},
	}
}

func SpanVSpanLblFromContainer(cont *container.Container) *SpanVSpanLbl {

	return SpanVSpanLblFromContainerList(cont, 0)
}

func SpanVSpanLblListFromContainer(cont *container.Container) []*SpanVSpanLbl {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanVSpanLbl, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanVSpanLblFromContainerList(cont, i)
	}

	return arr
}

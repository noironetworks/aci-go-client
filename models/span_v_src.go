package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanVSrcClassName = "spanVSrc"

type SpanVSrc struct {
	BaseAttributes
    SpanVSrcAttributes
}

type SpanVSrcAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias        string `json:",omitempty"`
	Dir		         string `json:",omitempty"`
}


func NewSpanVSrc(spanVSrcRn, parentDn, description string, spanVSrcAttr SpanVSrcAttributes) *SpanVSrc {
	dn := fmt.Sprintf("%s/%s", parentDn, spanVSrcRn)
	return &SpanVSrc{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         SpanVSrcClassName,
			Rn:                spanVSrcRn,
		},

		SpanVSrcAttributes: spanVSrcAttr,

	}
}

func (spanVSrc *SpanVSrc) ToMap() (map[string]string, error) {
	spanVSrcMap, err := spanVSrc.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(spanVSrcMap, "annotation",spanVSrc.Annotation)
	A(spanVSrcMap, "nameAlias",spanVSrc.NameAlias)
	A(spanVSrcMap, "dir",spanVSrc.Dir)
	return spanVSrcMap, err
}

func SpanVSrcFromContainerList(cont *container.Container, index int) *SpanVSrc {

	SpanVSrcCont := cont.S("imdata").Index(index).S(SpanVSrcClassName, "attributes")
	return &SpanVSrc {
		BaseAttributes{
			DistinguishedName: G(SpanVSrcCont, "dn"),
			Description:       G(SpanVSrcCont, "descr"),
			Status:            G(SpanVSrcCont, "status"),
			ClassName:         SpanVSrcClassName,
			Rn:                G(SpanVSrcCont, "rn"),
		},

		SpanVSrcAttributes{
			Annotation : G(SpanVSrcCont, "annotation"),
			NameAlias : G(SpanVSrcCont, "nameAlias"),
			Dir : G(SpanVSrcCont, "dir"),
		},
	}
}

func SpanVSrcFromContainer(cont *container.Container) *SpanVSrc {

	return SpanVSrcFromContainerList(cont, 0)
}

func SpanVSrcListFromContainer(cont *container.Container) []*SpanVSrc {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanVSrc, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanVSrcFromContainerList(cont, i)
	}

	return arr
}

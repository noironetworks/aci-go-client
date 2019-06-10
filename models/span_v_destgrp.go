package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanVDestGrpClassName = "spanVDestGrp"

type SpanVDestGrp struct {
	BaseAttributes
    SpanVDestGrpAttributes
}

type SpanVDestGrpAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias        string `json:",omitempty"`
}


func NewSpanVDestGrp(spanVDestGrpRn, parentDn, description string, spanVDestGrpAttr SpanVDestGrpAttributes) *SpanVDestGrp {
	dn := fmt.Sprintf("%s/%s", parentDn, spanVDestGrpRn)
	return &SpanVDestGrp{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         SpanVDestGrpClassName,
			Rn:                spanVDestGrpRn,
		},

		SpanVDestGrpAttributes: spanVDestGrpAttr,

	}
}

func (spanVDestGrp *SpanVDestGrp) ToMap() (map[string]string, error) {
	spanVDestGrpMap, err := spanVDestGrp.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(spanVDestGrpMap, "annotation",spanVDestGrp.Annotation)
	A(spanVDestGrpMap, "nameAlias",spanVDestGrp.NameAlias)
	return spanVDestGrpMap, err
}

func SpanVDestGrpFromContainerList(cont *container.Container, index int) *SpanVDestGrp {

	SpanVDestGrpCont := cont.S("imdata").Index(index).S(SpanVDestGrpClassName, "attributes")
	return &SpanVDestGrp {
		BaseAttributes{
			DistinguishedName: G(SpanVDestGrpCont, "dn"),
			Description:       G(SpanVDestGrpCont, "descr"),
			Status:            G(SpanVDestGrpCont, "status"),
			ClassName:         SpanVDestGrpClassName,
			Rn:                G(SpanVDestGrpCont, "rn"),
		},

		SpanVDestGrpAttributes{
			Annotation : G(SpanVDestGrpCont, "annotation"),
			NameAlias : G(SpanVDestGrpCont, "nameAlias"),
		},
	}
}

func SpanVDestGrpFromContainer(cont *container.Container) *SpanVDestGrp {

	return SpanVDestGrpFromContainerList(cont, 0)
}

func SpanVDestGrpListFromContainer(cont *container.Container) []*SpanVDestGrp {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanVDestGrp, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanVDestGrpFromContainerList(cont, i)
	}

	return arr
}

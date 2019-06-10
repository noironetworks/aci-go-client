package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanVSrcGrpClassName = "spanVSrcGrp"

type SpanVSrcGrp struct {
	BaseAttributes
    SpanVSrcGrpAttributes
}

type SpanVSrcGrpAttributes struct {
    Annotation       string `json:",omitempty"`
    NameAlias        string `json:",omitempty"`
	AdminSt          string `json:",omitempty"`
}


func NewSpanVSrcGrp(spanVSrcGrpRn, parentDn, description string, spanVSrcGrpAttr SpanVSrcGrpAttributes) *SpanVSrcGrp {
	dn := fmt.Sprintf("%s/%s", parentDn, spanVSrcGrpRn)
	return &SpanVSrcGrp{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         SpanVSrcGrpClassName,
			Rn:                spanVSrcGrpRn,
		},

		SpanVSrcGrpAttributes: spanVSrcGrpAttr,

	}
}

func (spanVSrcGrp *SpanVSrcGrp) ToMap() (map[string]string, error) {
	spanVSrcGrpMap, err := spanVSrcGrp.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(spanVSrcGrpMap, "annotation",spanVSrcGrp.Annotation)
	A(spanVSrcGrpMap, "nameAlias",spanVSrcGrp.NameAlias)
	A(spanVSrcGrpMap, "adminSt",spanVSrcGrp.AdminSt)
	return spanVSrcGrpMap, err
}

func SpanVSrcGrpFromContainerList(cont *container.Container, index int) *SpanVSrcGrp {

	SpanVSrcGrpCont := cont.S("imdata").Index(index).S(SpanVSrcGrpClassName, "attributes")
	return &SpanVSrcGrp {
		BaseAttributes{
			DistinguishedName: G(SpanVSrcGrpCont, "dn"),
			Description:       G(SpanVSrcGrpCont, "descr"),
			Status:            G(SpanVSrcGrpCont, "status"),
			ClassName:         SpanVSrcGrpClassName,
			Rn:                G(SpanVSrcGrpCont, "rn"),
		},

		SpanVSrcGrpAttributes{
			Annotation : G(SpanVSrcGrpCont, "annotation"),
			NameAlias : G(SpanVSrcGrpCont, "nameAlias"),
			AdminSt : G(SpanVSrcGrpCont, "adminSt"),
		},
	}
}

func SpanVSrcGrpFromContainer(cont *container.Container) *SpanVSrcGrp {

	return SpanVSrcGrpFromContainerList(cont, 0)
}

func SpanVSrcGrpListFromContainer(cont *container.Container) []*SpanVSrcGrp {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanVSrcGrp, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanVSrcGrpFromContainerList(cont, i)
	}

	return arr
}

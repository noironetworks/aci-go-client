package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanVEpgSummaryClassName = "spanVEpgSummary"

type SpanVEpgSummary struct {
	BaseAttributes
    SpanVEpgSummaryAttributes
}

type SpanVEpgSummaryAttributes struct {
    Annotation       string `json:",omitempty"`
    Dscp             string `json:",omitempty"`
    DstIp            string `json:",omitempty"`
    FlowId           string `json:",omitempty"`
    Mode             string `json:",omitempty"`
    Mtu              string `json:",omitempty"`
    NameAlias        string `json:",omitempty"`
    SrcIpPrefix      string `json:",omitempty"`
    Ttl              string `json:",omitempty"`
    VrfName          string `json:",omitempty"`

}


func NewSpanVEpgSummary(spanVEpgSummaryRn, parentDn, description string, spanVEpgSummaryAttr SpanVEpgSummaryAttributes) *SpanVEpgSummary {
	dn := fmt.Sprintf("%s/%s", parentDn, spanVEpgSummaryRn)
	return &SpanVEpgSummary{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         SpanVEpgSummaryClassName,
			Rn:                spanVEpgSummaryRn,
		},

		SpanVEpgSummaryAttributes: spanVEpgSummaryAttr,

	}
}

func (spanVEpgSummary *SpanVEpgSummary) ToMap() (map[string]string, error) {
	spanVEpgSummaryMap, err := spanVEpgSummary.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(spanVEpgSummaryMap, "annotation",spanVEpgSummary.Annotation)
    A(spanVEpgSummaryMap, "dscp",spanVEpgSummary.Dscp)
    A(spanVEpgSummaryMap, "dstIp",spanVEpgSummary.DstIp)
    A(spanVEpgSummaryMap, "flowId",spanVEpgSummary.FlowId)
    A(spanVEpgSummaryMap, "mode",spanVEpgSummary.Mode)
    A(spanVEpgSummaryMap, "mtu",spanVEpgSummary.Mtu)
    A(spanVEpgSummaryMap, "nameAlias",spanVEpgSummary.NameAlias)
    A(spanVEpgSummaryMap, "srcIpPrefix",spanVEpgSummary.SrcIpPrefix)
    A(spanVEpgSummaryMap, "ttl",spanVEpgSummary.Ttl)
    A(spanVEpgSummaryMap, "vrfName",spanVEpgSummary.VrfName)

	return spanVEpgSummaryMap, err
}

func SpanVEpgSummaryFromContainerList(cont *container.Container, index int) *SpanVEpgSummary {

	SpanVEpgSummaryCont := cont.S("imdata").Index(index).S(SpanVEpgSummaryClassName, "attributes")
	return &SpanVEpgSummary {
		BaseAttributes{
			DistinguishedName: G(SpanVEpgSummaryCont, "dn"),
			Description:       G(SpanVEpgSummaryCont, "descr"),
			Status:            G(SpanVEpgSummaryCont, "status"),
			ClassName:         SpanVEpgSummaryClassName,
			Rn:                G(SpanVEpgSummaryCont, "rn"),
		},

		SpanVEpgSummaryAttributes{
			Annotation : G(SpanVEpgSummaryCont, "annotation"),
			Dscp : G(SpanVEpgSummaryCont, "dscp"),
			DstIp : G(SpanVEpgSummaryCont, "dstIp"),
			FlowId : G(SpanVEpgSummaryCont, "flowId"),
			Mode : G(SpanVEpgSummaryCont, "mode"),
			Mtu : G(SpanVEpgSummaryCont, "mtu"),
			NameAlias : G(SpanVEpgSummaryCont, "nameAlias"),
			SrcIpPrefix : G(SpanVEpgSummaryCont, "srcIpPrefix"),
			Ttl : G(SpanVEpgSummaryCont, "ttl"),
			VrfName : G(SpanVEpgSummaryCont, "vrfName"),
        },

	}
}

func SpanVEpgSummaryFromContainer(cont *container.Container) *SpanVEpgSummary {

	return SpanVEpgSummaryFromContainerList(cont, 0)
}

func SpanVEpgSummaryListFromContainer(cont *container.Container) []*SpanVEpgSummary {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanVEpgSummary, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanVEpgSummaryFromContainerList(cont, i)
	}

	return arr
}

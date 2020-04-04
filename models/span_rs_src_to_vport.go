package models

import (
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const SpanRsSrcToVPortClassName = "spanRsSrcToVPort"

type SpanRsSrcToVPort struct {
	BaseAttributes
	SpanRsSrcToVPortAttributes
}

type SpanRsSrcToVPortAttributes struct {
	State             string `json:",omitempty"`
	Tdn               string `json:",omitempty"`
}


func SpanRsSrcToVPortFromContainerList(cont *container.Container, index int) *SpanRsSrcToVPort {

	SpanRsSrcToVPortCont := cont.S("imdata").Index(index).S(SpanRsSrcToVPortClassName, "attributes")
	return &SpanRsSrcToVPort{
		BaseAttributes{
			DistinguishedName: G(SpanRsSrcToVPortCont, "dn"),
			Status:            G(SpanRsSrcToVPortCont, "status"),
			ClassName:         SpanRsSrcToVPortClassName,
			Rn:                G(SpanRsSrcToVPortCont, "rn"),
		},

		SpanRsSrcToVPortAttributes{
			State             : G(SpanRsSrcToVPortCont, "state"),
			Tdn               : G(SpanRsSrcToVPortCont, "tDn"),
        },
	}
}

func SpanRsSrcToVPortFromContainer(cont *container.Container) *SpanRsSrcToVPort {

	return SpanRsSrcToVPortFromContainerList(cont, 0)
}

func SpanRsSrcToVPortListFromContainer(cont *container.Container) []*SpanRsSrcToVPort {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*SpanRsSrcToVPort, length)

	for i := 0; i < length; i++ {

		arr[i] = SpanRsSrcToVPortFromContainerList(cont, i)
	}

	return arr
}
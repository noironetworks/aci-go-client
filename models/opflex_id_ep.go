package models

import (
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const OpflexidepClassName = "opflexIDEp"

type OpflexIDEp struct {
	BaseAttributes
	OpflexIDEpAttributes
}
  
type OpflexIDEpAttributes struct {
	BrIfId            string `json:",omitempty"`
	CompHvDn          string `json:",omitempty"`
	CompVmDn          string `json:",omitempty"`
	ContainerName     string `json:",omitempty"`
	CtrlrName         string `json:",omitempty"`
	DomName           string `json:",omitempty"`
	DomPDn            string `json:",omitempty"`
	Encap             string `json:",omitempty"`
	EncapMode         string `json:",omitempty"`
	EpHostAddr        string `json:",omitempty"`
	EpgID             string `json:",omitempty"`
	EpgPKey           string `json:",omitempty"`
	Eppdn             string `json:",omitempty"`
	HypervisorName    string `json:",omitempty"`
	Id                string `json:",omitempty"`
	InstType          string `json:",omitempty"`
	Ip                string `json:",omitempty"`
	Mac               string `json:",omitempty"`
	McastAddr         string `json:",omitempty"`
	NameAlias         string `json:",omitempty"`
	Namespace         string `json:",omitempty"`
	PcIfId            string `json:",omitempty"`
	PortId            string `json:",omitempty"`
	State             string `json:",omitempty"`
	TransitionStatus  string `json:",omitempty"`
	VendorId          string `json:",omitempty"`
	VmAttr            string `json:",omitempty"`
	VmmSrc            string `json:",omitempty"`
}


func OpflexIDEpFromContainerList(cont *container.Container, index int) *OpflexIDEp {

	OpflexIDEpCont := cont.S("imdata").Index(index).S(OpflexidepClassName, "attributes")
	return &OpflexIDEp{
		BaseAttributes{
			DistinguishedName: G(OpflexIDEpCont, "dn"),
			Description:       G(OpflexIDEpCont, "descr"),
			Status:            G(OpflexIDEpCont, "status"),
			ClassName:         OpflexidepClassName,
			Rn:                G(OpflexIDEpCont, "rn"),
		},

		OpflexIDEpAttributes{
			BrIfId            : G(OpflexIDEpCont, "brIfId"),
			CompHvDn          : G(OpflexIDEpCont, "compHvDn"),
			CompVmDn          : G(OpflexIDEpCont, "compVmDn"),
			ContainerName     : G(OpflexIDEpCont, "containerName"),
			CtrlrName         : G(OpflexIDEpCont, "ctrlrName"),
			DomName           : G(OpflexIDEpCont, "domName"),
			DomPDn            : G(OpflexIDEpCont, "domPDn"),
			Encap             : G(OpflexIDEpCont, "encap"),
			EncapMode         : G(OpflexIDEpCont, "encapMode"),
			EpHostAddr        : G(OpflexIDEpCont, "epHostAddr"),
			EpgID             : G(OpflexIDEpCont, "epgID"),
			EpgPKey           : G(OpflexIDEpCont, "epgPKey"),
			Eppdn             : G(OpflexIDEpCont, "eppdn"),
			HypervisorName    : G(OpflexIDEpCont, "hypervisorName"),
			Id                : G(OpflexIDEpCont, "id"),
			InstType          : G(OpflexIDEpCont, "instType"),
			Ip                : G(OpflexIDEpCont, "ip"),
			Mac               : G(OpflexIDEpCont, "mac"),
			McastAddr         : G(OpflexIDEpCont, "mcastAddr"),
			NameAlias         : G(OpflexIDEpCont, "nameAlias"),
			Namespace         : G(OpflexIDEpCont, "namespace"),
			PcIfId            : G(OpflexIDEpCont, "pcIfId"),
			PortId            : G(OpflexIDEpCont, "portId"),
			State             : G(OpflexIDEpCont, "state"),
			TransitionStatus  : G(OpflexIDEpCont, "transitionStatus"),
			VendorId          : G(OpflexIDEpCont, "vendorId"),
			VmAttr            : G(OpflexIDEpCont, "vmAttr"),
			VmmSrc            : G(OpflexIDEpCont, "vmmSrc"),
        },
	}
}

func OpflexIDEpFromContainer(cont *container.Container) *OpflexIDEp {

	return OpflexIDEpFromContainerList(cont, 0)
}

func OpflexIDEpListFromContainer(cont *container.Container) []*OpflexIDEp {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*OpflexIDEp, length)

	for i := 0; i < length; i++ {

		arr[i] = OpflexIDEpFromContainerList(cont, i)
	}

	return arr
}

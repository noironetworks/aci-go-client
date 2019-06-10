package client

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSpanVSpanLbl(name string, vsrc_grp string, description string,
	spanVSpanLblattr models.SpanVSpanLblAttributes) (*models.SpanVSpanLbl, error) {
	rn := fmt.Sprintf("spanlbl-%s", name)
	parentDn := fmt.Sprintf("uni/infra/vsrcgrp-%s", vsrc_grp)
	spanVSpanLbl := models.NewSpanVSpanLbl(rn, parentDn, description, spanVSpanLblattr)
	err := sm.Save(spanVSpanLbl)
	return spanVSpanLbl, err
}


func (sm *ServiceManager) ReadSpanVSpanLbl(name string, vsrc_grp string ) (*models.SpanVSpanLbl, error) {
	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s/spanlbl-%s", vsrc_grp, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	spanVSpanLbl := models.SpanVSpanLblFromContainer(cont)
	return spanVSpanLbl, nil
}


func (sm *ServiceManager) DeleteSpanVSpanLbl(name string, vsrc_grp string ) (error) {
	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s/spanlbl-%s", vsrc_grp, name)
	return sm.DeleteByDn(dn, models.SpanSpanLblClassName)
}


func (sm *ServiceManager) UpdateSpanVSpanLbl(name string, vsrc_grp string, description string,
	spanVSpanLblattr models.SpanVSpanLblAttributes) (*models.SpanVSpanLbl, error) {
	rn := fmt.Sprintf("vsrc-%s", name)
	parentDn := fmt.Sprintf("uni/infra/vsrcgrp-%s", vsrc_grp)
	spanVSpanLbl := models.NewSpanVSpanLbl(rn, parentDn, description, spanVSpanLblattr)

	spanVSpanLbl.Status = "modified"
	err := sm.Save(spanVSpanLbl)
	return spanVSpanLbl, err
}


//func (sm *ServiceManager) ListSpanVSpanLbl(vsrc_grp string) ([]*models.SpanVSpanLbl, error) {
func (sm *ServiceManager) ListSpanVSpanLbl() ([]*models.SpanVSpanLbl, error) {


	//
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/spanSpanLbl.json", baseurlStr)
	//fmt.Println(dnUrl)
	cont, err := sm.GetViaURL(dnUrl)
	//list := models.SpanVSrcGrpListFromContainer(cont)
	//fmt.Println(cont)
//	os.Exit(1)
	//list := models.OpflexIDEpListFromContainer(cont)

	//


	//baseurlStr := "/api/node/class"
	//dnUrl := fmt.Sprintf("%s/uni/infra/vsrcgrp-%s.json", baseurlStr, vsrc_grp)
	//dnUrl := fmt.Sprintf("%s/vsrcgrp-%s/spanSpanLbl.json", baseurlStr, vsrc_grp)
	//dnUrl := fmt.Sprintf("%s/vsrcgrp-%s.json?query-target=children", baseurlStr ,vsrc_grp)

	//dnUrl := fmt.Sprintf("%s/spanSpanLbl.json)", baseurlStr)



//	cont, err := sm.GetViaURL(dnUrl)
//	fmt.Println(cont)
	list := models.SpanVSpanLblListFromContainer(cont)
	//for _,v := range(list) {
	//	fmt.Println(models.GetMOName(v.DistinguishedName))
	//}
	//fmt.Println(list)
	//os.Exit(1)

	return list, err
}

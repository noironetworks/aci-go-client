package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSpanVSrcGrp(name string, description string,
	spanVSrcGrpattr models.SpanVSrcGrpAttributes) (*models.SpanVSrcGrp, error) {
	rn := fmt.Sprintf("vsrcgrp-%s", name)
	parentDn := fmt.Sprintf("uni/infra")
	spanVSrcGrp := models.NewSpanVSrcGrp(rn, parentDn, description, spanVSrcGrpattr)
	err := sm.Save(spanVSrcGrp)
	return spanVSrcGrp, err
}


func (sm *ServiceManager) ReadSpanVSrcGrp(name string) (*models.SpanVSrcGrp, error) {
	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	spanVSrcGrp := models.SpanVSrcGrpFromContainer(cont)
	return spanVSrcGrp, nil
}


func (sm *ServiceManager) DeleteSpanVSrcGrp(name string) (error) {
	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s", name)
	return sm.DeleteByDn(dn, models.SpanVSrcGrpClassName)
}


func (sm *ServiceManager) UpdateSpanVSrcGrp(name string, description string,
	spanVSrcGrpattr models.SpanVSrcGrpAttributes) (*models.SpanVSrcGrp, error) {
	rn := fmt.Sprintf("%s", name)
	parentDn := fmt.Sprintf("uni/infra")
	spanVSrcGrp := models.NewSpanVSrcGrp(rn, parentDn, description, spanVSrcGrpattr)

	spanVSrcGrp.Status = "modified"
	err := sm.Save(spanVSrcGrp)
	return spanVSrcGrp, err
}

func (sm *ServiceManager) StateUpdateSpanVSrcGrp(name string, state string) (*models.SpanVSrcGrp, error) {

	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	spanVSrcGrp := models.SpanVSrcGrpFromContainer(cont)
	spanVSrcGrp.AdminSt = state
	err = sm.Save(spanVSrcGrp)
	return spanVSrcGrp, err
}

func (sm *ServiceManager) ListSpanVSrcGrp() ([]*models.SpanVSrcGrp, error) {

	baseurlStr := "/api/node/class"
	//dnUrl := fmt.Sprintf("%s/uni/infra/spanSrcGrp.json", baseurlStr)
	dnUrl := fmt.Sprintf("%s/spanVSrcGrp.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.SpanVSrcGrpListFromContainer(cont)

	return list, err
}

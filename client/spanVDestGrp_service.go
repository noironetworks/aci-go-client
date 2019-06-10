package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSpanVDestGrp(name string, description string,
	spanVDestGrpattr models.SpanVDestGrpAttributes) (*models.SpanVDestGrp, error) {
	rn := fmt.Sprintf("vdestgrp-%s", name)
	parentDn := fmt.Sprintf("uni/infra")
	spanVDestGrp := models.NewSpanVDestGrp(rn, parentDn, description, spanVDestGrpattr)
	err := sm.Save(spanVDestGrp)
	return spanVDestGrp, err
}


func (sm *ServiceManager) ReadSpanVDestGrp(name string) (*models.SpanVDestGrp, error) {
	dn := fmt.Sprintf("uni/infra/vdestgrp-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	spanVDestGrp := models.SpanVDestGrpFromContainer(cont)
	return spanVDestGrp, nil
}


func (sm *ServiceManager) DeleteSpanVDestGrp(name string) (error) {
	dn := fmt.Sprintf("uni/infra/vdestgrp-%s", name)
	return sm.DeleteByDn(dn, models.SpanVDestGrpClassName)
}


func (sm *ServiceManager) UpdateSpanVDestGrp(name string, description string,
	spanVDestGrpattr models.SpanVDestGrpAttributes) (*models.SpanVDestGrp, error) {
	rn := fmt.Sprintf("%s", name)
	parentDn := fmt.Sprintf("uni/infra")
	spanVDestGrp := models.NewSpanVDestGrp(rn, parentDn, description, spanVDestGrpattr)

	spanVDestGrp.Status = "modified"
	err := sm.Save(spanVDestGrp)
	return spanVDestGrp, err
}


func (sm *ServiceManager) ListSpanVDestGrp() ([]*models.SpanVDestGrp, error) {

	baseurlStr := "/api/node/class"
	//dnUrl := fmt.Sprintf("%s/uni/infra/spanDestGrp.json", baseurlStr)
	dnUrl := fmt.Sprintf("%s/spanVDestGrp.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.SpanVDestGrpListFromContainer(cont)

	return list, err
}

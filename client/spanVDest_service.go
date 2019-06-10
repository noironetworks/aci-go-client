package client

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSpanVDest(name string, vdest_grp string, description string,
	spanVDestattr models.SpanVDestAttributes) (*models.SpanVDest, error) {
	rn := fmt.Sprintf("vdest-%s", name)
	parentDn := fmt.Sprintf("uni/infra/vdestgrp-%s", vdest_grp)
	spanVDest := models.NewSpanVDest(rn, parentDn, description, spanVDestattr)
	err := sm.Save(spanVDest)
	return spanVDest, err
}


func (sm *ServiceManager) ReadSpanVDest(name string, vdest_grp string ) (*models.SpanVDest, error) {
	dn := fmt.Sprintf("uni/infra/vdestgrp-%s/vdest-%s", vdest_grp, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	spanVDest := models.SpanVDestFromContainer(cont)
	return spanVDest, nil
}


func (sm *ServiceManager) DeleteSpanVDest(name string, vdest_grp string ) (error) {
	dn := fmt.Sprintf("uni/infra/vdestgrp-%s/vdest-%s", vdest_grp, name)
	return sm.DeleteByDn(dn, models.SpanVDestClassName)
}


func (sm *ServiceManager) UpdateSpanVDest(name string, vdest_grp string, description string,
	spanVDestattr models.SpanVDestAttributes) (*models.SpanVDest, error) {
	rn := fmt.Sprintf("vdest-%s", name)
	parentDn := fmt.Sprintf("uni/infra/vdestgrp-%s", vdest_grp)
	spanVDest := models.NewSpanVDest(rn, parentDn, description, spanVDestattr)

	spanVDest.Status = "modified"
	err := sm.Save(spanVDest)
	return spanVDest, err
}


func (sm *ServiceManager) ListSpanVDest() ([]*models.SpanVDest, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/spanVDest.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.SpanVDestListFromContainer(cont)

	return list, err
}
package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSpanVEpgSummary(vdest_grp string, vdest string, description string,
	spanVEpgSummaryattr models.SpanVEpgSummaryAttributes) (*models.SpanVEpgSummary, error) {
	rn := fmt.Sprintf("vepgsummary")
	parentDn := fmt.Sprintf("uni/infra/vdestgrp-%s/vdest-%s", vdest_grp ,vdest)
	spanVEpgSummary := models.NewSpanVEpgSummary(rn, parentDn, description, spanVEpgSummaryattr)
	err := sm.Save(spanVEpgSummary)
	return spanVEpgSummary, err
}


func (sm *ServiceManager) ReadSpanVEpgSummary(vdest_grp string, vdest string ) (*models.SpanVEpgSummary, error) {
	dn := fmt.Sprintf("uni/infra/vdestgrp-%s/vdest-%s/vepgsummary", vdest_grp ,vdest)
	cont, err := sm.Get(dn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	spanVEpgSummary := models.SpanVEpgSummaryFromContainer(cont)
	return spanVEpgSummary, nil
}


func (sm *ServiceManager) DeleteSpanVEpgSummary(vdest_grp string, vdest string ) (error) {
	dn := fmt.Sprintf("uni/infra/vdestgrp-%s/vdest-%s/vepgsummary", vdest_grp ,vdest)
	return sm.DeleteByDn(dn, models.SpanVEpgSummaryClassName)
}


func (sm *ServiceManager) UpdateSpanVEpgSummary(vdest_grp string, vdest string, description string,
	spanVEpgSummaryattr models.SpanVEpgSummaryAttributes) (*models.SpanVEpgSummary, error) {
	rn := fmt.Sprintf("vepgsummary")
	parentDn := fmt.Sprintf("uni/infra/vdestgrp-%s/vdest-%s", vdest_grp ,vdest)
	spanVEpgSummary := models.NewSpanVEpgSummary(rn, parentDn, description, spanVEpgSummaryattr)

	spanVEpgSummary.Status = "modified"
	err := sm.Save(spanVEpgSummary)
	return spanVEpgSummary, err
}


func (sm *ServiceManager) ListSpanVEpgSummary(vdest_grp string, vdest string ) ([]*models.SpanVEpgSummary, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/infra/vdestgrp-%s/vdest-%s.json", baseurlStr, vdest_grp, vdest)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.SpanVEpgSummaryListFromContainer(cont)

	return list, err
}

package client

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/container"

	"github.com/ciscoecosystem/aci-go-client/models"
)









func (sm *ServiceManager) CreateSpanVSrc(name string, vsrc_grp string, description string,
	spanVSrcattr models.SpanVSrcAttributes) (*models.SpanVSrc, error) {
	rn := fmt.Sprintf("vsrc-%s", name)
	parentDn := fmt.Sprintf("uni/infra/vsrcgrp-%s", vsrc_grp)
	spanVSrc := models.NewSpanVSrc(rn, parentDn, description, spanVSrcattr)
	err := sm.Save(spanVSrc)
	return spanVSrc, err
}


func (sm *ServiceManager) ReadSpanVSrc(name string, vsrc_grp string ) (*models.SpanVSrc, error) {
	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s/vsrc-%s", vsrc_grp, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	spanVSrc := models.SpanVSrcFromContainer(cont)
	return spanVSrc, nil
}


func (sm *ServiceManager) DeleteSpanVSrc(name string, vsrc_grp string ) (error) {
	dn := fmt.Sprintf("uni/infra/vsrcgrp-%s/vsrc-%s", vsrc_grp, name)
	return sm.DeleteByDn(dn, models.SpanVSrcClassName)
}


func (sm *ServiceManager) UpdateSpanVSrc(name string, vsrc_grp string, description string,
	spanVSrcattr models.SpanVSrcAttributes) (*models.SpanVSrc, error) {
	rn := fmt.Sprintf("vsrc-%s", name)
	parentDn := fmt.Sprintf("uni/infra/vsrcgrp-%s", vsrc_grp)
	spanVSrc := models.NewSpanVSrc(rn, parentDn, description, spanVSrcattr)

	spanVSrc.Status = "modified"
	err := sm.Save(spanVSrc)
	return spanVSrc, err
}


func (sm *ServiceManager) ListSpanVSrc(vsrc_grp string) ([]*models.SpanVSrc, error) {

	baseurlStr := "/api/node/class"
	//dnUrl := fmt.Sprintf("%s/uni/infra/vsrcgrp-%s.json", baseurlStr, vsrc_grp)
	dnUrl := fmt.Sprintf("%s/spanVSrc.json?query-target-filter=and(wcard(spanVSrc.dn,\"%s\"))", baseurlStr, vsrc_grp)
	//dnUrl := fmt.Sprintf("%s/spanSpanLbl.json?query-target-filter=and(wcard(spanSpanLbl.dn,\"%s\"))", baseurlStr, vsrc_grp)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.SpanVSrcListFromContainer(cont)

	return list, err
}

//spanRsSrcToPathEp --> topology/pod-1/paths-101/pathep-[eth1/1] (pod: 1, node: "101", port: eth1/1)
//spanRsSrcToVPort --> uni/tn-kube/ap-kubernetes/epg-kube-nodes/cep-88:1D:FC:F2:FB:59 (tenant: "t1", ap: "ap1", epg: "myepg", cep: "88:1D:FC:F2:FB:59")
//spanRsSrcToEpg --> uni/tn-kube/ap-kubernetes/epg-kube-default (tenant: "t1", ap: "ap1", epg: "myepg")

func (sm *ServiceManager) CreateRelationSpanRsSrcToPathEp( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rssrcToPathEp-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"
			}
		}
	}`, "spanRsSrcToPathEp", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	_ = cont
	//fmt.Printf("%+v", cont)

	return nil
}


func (sm *ServiceManager) CreateRelationSpanRsSrcToVPort( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rssrcToVPort-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"
			}
		}
	}`, "spanRsSrcToVPort", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	_ = cont
	//fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) CreateRelationSpanRsSrcToEpg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rssrcToEpg", parentDn)
	//tdn := fmt.Sprintf("%s/rssrcToVPort-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tDn": "%s"
			}
		}
	}`, "spanRsSrcToEpg", dn, tDn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	_ = cont
	//fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) ReadRelationSpanRsSrcToVPort( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/%s/%s.json",baseurlStr,parentDn,"fvRsBDSubnetToProfile")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsBDSubnetToProfile")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnRtctrlProfileName")
		return dat, err
	} else {
		return nil,err
	}






}
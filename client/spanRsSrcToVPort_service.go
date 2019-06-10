package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)







func (sm *ServiceManager) ListSpanRsSrcToVPort() ([]*models.SpanRsSrcToVPort, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/spanRsSrcToVPort.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	//list := models.SpanVSrcGrpListFromContainer(cont)
	list := models.SpanRsSrcToVPortListFromContainer(cont)


	return list, err
}

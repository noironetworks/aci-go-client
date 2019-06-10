package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)







func (sm *ServiceManager) ListOpflexIDEp() ([]*models.OpflexIDEp, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/opflexIDEp.json", baseurlStr)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.OpflexIDEpListFromContainer(cont)


	return list, err
}

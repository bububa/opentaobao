package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicestoredeleteservicestorecoverservice 删除网点覆盖的服务
// tmall.servicecenter.servicestore.deleteservicestorecoverservice
//
// 天猫服务平台删除网点覆盖的服务，
// 必选字段：serviceStoreCode、bizType
func Tmallservicecenterservicestoredeleteservicestorecoverservice(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicestoredeleteservicestorecoverserviceAPIRequest, session string) (*tmallservice.TmallservicecenterservicestoredeleteservicestorecoverserviceAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicestoredeleteservicestorecoverserviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

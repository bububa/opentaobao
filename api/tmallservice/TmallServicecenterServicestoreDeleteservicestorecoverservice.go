package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreDeleteservicestorecoverservice 删除网点覆盖的服务
// tmall.servicecenter.servicestore.deleteservicestorecoverservice
//
// 天猫服务平台删除网点覆盖的服务，
// 必选字段：serviceStoreCode、bizType
func TmallServicecenterServicestoreDeleteservicestorecoverservice(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest, resp *tmallservice.TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

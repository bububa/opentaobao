package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreDeleteservicestorecoverservice 删除网点覆盖的服务
// tmall.servicecenter.servicestore.deleteservicestorecoverservice
//
// 天猫服务平台删除网点覆盖的服务，
// 必选字段：serviceStoreCode、bizType
func TmallServicecenterServicestoreDeleteservicestorecoverservice(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIRequest, resp *tmallservice.TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

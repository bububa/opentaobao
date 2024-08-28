package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterServicestoreDeleteservicestorecapacity 删除网点容量
// tmall.servicecenter.servicestore.deleteservicestorecapacity
//
// 删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
// 必选字段：serviceStoreCode、bizType
func TmallServicecenterServicestoreDeleteservicestorecapacity(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterServicestoreDeleteservicestorecapacityAPIRequest, resp *tmallsc.TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

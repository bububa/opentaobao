package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterServicestoreCreateservicestorecoverservice 新增网点覆盖的服务
// tmall.servicecenter.servicestore.createservicestorecoverservice
//
// 新增网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要新增的网点覆盖的服务已存在，会新增失败。
// 网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
func TmallServicecenterServicestoreCreateservicestorecoverservice(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterServicestoreCreateservicestorecoverserviceAPIRequest, resp *tmallsc.TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

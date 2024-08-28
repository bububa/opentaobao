package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerRegister 服务商添加工人
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
func AlibabaSscSupplyplatformServiceworkerRegister(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

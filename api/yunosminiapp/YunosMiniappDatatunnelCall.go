package yunosminiapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosminiapp"
)

// YunosMiniappDatatunnelCall 车载小程序外部服务调用
// yunos.miniapp.datatunnel.call
//
// 对客户提供的api进行统一封装调用。
func YunosMiniappDatatunnelCall(ctx context.Context, clt *core.SDKClient, req *yunosminiapp.YunosMiniappDatatunnelCallAPIRequest, resp *yunosminiapp.YunosMiniappDatatunnelCallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

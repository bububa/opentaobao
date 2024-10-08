package yunos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunos"
)

// AliyunCosmoGatewayInvoke alios cosmo服务调用
// aliyun.cosmo.gateway.invoke
//
// AliOS cosmo服务分发平台对外调用接口
func AliyunCosmoGatewayInvoke(ctx context.Context, clt *core.SDKClient, req *yunos.AliyunCosmoGatewayInvokeAPIRequest, resp *yunos.AliyunCosmoGatewayInvokeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package rhino

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// TaobaoRhinoCrmGateway crm实体变更回调接口
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
func TaobaoRhinoCrmGateway(ctx context.Context, clt *core.SDKClient, req *rhino.TaobaoRhinoCrmGatewayAPIRequest, resp *rhino.TaobaoRhinoCrmGatewayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

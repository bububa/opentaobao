package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractIsvGateway isv调用gateway
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
func AlibabaInteractIsvGateway(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractIsvGatewayAPIRequest, resp *interact.AlibabaInteractIsvGatewayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

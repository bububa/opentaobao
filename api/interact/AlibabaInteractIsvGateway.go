package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractIsvGateway isv调用gateway
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
func AlibabaInteractIsvGateway(clt *core.SDKClient, req *interact.AlibabaInteractIsvGatewayAPIRequest, resp *interact.AlibabaInteractIsvGatewayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

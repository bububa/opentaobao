package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractisvgateway isv调用gateway
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
func Alibabainteractisvgateway(clt *core.SDKClient, req *interact.AlibabainteractisvgatewayAPIRequest, session string) (*interact.AlibabainteractisvgatewayAPIResponse, error) {
	var resp interact.AlibabainteractisvgatewayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

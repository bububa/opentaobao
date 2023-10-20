package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorauthorize 客户端授权页
// alibaba.interact.sensor.authorize
//
// 客户端授权页
func Alibabainteractsensorauthorize(clt *core.SDKClient, req *interact.AlibabainteractsensorauthorizeAPIRequest, session string) (*interact.AlibabainteractsensorauthorizeAPIResponse, error) {
	var resp interact.AlibabainteractsensorauthorizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

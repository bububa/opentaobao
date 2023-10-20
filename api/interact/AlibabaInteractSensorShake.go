package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorshake 摇一摇
// alibaba.interact.sensor.shake
//
// 摇一摇
func Alibabainteractsensorshake(clt *core.SDKClient, req *interact.AlibabainteractsensorshakeAPIRequest, session string) (*interact.AlibabainteractsensorshakeAPIResponse, error) {
	var resp interact.AlibabainteractsensorshakeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

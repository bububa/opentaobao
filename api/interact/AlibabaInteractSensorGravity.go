package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorgravity 重力感应
// alibaba.interact.sensor.gravity
//
// native获取重力感应
func Alibabainteractsensorgravity(clt *core.SDKClient, req *interact.AlibabainteractsensorgravityAPIRequest, session string) (*interact.AlibabainteractsensorgravityAPIResponse, error) {
	var resp interact.AlibabainteractsensorgravityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

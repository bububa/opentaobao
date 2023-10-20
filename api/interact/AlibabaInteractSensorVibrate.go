package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorvibrate 客户端震动
// alibaba.interact.sensor.vibrate
//
// 客户端震动
func Alibabainteractsensorvibrate(clt *core.SDKClient, req *interact.AlibabainteractsensorvibrateAPIRequest, session string) (*interact.AlibabainteractsensorvibrateAPIResponse, error) {
	var resp interact.AlibabainteractsensorvibrateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

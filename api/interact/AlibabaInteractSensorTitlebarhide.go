package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensortitlebarhide 隐藏titleBar
// alibaba.interact.sensor.titlebarhide
//
// 隐藏titleBar
func Alibabainteractsensortitlebarhide(clt *core.SDKClient, req *interact.AlibabainteractsensortitlebarhideAPIRequest, session string) (*interact.AlibabainteractsensortitlebarhideAPIResponse, error) {
	var resp interact.AlibabainteractsensortitlebarhideAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

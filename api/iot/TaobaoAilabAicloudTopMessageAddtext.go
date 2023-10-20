package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopmessageaddtext 精灵代说
// taobao.ailab.aicloud.top.message.addtext
//
// 精灵代说
func Taobaoailabaicloudtopmessageaddtext(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopmessageaddtextAPIRequest, session string) (*iot.TaobaoailabaicloudtopmessageaddtextAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopmessageaddtextAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

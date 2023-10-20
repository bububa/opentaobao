package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopmessagelist 获取留言列表
// taobao.ailab.aicloud.top.message.list
//
// 根据指定参数获取留言列表
func Taobaoailabaicloudtopmessagelist(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopmessagelistAPIRequest, session string) (*iot.TaobaoailabaicloudtopmessagelistAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopmessagelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

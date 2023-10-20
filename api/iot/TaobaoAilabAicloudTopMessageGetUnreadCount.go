package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopmessagegetunreadcount 获取未读的消息数量
// taobao.ailab.aicloud.top.message.get.unread.count
//
// 开放获取未读留言数量的接口
func Taobaoailabaicloudtopmessagegetunreadcount(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopmessagegetunreadcountAPIRequest, session string) (*iot.TaobaoailabaicloudtopmessagegetunreadcountAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopmessagegetunreadcountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

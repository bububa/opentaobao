package alimsg

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimsg"
)

// Alibabaidleordermsgsend 虚拟发货消息发送接口
// alibaba.idle.order.msg.send
//
// 用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
func Alibabaidleordermsgsend(clt *core.SDKClient, req *alimsg.AlibabaidleordermsgsendAPIRequest, session string) (*alimsg.AlibabaidleordermsgsendAPIResponse, error) {
	var resp alimsg.AlibabaidleordermsgsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

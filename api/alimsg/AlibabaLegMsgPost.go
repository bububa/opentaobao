package alimsg

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimsg"
)

// Alibabalegmsgpost 集团法务消息发送
// alibaba.leg.msg.post
//
// 消息发送能力
func Alibabalegmsgpost(clt *core.SDKClient, req *alimsg.AlibabalegmsgpostAPIRequest, session string) (*alimsg.AlibabalegmsgpostAPIResponse, error) {
	var resp alimsg.AlibabalegmsgpostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

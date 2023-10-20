package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcmessagesconfirm 确认消费消息的状态
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
func Taobaotmcmessagesconfirm(clt *core.SDKClient, req *tmc.TaobaotmcmessagesconfirmAPIRequest, session string) (*tmc.TaobaotmcmessagesconfirmAPIResponse, error) {
	var resp tmc.TaobaotmcmessagesconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

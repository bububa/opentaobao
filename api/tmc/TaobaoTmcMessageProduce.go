package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcmessageproduce 发布单条消息
// taobao.tmc.message.produce
//
// 发布单条消息
func Taobaotmcmessageproduce(clt *core.SDKClient, req *tmc.TaobaotmcmessageproduceAPIRequest, session string) (*tmc.TaobaotmcmessageproduceAPIResponse, error) {
	var resp tmc.TaobaotmcmessageproduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

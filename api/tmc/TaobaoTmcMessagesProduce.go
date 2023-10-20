package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcmessagesproduce 批量发送消息
// taobao.tmc.messages.produce
//
// 批量发送消息
func Taobaotmcmessagesproduce(clt *core.SDKClient, req *tmc.TaobaotmcmessagesproduceAPIRequest, session string) (*tmc.TaobaotmcmessagesproduceAPIResponse, error) {
	var resp tmc.TaobaotmcmessagesproduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

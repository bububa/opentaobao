package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcMessageProduce 发布单条消息
// taobao.tmc.message.produce
//
// 发布单条消息
func TaobaoTmcMessageProduce(clt *core.SDKClient, req *tmc.TaobaoTmcMessageProduceAPIRequest, session string) (*tmc.TaobaoTmcMessageProduceAPIResponse, error) {
	var resp tmc.TaobaoTmcMessageProduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

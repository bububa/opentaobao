package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoqimeneventsproduce 批量发送奇门事件
// taobao.qimen.events.produce
//
// 批量发送消息
func Taobaoqimeneventsproduce(clt *core.SDKClient, req *util.TaobaoqimeneventsproduceAPIRequest, session string) (*util.TaobaoqimeneventsproduceAPIResponse, error) {
	var resp util.TaobaoqimeneventsproduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

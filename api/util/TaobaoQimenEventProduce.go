package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoqimeneventproduce 发出奇门事件
// taobao.qimen.event.produce
//
// 当订单被处理时，用于通知奇门系统。
func Taobaoqimeneventproduce(clt *core.SDKClient, req *util.TaobaoqimeneventproduceAPIRequest, session string) (*util.TaobaoqimeneventproduceAPIResponse, error) {
	var resp util.TaobaoqimeneventproduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

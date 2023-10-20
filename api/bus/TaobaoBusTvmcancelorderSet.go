package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobustvmcancelorderset 线下自助机未付款取消订单
// taobao.bus.tvmcancelorder.set
//
// 自助机汽车票未付款取消订单
func Taobaobustvmcancelorderset(clt *core.SDKClient, req *bus.TaobaobustvmcancelordersetAPIRequest, session string) (*bus.TaobaobustvmcancelordersetAPIResponse, error) {
	var resp bus.TaobaobustvmcancelordersetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

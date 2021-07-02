package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsOrderShengxianConfirm 物流宝生鲜冷链的发货
// taobao.logistics.order.shengxian.confirm
//
// 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
func TaobaoLogisticsOrderShengxianConfirm(clt *core.SDKClient, req *logistic.TaobaoLogisticsOrderShengxianConfirmAPIRequest, session string) (*logistic.TaobaoLogisticsOrderShengxianConfirmAPIResponse, error) {
	var resp logistic.TaobaoLogisticsOrderShengxianConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* TaobaoLogisticsOfflineSend
自己联系物流（线下物流）发货
taobao.logistics.offline.send

用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
func TaobaoLogisticsOfflineSend(clt *core.SDKClient, req *logistic.TaobaoLogisticsOfflineSendAPIRequest, session string) (*logistic.TaobaoLogisticsOfflineSendAPIResponse, error) {
	var resp logistic.TaobaoLogisticsOfflineSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

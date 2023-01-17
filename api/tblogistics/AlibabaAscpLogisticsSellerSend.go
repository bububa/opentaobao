package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerSend 商家配送发货
// alibaba.ascp.logistics.seller.send
//
// 该API提供商家配送发货能力，使用该接口发货，交易订单状态会直接变成卖家已发货
func AlibabaAscpLogisticsSellerSend(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerSendAPIRequest, session string) (*tblogistics.AlibabaAscpLogisticsSellerSendAPIResponse, error) {
	var resp tblogistics.AlibabaAscpLogisticsSellerSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

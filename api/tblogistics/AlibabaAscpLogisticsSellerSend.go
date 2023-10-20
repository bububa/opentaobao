package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticssellersend 商家配送发货
// alibaba.ascp.logistics.seller.send
//
// 该API提供商家配送发货能力，使用该接口发货，交易订单状态会直接变成卖家已发货
func Alibabaascplogisticssellersend(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticssellersendAPIRequest, session string) (*tblogistics.AlibabaascplogisticssellersendAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticssellersendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

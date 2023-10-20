package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlineDeliveryorderGet 同城配送在线下单获取配送单
// alibaba.ascp.logistics.instantsonline.deliveryorder.get
//
// 同城配送在线下单获取配送单
func AlibabaAscpLogisticsInstantsonlineDeliveryorderGet(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIRequest, session string) (*tblogistics.AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse, error) {
	var resp tblogistics.AlibabaAscpLogisticsInstantsonlineDeliveryorderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

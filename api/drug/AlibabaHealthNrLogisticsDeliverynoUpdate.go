package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaHealthNrLogisticsDeliverynoUpdate 上传订单同城快递单号
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
func AlibabaHealthNrLogisticsDeliverynoUpdate(clt *core.SDKClient, req *drug.AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest, resp *drug.AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

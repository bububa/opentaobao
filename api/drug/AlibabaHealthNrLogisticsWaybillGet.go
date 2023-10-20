package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaHealthNrLogisticsWaybillGet 电子面单查询接口
// alibaba.health.nr.logistics.waybill.get
//
// 商家登录后根据订单号查询物流单号及电子面单信息
func AlibabaHealthNrLogisticsWaybillGet(clt *core.SDKClient, req *drug.AlibabaHealthNrLogisticsWaybillGetAPIRequest, resp *drug.AlibabaHealthNrLogisticsWaybillGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

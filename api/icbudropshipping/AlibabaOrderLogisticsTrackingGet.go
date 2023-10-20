package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaOrderLogisticsTrackingGet 阿里巴巴订单物流轨迹查询
// alibaba.order.logistics.tracking.get
//
// 阿里巴巴订单物流轨迹查询
func AlibabaOrderLogisticsTrackingGet(clt *core.SDKClient, req *icbudropshipping.AlibabaOrderLogisticsTrackingGetAPIRequest, resp *icbudropshipping.AlibabaOrderLogisticsTrackingGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

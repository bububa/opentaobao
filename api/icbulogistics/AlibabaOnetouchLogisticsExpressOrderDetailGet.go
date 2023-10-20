package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressOrderDetailGet 订单详细信息(面单及仓库信息)
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
func AlibabaOnetouchLogisticsExpressOrderDetailGet(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

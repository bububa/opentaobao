package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressLogisticsOrderCreate 快递下单
// alibaba.onetouch.logistics.express.logistics.order.create
//
// 快递下单
func AlibabaOnetouchLogisticsExpressLogisticsOrderCreate(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

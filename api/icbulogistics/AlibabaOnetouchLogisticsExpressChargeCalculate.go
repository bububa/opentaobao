package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressChargeCalculate 计算快递运费&下单参数校验
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
func AlibabaOnetouchLogisticsExpressChargeCalculate(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

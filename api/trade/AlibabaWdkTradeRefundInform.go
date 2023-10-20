package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeRefundInform 外部渠道通知淘鲜达退款成功接口
// alibaba.wdk.trade.refund.inform
//
// 该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
func AlibabaWdkTradeRefundInform(clt *core.SDKClient, req *trade.AlibabaWdkTradeRefundInformAPIRequest, resp *trade.AlibabaWdkTradeRefundInformAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

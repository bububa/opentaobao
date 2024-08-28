package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeRefundInform 外部渠道通知淘鲜达退款成功接口
// alibaba.wdk.trade.refund.inform
//
// 该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
func AlibabaWdkTradeRefundInform(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaWdkTradeRefundInformAPIRequest, resp *trade.AlibabaWdkTradeRefundInformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

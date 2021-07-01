package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeRefundInformAPIRequest
外部渠道通知淘鲜达退款成功接口 API请求
alibaba.wdk.trade.refund.inform

该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。 */
type AlibabaWdkTradeRefundInformAPIRequest struct {
	model.Params
	// 通知退款成功请求
	_informRefundSuccessRequest *InformRefundSuccessRequest
}

// New

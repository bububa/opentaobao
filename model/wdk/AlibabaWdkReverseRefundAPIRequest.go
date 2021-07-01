package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkReverseRefundAPIRequest
退款打款 API请求
alibaba.wdk.reverse.refund

五道口退款打款开放能力接口 */
type AlibabaWdkReverseRefundAPIRequest struct {
	model.Params
	// 退款打款请求
	_openRefundReqDTO *OpenRefundReqDto
}

// New

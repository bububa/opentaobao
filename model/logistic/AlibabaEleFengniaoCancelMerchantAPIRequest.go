package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoCancelMerchantAPIRequest
商户取消 API请求
alibaba.ele.fengniao.cancel.merchant

商户取消配送 */
type AlibabaEleFengniaoCancelMerchantAPIRequest struct {
	model.Params
	// 参数param
	_param *Param
}

// New

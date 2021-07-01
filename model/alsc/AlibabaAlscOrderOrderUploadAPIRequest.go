package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscOrderOrderUploadAPIRequest
订单回流 API请求
alibaba.alsc.order.order.upload

第三方订单回流 */
type AlibabaAlscOrderOrderUploadAPIRequest struct {
	model.Params
	// 订单回流参数
	_paramBackflowRequest *BackflowRequest
}

// New

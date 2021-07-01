package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthShortUrlGetAPIRequest
支付宝短链跳转三方h5通用接口 API请求
alibaba.alihealth.short.url.get

支付宝短链跳转三方h5通用接口 */
type AlibabaAlihealthShortUrlGetAPIRequest struct {
	model.Params
	// 三方h5
	_url string
	// 参数替换列表
	_params []string
}

// New

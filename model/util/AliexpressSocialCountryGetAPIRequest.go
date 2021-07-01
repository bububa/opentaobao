package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialCountryGetAPIRequest
获取国家列表 API请求
aliexpress.social.country.get

获取目前AE支持的国家列表 */
type AliexpressSocialCountryGetAPIRequest struct {
	model.Params
	// 语言
	_language string
}

// New

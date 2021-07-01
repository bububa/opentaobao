package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialCurrencyGetAPIRequest
币种获取接口 API请求
aliexpress.social.currency.get

获取目前AE社交支持的币种 */
type AliexpressSocialCurrencyGetAPIRequest struct {
	model.Params
}

// New

package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialDiscategoryGetAPIRequest
展示类目获取接口 API请求
aliexpress.social.discategory.get

AE展示类目获取接口 */
type AliexpressSocialDiscategoryGetAPIRequest struct {
	model.Params
	// Locale值，格式为language+"_"+country
	_locale string
}

// New

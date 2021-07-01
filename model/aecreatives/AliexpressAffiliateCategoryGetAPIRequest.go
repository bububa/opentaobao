package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateCategoryGetAPIRequest
AE流量推广类目信息获取API API请求
aliexpress.affiliate.category.get

获取AE流量推广类目的API */
type AliexpressAffiliateCategoryGetAPIRequest struct {
	model.Params
	// 请求安全签名
	_appSignature string
}

// New

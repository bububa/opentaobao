package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDropshippingTokenCreateAPIRequest
国际站dropshipping 选品token 创建 API请求
alibaba.dropshipping.token.create

国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆 */
type AlibabaDropshippingTokenCreateAPIRequest struct {
	model.Params
}

// New

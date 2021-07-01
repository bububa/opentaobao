package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialItemPromotionAPIRequest
获取推广链接 API请求
aliexpress.social.item.promotion

获取商品社交推广链接 */
type AliexpressSocialItemPromotionAPIRequest struct {
	model.Params
	// 推广的商品链接
	_targetUrl string
	// 子渠道号
	_af string
	// campaign Id
	_cn string
	// creative id
	_cv string
	// click id
	_dp string
}

// New

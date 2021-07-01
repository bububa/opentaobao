package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTaobaoShopCatNeoGetAPIRequest
店铺mtop接口鉴权虚拟api API请求
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用 */
type AlibabaTaobaoShopCatNeoGetAPIRequest struct {
	model.Params
	// 客户端鉴权虚拟api
	_unNamed string
}

// New

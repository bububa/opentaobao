package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDataItemGetAPIRequest
获取商品 API请求
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api */
type AlibabaDataItemGetAPIRequest struct {
	model.Params
	// 获取商品信息，作为客户端Weex鉴权的虚拟api
	_unNamed string
}

// New

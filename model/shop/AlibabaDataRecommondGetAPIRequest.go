package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDataRecommondGetAPIRequest
获取推荐信息 API请求
alibaba.data.recommond.get

获取优惠券信息，仅作客户端鉴权虚拟api使用 */
type AlibabaDataRecommondGetAPIRequest struct {
	model.Params
	// 客户端鉴权虚拟api使用
	_unNamed string
}

// New

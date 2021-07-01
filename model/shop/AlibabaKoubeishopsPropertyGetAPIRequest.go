package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKoubeishopsPropertyGetAPIRequest
口碑店铺列表推荐 API请求
alibaba.koubeishops.property.get

推荐用户附近的美食门店 */
type AlibabaKoubeishopsPropertyGetAPIRequest struct {
	model.Params
	// 入参
	_paramOpenApiSearchRequest *OpenApiSearchRequest
}

// New

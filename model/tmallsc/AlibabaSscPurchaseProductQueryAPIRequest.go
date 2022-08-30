package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscPurchaseProductQueryAPIRequest 查询已采购的服务产品 API请求
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
type AlibabaSscPurchaseProductQueryAPIRequest struct {
	model.Params
}

// NewAlibabaSscPurchaseProductQueryRequest 初始化AlibabaSscPurchaseProductQueryAPIRequest对象
func NewAlibabaSscPurchaseProductQueryRequest() *AlibabaSscPurchaseProductQueryAPIRequest {
	return &AlibabaSscPurchaseProductQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscPurchaseProductQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.purchase.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscPurchaseProductQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

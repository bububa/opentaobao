package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripshipproductsyncbaseAPIRequest 基础信息修改回调 API请求
// alitrip.ship.product.syncbase
//
// 基础信息修改回调
type AlitripshipproductsyncbaseAPIRequest struct {
	model.Params
}

// NewAlitripshipproductsyncbaseRequest 初始化AlitripshipproductsyncbaseAPIRequest对象
func NewAlitripshipproductsyncbaseRequest() *AlitripshipproductsyncbaseAPIRequest {
	return &AlitripshipproductsyncbaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripshipproductsyncbaseAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.product.syncbase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripshipproductsyncbaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripshipproductsyncbaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

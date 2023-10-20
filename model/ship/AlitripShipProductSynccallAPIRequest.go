package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripshipproductsynccallAPIRequest 全量同步回调 API请求
// alitrip.ship.product.synccall
//
// 全量同步接口
type AlitripshipproductsynccallAPIRequest struct {
	model.Params
}

// NewAlitripshipproductsynccallRequest 初始化AlitripshipproductsynccallAPIRequest对象
func NewAlitripshipproductsynccallRequest() *AlitripshipproductsynccallAPIRequest {
	return &AlitripshipproductsynccallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripshipproductsynccallAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.product.synccall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripshipproductsynccallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripshipproductsynccallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

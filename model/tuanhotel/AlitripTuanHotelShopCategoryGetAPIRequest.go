package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptuanhotelshopcategorygetAPIRequest 商家店铺类目查询 API请求
// alitrip.tuan.hotel.shop.category.get
//
// 查询商家店铺类目信息
type AlitriptuanhotelshopcategorygetAPIRequest struct {
	model.Params
}

// NewAlitriptuanhotelshopcategorygetRequest 初始化AlitriptuanhotelshopcategorygetAPIRequest对象
func NewAlitriptuanhotelshopcategorygetRequest() *AlitriptuanhotelshopcategorygetAPIRequest {
	return &AlitriptuanhotelshopcategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptuanhotelshopcategorygetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.shop.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptuanhotelshopcategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptuanhotelshopcategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

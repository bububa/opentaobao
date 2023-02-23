package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelShopCategoryGetAPIRequest 商家店铺类目查询 API请求
// alitrip.tuan.hotel.shop.category.get
//
// 查询商家店铺类目信息
type AlitripTuanHotelShopCategoryGetAPIRequest struct {
	model.Params
}

// NewAlitripTuanHotelShopCategoryGetRequest 初始化AlitripTuanHotelShopCategoryGetAPIRequest对象
func NewAlitripTuanHotelShopCategoryGetRequest() *AlitripTuanHotelShopCategoryGetAPIRequest {
	return &AlitripTuanHotelShopCategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelShopCategoryGetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.shop.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelShopCategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTuanHotelShopCategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

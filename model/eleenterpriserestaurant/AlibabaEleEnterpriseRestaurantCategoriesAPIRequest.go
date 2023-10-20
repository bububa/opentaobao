package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriserestaurantcategoriesAPIRequest 餐厅分类 API请求
// alibaba.ele.enterprise.restaurant.categories
//
// 餐厅分类
type AlibabaeleenterpriserestaurantcategoriesAPIRequest struct {
	model.Params
}

// NewAlibabaeleenterpriserestaurantcategoriesRequest 初始化AlibabaeleenterpriserestaurantcategoriesAPIRequest对象
func NewAlibabaeleenterpriserestaurantcategoriesRequest() *AlibabaeleenterpriserestaurantcategoriesAPIRequest {
	return &AlibabaeleenterpriserestaurantcategoriesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriserestaurantcategoriesAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.categories"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriserestaurantcategoriesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriserestaurantcategoriesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

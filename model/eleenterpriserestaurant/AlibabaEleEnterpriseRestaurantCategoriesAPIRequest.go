package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantCategoriesAPIRequest 餐厅分类 API请求
// alibaba.ele.enterprise.restaurant.categories
//
// 餐厅分类
type AlibabaEleEnterpriseRestaurantCategoriesAPIRequest struct {
	model.Params
}

// NewAlibabaEleEnterpriseRestaurantCategoriesRequest 初始化AlibabaEleEnterpriseRestaurantCategoriesAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantCategoriesRequest() *AlibabaEleEnterpriseRestaurantCategoriesAPIRequest {
	return &AlibabaEleEnterpriseRestaurantCategoriesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.categories"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

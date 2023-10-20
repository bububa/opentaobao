package eleenterpriserestaurant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) Reset() {
	r.Params.ToZero()
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

var poolAlibabaEleEnterpriseRestaurantCategoriesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseRestaurantCategoriesRequest()
	},
}

// GetAlibabaEleEnterpriseRestaurantCategoriesRequest 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantCategoriesAPIRequest
func GetAlibabaEleEnterpriseRestaurantCategoriesAPIRequest() *AlibabaEleEnterpriseRestaurantCategoriesAPIRequest {
	return poolAlibabaEleEnterpriseRestaurantCategoriesAPIRequest.Get().(*AlibabaEleEnterpriseRestaurantCategoriesAPIRequest)
}

// ReleaseAlibabaEleEnterpriseRestaurantCategoriesAPIRequest 将 AlibabaEleEnterpriseRestaurantCategoriesAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantCategoriesAPIRequest(v *AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantCategoriesAPIRequest.Put(v)
}

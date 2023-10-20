package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// Alibabaeleenterpriserestaurantcategories 餐厅分类
// alibaba.ele.enterprise.restaurant.categories
//
// 餐厅分类
func Alibabaeleenterpriserestaurantcategories(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaeleenterpriserestaurantcategoriesAPIRequest, session string) (*eleenterpriserestaurant.AlibabaeleenterpriserestaurantcategoriesAPIResponse, error) {
	var resp eleenterpriserestaurant.AlibabaeleenterpriserestaurantcategoriesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

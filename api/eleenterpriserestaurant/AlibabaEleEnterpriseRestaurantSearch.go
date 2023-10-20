package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// Alibabaeleenterpriserestaurantsearch 餐厅列表
// alibaba.ele.enterprise.restaurant.search
//
// 餐厅列表
func Alibabaeleenterpriserestaurantsearch(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaeleenterpriserestaurantsearchAPIRequest, session string) (*eleenterpriserestaurant.AlibabaeleenterpriserestaurantsearchAPIResponse, error) {
	var resp eleenterpriserestaurant.AlibabaeleenterpriserestaurantsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

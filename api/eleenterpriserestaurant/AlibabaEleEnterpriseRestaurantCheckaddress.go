package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// Alibabaeleenterpriserestaurantcheckaddress 检查地址是否在餐厅配送范围内
// alibaba.ele.enterprise.restaurant.checkaddress
//
// 检查地址是否在餐厅配送范围内
func Alibabaeleenterpriserestaurantcheckaddress(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaeleenterpriserestaurantcheckaddressAPIRequest, session string) (*eleenterpriserestaurant.AlibabaeleenterpriserestaurantcheckaddressAPIResponse, error) {
	var resp eleenterpriserestaurant.AlibabaeleenterpriserestaurantcheckaddressAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

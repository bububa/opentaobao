package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// Alibabaeleenterpriserestaurantmix 混合搜索店铺
// alibaba.ele.enterprise.restaurant.mix
//
// 混合搜索店铺
func Alibabaeleenterpriserestaurantmix(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaeleenterpriserestaurantmixAPIRequest, session string) (*eleenterpriserestaurant.AlibabaeleenterpriserestaurantmixAPIResponse, error) {
	var resp eleenterpriserestaurant.AlibabaeleenterpriserestaurantmixAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// Alibabaeleenterpriserestaurantget 查询餐厅信息
// alibaba.ele.enterprise.restaurant.get
//
// 查询餐厅信息
func Alibabaeleenterpriserestaurantget(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaeleenterpriserestaurantgetAPIRequest, session string) (*eleenterpriserestaurant.AlibabaeleenterpriserestaurantgetAPIResponse, error) {
	var resp eleenterpriserestaurant.AlibabaeleenterpriserestaurantgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

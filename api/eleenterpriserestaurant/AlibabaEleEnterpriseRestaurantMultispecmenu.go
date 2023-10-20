package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// Alibabaeleenterpriserestaurantmultispecmenu 查询餐厅菜单
// alibaba.ele.enterprise.restaurant.multispecmenu
//
// 查询餐厅菜单
func Alibabaeleenterpriserestaurantmultispecmenu(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaeleenterpriserestaurantmultispecmenuAPIRequest, session string) (*eleenterpriserestaurant.AlibabaeleenterpriserestaurantmultispecmenuAPIResponse, error) {
	var resp eleenterpriserestaurant.AlibabaeleenterpriserestaurantmultispecmenuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

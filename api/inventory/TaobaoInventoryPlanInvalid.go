package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Taobaoinventoryplaninvalid 失效计划库存
// taobao.inventory.plan.invalid
//
// 计划库存的失效服务
func Taobaoinventoryplaninvalid(clt *core.SDKClient, req *inventory.TaobaoinventoryplaninvalidAPIRequest, session string) (*inventory.TaobaoinventoryplaninvalidAPIResponse, error) {
	var resp inventory.TaobaoinventoryplaninvalidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

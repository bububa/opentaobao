package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanInvalid 失效计划库存
// taobao.inventory.plan.invalid
//
// 计划库存的失效服务
func TaobaoInventoryPlanInvalid(clt *core.SDKClient, req *inventory.TaobaoInventoryPlanInvalidAPIRequest, session string) (*inventory.TaobaoInventoryPlanInvalidAPIResponse, error) {
	var resp inventory.TaobaoInventoryPlanInvalidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

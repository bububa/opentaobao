package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

/* TaobaoInventoryPlanEdit
设置计划库存
taobao.inventory.plan.edit

初始化计划库存，或者编辑已经存在的计划库存 */
func TaobaoInventoryPlanEdit(clt *core.SDKClient, req *inventory.TaobaoInventoryPlanEditAPIRequest, session string) (*inventory.TaobaoInventoryPlanEditAPIResponse, error) {
	var resp inventory.TaobaoInventoryPlanEditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

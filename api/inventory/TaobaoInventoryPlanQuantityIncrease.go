package inventory

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/inventory"
)

/* 
计划库存的增量编辑 
taobao.inventory.plan.quantity.increase

计划库存的增量编辑
*/
func TaobaoInventoryPlanQuantityIncrease(clt *core.SDKClient, req *inventory.TaobaoInventoryPlanQuantityIncreaseAPIRequest, session string) (*inventory.TaobaoInventoryPlanQuantityIncreaseAPIResponse, error) {
    var resp inventory.TaobaoInventoryPlanQuantityIncreaseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

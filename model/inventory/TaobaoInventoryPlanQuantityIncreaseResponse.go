package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存的增量编辑 APIResponse
taobao.inventory.plan.quantity.increase

计划库存的增量编辑
*/
type TaobaoInventoryPlanQuantityIncreaseAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryPlanQuantityIncreaseResponse
}

type TaobaoInventoryPlanQuantityIncreaseResponse struct {
    XMLName xml.Name `xml:"inventory_plan_quantity_increase_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 批量返回结果
    
    Result   *BatchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

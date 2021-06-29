package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存查询 APIResponse
taobao.inventory.plan.query

计划库存查询
*/
type TaobaoInventoryPlanQueryAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryPlanQueryResponse
}

type TaobaoInventoryPlanQueryResponse struct {
    XMLName xml.Name `xml:"inventory_plan_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoInventoryPlanQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

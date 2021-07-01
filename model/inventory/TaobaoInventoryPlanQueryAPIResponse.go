package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划库存查询 API返回值 
taobao.inventory.plan.query

计划库存查询
*/
type TaobaoInventoryPlanQueryAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryPlanQueryAPIResponseModel
}

// 计划库存查询 成功返回结果
type TaobaoInventoryPlanQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"inventory_plan_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoInventoryPlanQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

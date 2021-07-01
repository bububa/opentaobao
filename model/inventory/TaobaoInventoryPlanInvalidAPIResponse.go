package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
失效计划库存 API返回值 
taobao.inventory.plan.invalid

计划库存的失效服务
*/
type TaobaoInventoryPlanInvalidAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryPlanInvalidAPIResponseModel
}

// 失效计划库存 成功返回结果
type TaobaoInventoryPlanInvalidAPIResponseModel struct {
    XMLName xml.Name `xml:"inventory_plan_invalid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 批量返回结果
    Result   *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

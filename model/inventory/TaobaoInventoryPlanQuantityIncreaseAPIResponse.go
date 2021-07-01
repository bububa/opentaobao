package inventory

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryPlanQuantityIncreaseAPIResponse
计划库存的增量编辑 API返回值
taobao.inventory.plan.quantity.increase

计划库存的增量编辑 */
type TaobaoInventoryPlanQuantityIncreaseAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryPlanQuantityIncreaseAPIResponseModel
}

// TaobaoInventoryPlanQuantityIncreaseAPIResponseModel is 计划库存的增量编辑 成功返回结果
type TaobaoInventoryPlanQuantityIncreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_quantity_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量返回结果
	Result *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

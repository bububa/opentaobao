package inventory

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanQuantityIncreaseAPIResponse 计划库存的增量编辑 API返回值
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
type TaobaoInventoryPlanQuantityIncreaseAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryPlanQuantityIncreaseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanQuantityIncreaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryPlanQuantityIncreaseAPIResponseModel).Reset()
}

// TaobaoInventoryPlanQuantityIncreaseAPIResponseModel is 计划库存的增量编辑 成功返回结果
type TaobaoInventoryPlanQuantityIncreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_quantity_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量返回结果
	Result *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanQuantityIncreaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoInventoryPlanQuantityIncreaseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryPlanQuantityIncreaseAPIResponse)
	},
}

// GetTaobaoInventoryPlanQuantityIncreaseAPIResponse 从 sync.Pool 获取 TaobaoInventoryPlanQuantityIncreaseAPIResponse
func GetTaobaoInventoryPlanQuantityIncreaseAPIResponse() *TaobaoInventoryPlanQuantityIncreaseAPIResponse {
	return poolTaobaoInventoryPlanQuantityIncreaseAPIResponse.Get().(*TaobaoInventoryPlanQuantityIncreaseAPIResponse)
}

// ReleaseTaobaoInventoryPlanQuantityIncreaseAPIResponse 将 TaobaoInventoryPlanQuantityIncreaseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryPlanQuantityIncreaseAPIResponse(v *TaobaoInventoryPlanQuantityIncreaseAPIResponse) {
	v.Reset()
	poolTaobaoInventoryPlanQuantityIncreaseAPIResponse.Put(v)
}

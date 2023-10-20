package inventory

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanInvalidAPIResponse 失效计划库存 API返回值
// taobao.inventory.plan.invalid
//
// 计划库存的失效服务
type TaobaoInventoryPlanInvalidAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryPlanInvalidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanInvalidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryPlanInvalidAPIResponseModel).Reset()
}

// TaobaoInventoryPlanInvalidAPIResponseModel is 失效计划库存 成功返回结果
type TaobaoInventoryPlanInvalidAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_invalid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量返回结果
	Result *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanInvalidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoInventoryPlanInvalidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryPlanInvalidAPIResponse)
	},
}

// GetTaobaoInventoryPlanInvalidAPIResponse 从 sync.Pool 获取 TaobaoInventoryPlanInvalidAPIResponse
func GetTaobaoInventoryPlanInvalidAPIResponse() *TaobaoInventoryPlanInvalidAPIResponse {
	return poolTaobaoInventoryPlanInvalidAPIResponse.Get().(*TaobaoInventoryPlanInvalidAPIResponse)
}

// ReleaseTaobaoInventoryPlanInvalidAPIResponse 将 TaobaoInventoryPlanInvalidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryPlanInvalidAPIResponse(v *TaobaoInventoryPlanInvalidAPIResponse) {
	v.Reset()
	poolTaobaoInventoryPlanInvalidAPIResponse.Put(v)
}

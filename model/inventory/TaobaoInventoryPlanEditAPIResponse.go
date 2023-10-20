package inventory

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanEditAPIResponse 设置计划库存 API返回值
// taobao.inventory.plan.edit
//
// 初始化计划库存，或者编辑已经存在的计划库存
type TaobaoInventoryPlanEditAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryPlanEditAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryPlanEditAPIResponseModel).Reset()
}

// TaobaoInventoryPlanEditAPIResponseModel is 设置计划库存 成功返回结果
type TaobaoInventoryPlanEditAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批量返回结果
	Result *BatchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoInventoryPlanEditAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryPlanEditAPIResponse)
	},
}

// GetTaobaoInventoryPlanEditAPIResponse 从 sync.Pool 获取 TaobaoInventoryPlanEditAPIResponse
func GetTaobaoInventoryPlanEditAPIResponse() *TaobaoInventoryPlanEditAPIResponse {
	return poolTaobaoInventoryPlanEditAPIResponse.Get().(*TaobaoInventoryPlanEditAPIResponse)
}

// ReleaseTaobaoInventoryPlanEditAPIResponse 将 TaobaoInventoryPlanEditAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryPlanEditAPIResponse(v *TaobaoInventoryPlanEditAPIResponse) {
	v.Reset()
	poolTaobaoInventoryPlanEditAPIResponse.Put(v)
}

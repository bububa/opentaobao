package inventory

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanQueryAPIResponse 计划库存查询 API返回值
// taobao.inventory.plan.query
//
// 计划库存查询
type TaobaoInventoryPlanQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryPlanQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryPlanQueryAPIResponseModel).Reset()
}

// TaobaoInventoryPlanQueryAPIResponseModel is 计划库存查询 成功返回结果
type TaobaoInventoryPlanQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_plan_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoInventoryPlanQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryPlanQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoInventoryPlanQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryPlanQueryAPIResponse)
	},
}

// GetTaobaoInventoryPlanQueryAPIResponse 从 sync.Pool 获取 TaobaoInventoryPlanQueryAPIResponse
func GetTaobaoInventoryPlanQueryAPIResponse() *TaobaoInventoryPlanQueryAPIResponse {
	return poolTaobaoInventoryPlanQueryAPIResponse.Get().(*TaobaoInventoryPlanQueryAPIResponse)
}

// ReleaseTaobaoInventoryPlanQueryAPIResponse 将 TaobaoInventoryPlanQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryPlanQueryAPIResponse(v *TaobaoInventoryPlanQueryAPIResponse) {
	v.Reset()
	poolTaobaoInventoryPlanQueryAPIResponse.Put(v)
}

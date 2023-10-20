package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecyclePredeductOldGetAPIResponse 查询回收单前置抵扣详情 API返回值
// taobao.recycle.prededuct.old.get
//
// 查询回收单前置抵扣详情
type TaobaoRecyclePredeductOldGetAPIResponse struct {
	model.CommonResponse
	TaobaoRecyclePredeductOldGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRecyclePredeductOldGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRecyclePredeductOldGetAPIResponseModel).Reset()
}

// TaobaoRecyclePredeductOldGetAPIResponseModel is 查询回收单前置抵扣详情 成功返回结果
type TaobaoRecyclePredeductOldGetAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_prededuct_old_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回收单前置抵扣详情
	Data *OfnRecycleOrderPreDeductDetailDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRecyclePredeductOldGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoRecyclePredeductOldGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRecyclePredeductOldGetAPIResponse)
	},
}

// GetTaobaoRecyclePredeductOldGetAPIResponse 从 sync.Pool 获取 TaobaoRecyclePredeductOldGetAPIResponse
func GetTaobaoRecyclePredeductOldGetAPIResponse() *TaobaoRecyclePredeductOldGetAPIResponse {
	return poolTaobaoRecyclePredeductOldGetAPIResponse.Get().(*TaobaoRecyclePredeductOldGetAPIResponse)
}

// ReleaseTaobaoRecyclePredeductOldGetAPIResponse 将 TaobaoRecyclePredeductOldGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRecyclePredeductOldGetAPIResponse(v *TaobaoRecyclePredeductOldGetAPIResponse) {
	v.Reset()
	poolTaobaoRecyclePredeductOldGetAPIResponse.Put(v)
}

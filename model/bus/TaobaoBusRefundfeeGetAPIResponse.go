package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundfeeGetAPIResponse 查询退票费用明细 API返回值
// taobao.bus.refundfee.get
//
// 查询退票的费用信息
type TaobaoBusRefundfeeGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusRefundfeeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusRefundfeeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusRefundfeeGetAPIResponseModel).Reset()
}

// TaobaoBusRefundfeeGetAPIResponseModel is 查询退票费用明细 成功返回结果
type TaobaoBusRefundfeeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_refundfee_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *B2BQueryRefundFeeRp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusRefundfeeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusRefundfeeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusRefundfeeGetAPIResponse)
	},
}

// GetTaobaoBusRefundfeeGetAPIResponse 从 sync.Pool 获取 TaobaoBusRefundfeeGetAPIResponse
func GetTaobaoBusRefundfeeGetAPIResponse() *TaobaoBusRefundfeeGetAPIResponse {
	return poolTaobaoBusRefundfeeGetAPIResponse.Get().(*TaobaoBusRefundfeeGetAPIResponse)
}

// ReleaseTaobaoBusRefundfeeGetAPIResponse 将 TaobaoBusRefundfeeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusRefundfeeGetAPIResponse(v *TaobaoBusRefundfeeGetAPIResponse) {
	v.Reset()
	poolTaobaoBusRefundfeeGetAPIResponse.Put(v)
}

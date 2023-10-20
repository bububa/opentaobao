package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseTailpaymentbackAPIResponse 尾款处置方案回传 API返回值
// tmall.car.lease.tailpaymentback
//
// 尾款处置方案回传
type TmallCarLeaseTailpaymentbackAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseTailpaymentbackAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseTailpaymentbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseTailpaymentbackAPIResponseModel).Reset()
}

// TmallCarLeaseTailpaymentbackAPIResponseModel is 尾款处置方案回传 成功返回结果
type TmallCarLeaseTailpaymentbackAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_tailpaymentback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallCarLeaseTailpaymentbackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseTailpaymentbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseTailpaymentbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseTailpaymentbackAPIResponse)
	},
}

// GetTmallCarLeaseTailpaymentbackAPIResponse 从 sync.Pool 获取 TmallCarLeaseTailpaymentbackAPIResponse
func GetTmallCarLeaseTailpaymentbackAPIResponse() *TmallCarLeaseTailpaymentbackAPIResponse {
	return poolTmallCarLeaseTailpaymentbackAPIResponse.Get().(*TmallCarLeaseTailpaymentbackAPIResponse)
}

// ReleaseTmallCarLeaseTailpaymentbackAPIResponse 将 TmallCarLeaseTailpaymentbackAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseTailpaymentbackAPIResponse(v *TmallCarLeaseTailpaymentbackAPIResponse) {
	v.Reset()
	poolTmallCarLeaseTailpaymentbackAPIResponse.Put(v)
}

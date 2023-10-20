package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseRiskcallbackAPIResponse 整车租赁风控模型回调 API返回值
// tmall.car.lease.riskcallback
//
// 租赁公司回调风控结果
type TmallCarLeaseRiskcallbackAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseRiskcallbackAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseRiskcallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseRiskcallbackAPIResponseModel).Reset()
}

// TmallCarLeaseRiskcallbackAPIResponseModel is 整车租赁风控模型回调 成功返回结果
type TmallCarLeaseRiskcallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_riskcallback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集合
	Result *TmallCarLeaseRiskcallbackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseRiskcallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseRiskcallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseRiskcallbackAPIResponse)
	},
}

// GetTmallCarLeaseRiskcallbackAPIResponse 从 sync.Pool 获取 TmallCarLeaseRiskcallbackAPIResponse
func GetTmallCarLeaseRiskcallbackAPIResponse() *TmallCarLeaseRiskcallbackAPIResponse {
	return poolTmallCarLeaseRiskcallbackAPIResponse.Get().(*TmallCarLeaseRiskcallbackAPIResponse)
}

// ReleaseTmallCarLeaseRiskcallbackAPIResponse 将 TmallCarLeaseRiskcallbackAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseRiskcallbackAPIResponse(v *TmallCarLeaseRiskcallbackAPIResponse) {
	v.Reset()
	poolTmallCarLeaseRiskcallbackAPIResponse.Put(v)
}

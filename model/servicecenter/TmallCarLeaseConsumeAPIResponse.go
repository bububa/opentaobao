package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseConsumeAPIResponse 汽车租赁核销 API返回值
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
type TmallCarLeaseConsumeAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseConsumeAPIResponseModel).Reset()
}

// TmallCarLeaseConsumeAPIResponseModel is 汽车租赁核销 成功返回结果
type TmallCarLeaseConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集合
	Result *TmallCarLeaseConsumeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseConsumeAPIResponse)
	},
}

// GetTmallCarLeaseConsumeAPIResponse 从 sync.Pool 获取 TmallCarLeaseConsumeAPIResponse
func GetTmallCarLeaseConsumeAPIResponse() *TmallCarLeaseConsumeAPIResponse {
	return poolTmallCarLeaseConsumeAPIResponse.Get().(*TmallCarLeaseConsumeAPIResponse)
}

// ReleaseTmallCarLeaseConsumeAPIResponse 将 TmallCarLeaseConsumeAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseConsumeAPIResponse(v *TmallCarLeaseConsumeAPIResponse) {
	v.Reset()
	poolTmallCarLeaseConsumeAPIResponse.Put(v)
}

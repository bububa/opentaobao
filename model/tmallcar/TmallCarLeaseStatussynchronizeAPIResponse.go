package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseStatussynchronizeAPIResponse 天猫开新车租后状态同步 API返回值
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
type TmallCarLeaseStatussynchronizeAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseStatussynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseStatussynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseStatussynchronizeAPIResponseModel).Reset()
}

// TmallCarLeaseStatussynchronizeAPIResponseModel is 天猫开新车租后状态同步 成功返回结果
type TmallCarLeaseStatussynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_statussynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseStatussynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseStatussynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseStatussynchronizeAPIResponse)
	},
}

// GetTmallCarLeaseStatussynchronizeAPIResponse 从 sync.Pool 获取 TmallCarLeaseStatussynchronizeAPIResponse
func GetTmallCarLeaseStatussynchronizeAPIResponse() *TmallCarLeaseStatussynchronizeAPIResponse {
	return poolTmallCarLeaseStatussynchronizeAPIResponse.Get().(*TmallCarLeaseStatussynchronizeAPIResponse)
}

// ReleaseTmallCarLeaseStatussynchronizeAPIResponse 将 TmallCarLeaseStatussynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseStatussynchronizeAPIResponse(v *TmallCarLeaseStatussynchronizeAPIResponse) {
	v.Reset()
	poolTmallCarLeaseStatussynchronizeAPIResponse.Put(v)
}

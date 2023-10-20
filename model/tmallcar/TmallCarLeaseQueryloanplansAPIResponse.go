package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseQueryloanplansAPIResponse 天猫开新车租后查询还款计划 API返回值
// tmall.car.lease.queryloanplans
//
// 天猫开新车租后查询还款计划
type TmallCarLeaseQueryloanplansAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseQueryloanplansAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseQueryloanplansAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseQueryloanplansAPIResponseModel).Reset()
}

// TmallCarLeaseQueryloanplansAPIResponseModel is 天猫开新车租后查询还款计划 成功返回结果
type TmallCarLeaseQueryloanplansAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_queryloanplans_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseQueryloanplansAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseQueryloanplansAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseQueryloanplansAPIResponse)
	},
}

// GetTmallCarLeaseQueryloanplansAPIResponse 从 sync.Pool 获取 TmallCarLeaseQueryloanplansAPIResponse
func GetTmallCarLeaseQueryloanplansAPIResponse() *TmallCarLeaseQueryloanplansAPIResponse {
	return poolTmallCarLeaseQueryloanplansAPIResponse.Get().(*TmallCarLeaseQueryloanplansAPIResponse)
}

// ReleaseTmallCarLeaseQueryloanplansAPIResponse 将 TmallCarLeaseQueryloanplansAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseQueryloanplansAPIResponse(v *TmallCarLeaseQueryloanplansAPIResponse) {
	v.Reset()
	poolTmallCarLeaseQueryloanplansAPIResponse.Put(v)
}

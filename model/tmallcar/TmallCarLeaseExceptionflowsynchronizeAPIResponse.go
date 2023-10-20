package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseExceptionflowsynchronizeAPIResponse 天猫开新车租后异常流线下处理状态通知接口 API返回值
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
type TmallCarLeaseExceptionflowsynchronizeAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseExceptionflowsynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseExceptionflowsynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseExceptionflowsynchronizeAPIResponseModel).Reset()
}

// TmallCarLeaseExceptionflowsynchronizeAPIResponseModel is 天猫开新车租后异常流线下处理状态通知接口 成功返回结果
type TmallCarLeaseExceptionflowsynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_exceptionflowsynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseExceptionflowsynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseExceptionflowsynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseExceptionflowsynchronizeAPIResponse)
	},
}

// GetTmallCarLeaseExceptionflowsynchronizeAPIResponse 从 sync.Pool 获取 TmallCarLeaseExceptionflowsynchronizeAPIResponse
func GetTmallCarLeaseExceptionflowsynchronizeAPIResponse() *TmallCarLeaseExceptionflowsynchronizeAPIResponse {
	return poolTmallCarLeaseExceptionflowsynchronizeAPIResponse.Get().(*TmallCarLeaseExceptionflowsynchronizeAPIResponse)
}

// ReleaseTmallCarLeaseExceptionflowsynchronizeAPIResponse 将 TmallCarLeaseExceptionflowsynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseExceptionflowsynchronizeAPIResponse(v *TmallCarLeaseExceptionflowsynchronizeAPIResponse) {
	v.Reset()
	poolTmallCarLeaseExceptionflowsynchronizeAPIResponse.Put(v)
}

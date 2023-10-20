package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopInterceptAPIResponse 阿里巴巴.天猫. 履约订单. 配送拦截 API返回值
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
type AlibabaTianmaoUopInterceptAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoUopInterceptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianmaoUopInterceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianmaoUopInterceptAPIResponseModel).Reset()
}

// AlibabaTianmaoUopInterceptAPIResponseModel is 阿里巴巴.天猫. 履约订单. 配送拦截 成功返回结果
type AlibabaTianmaoUopInterceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_uop_intercept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianmaoUopInterceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTianmaoUopInterceptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianmaoUopInterceptAPIResponse)
	},
}

// GetAlibabaTianmaoUopInterceptAPIResponse 从 sync.Pool 获取 AlibabaTianmaoUopInterceptAPIResponse
func GetAlibabaTianmaoUopInterceptAPIResponse() *AlibabaTianmaoUopInterceptAPIResponse {
	return poolAlibabaTianmaoUopInterceptAPIResponse.Get().(*AlibabaTianmaoUopInterceptAPIResponse)
}

// ReleaseAlibabaTianmaoUopInterceptAPIResponse 将 AlibabaTianmaoUopInterceptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianmaoUopInterceptAPIResponse(v *AlibabaTianmaoUopInterceptAPIResponse) {
	v.Reset()
	poolAlibabaTianmaoUopInterceptAPIResponse.Put(v)
}

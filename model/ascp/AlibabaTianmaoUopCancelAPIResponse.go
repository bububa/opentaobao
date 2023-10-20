package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopCancelAPIResponse 阿里巴巴.天猫. 履约订单. 取消 API返回值
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
type AlibabaTianmaoUopCancelAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoUopCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianmaoUopCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianmaoUopCancelAPIResponseModel).Reset()
}

// AlibabaTianmaoUopCancelAPIResponseModel is 阿里巴巴.天猫. 履约订单. 取消 成功返回结果
type AlibabaTianmaoUopCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_uop_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianmaoUopCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTianmaoUopCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianmaoUopCancelAPIResponse)
	},
}

// GetAlibabaTianmaoUopCancelAPIResponse 从 sync.Pool 获取 AlibabaTianmaoUopCancelAPIResponse
func GetAlibabaTianmaoUopCancelAPIResponse() *AlibabaTianmaoUopCancelAPIResponse {
	return poolAlibabaTianmaoUopCancelAPIResponse.Get().(*AlibabaTianmaoUopCancelAPIResponse)
}

// ReleaseAlibabaTianmaoUopCancelAPIResponse 将 AlibabaTianmaoUopCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianmaoUopCancelAPIResponse(v *AlibabaTianmaoUopCancelAPIResponse) {
	v.Reset()
	poolAlibabaTianmaoUopCancelAPIResponse.Put(v)
}

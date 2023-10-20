package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopConsignAPIResponse 阿里巴巴.天猫. 履约订单. 发货 API返回值
// alibaba.tianmao.uop.consign
//
// 阿里巴巴.天猫. 履约订单. 发货
type AlibabaTianmaoUopConsignAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoUopConsignAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianmaoUopConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianmaoUopConsignAPIResponseModel).Reset()
}

// AlibabaTianmaoUopConsignAPIResponseModel is 阿里巴巴.天猫. 履约订单. 发货 成功返回结果
type AlibabaTianmaoUopConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_uop_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianmaoUopConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTianmaoUopConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianmaoUopConsignAPIResponse)
	},
}

// GetAlibabaTianmaoUopConsignAPIResponse 从 sync.Pool 获取 AlibabaTianmaoUopConsignAPIResponse
func GetAlibabaTianmaoUopConsignAPIResponse() *AlibabaTianmaoUopConsignAPIResponse {
	return poolAlibabaTianmaoUopConsignAPIResponse.Get().(*AlibabaTianmaoUopConsignAPIResponse)
}

// ReleaseAlibabaTianmaoUopConsignAPIResponse 将 AlibabaTianmaoUopConsignAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianmaoUopConsignAPIResponse(v *AlibabaTianmaoUopConsignAPIResponse) {
	v.Reset()
	poolAlibabaTianmaoUopConsignAPIResponse.Put(v)
}

package yunos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunCosmoGatewayInvokeAPIResponse alios cosmo服务调用 API返回值
// aliyun.cosmo.gateway.invoke
//
// AliOS cosmo服务分发平台对外调用接口
type AliyunCosmoGatewayInvokeAPIResponse struct {
	model.CommonResponse
	AliyunCosmoGatewayInvokeAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunCosmoGatewayInvokeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunCosmoGatewayInvokeAPIResponseModel).Reset()
}

// AliyunCosmoGatewayInvokeAPIResponseModel is alios cosmo服务调用 成功返回结果
type AliyunCosmoGatewayInvokeAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_cosmo_gateway_invoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RdamResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliyunCosmoGatewayInvokeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliyunCosmoGatewayInvokeAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunCosmoGatewayInvokeAPIResponse)
	},
}

// GetAliyunCosmoGatewayInvokeAPIResponse 从 sync.Pool 获取 AliyunCosmoGatewayInvokeAPIResponse
func GetAliyunCosmoGatewayInvokeAPIResponse() *AliyunCosmoGatewayInvokeAPIResponse {
	return poolAliyunCosmoGatewayInvokeAPIResponse.Get().(*AliyunCosmoGatewayInvokeAPIResponse)
}

// ReleaseAliyunCosmoGatewayInvokeAPIResponse 将 AliyunCosmoGatewayInvokeAPIResponse 保存到 sync.Pool
func ReleaseAliyunCosmoGatewayInvokeAPIResponse(v *AliyunCosmoGatewayInvokeAPIResponse) {
	v.Reset()
	poolAliyunCosmoGatewayInvokeAPIResponse.Put(v)
}

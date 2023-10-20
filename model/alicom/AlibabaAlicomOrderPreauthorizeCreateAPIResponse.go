package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderPreauthorizeCreateAPIResponse 业务办理结果 API返回值
// alibaba.alicom.order.preauthorize.create
//
// 授授权:签约结果通知
type AlibabaAlicomOrderPreauthorizeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomOrderPreauthorizeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel).Reset()
}

// AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel is 业务办理结果 成功返回结果
type AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomOrderPreauthorizeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomOrderPreauthorizeCreateAPIResponse)
	},
}

// GetAlibabaAlicomOrderPreauthorizeCreateAPIResponse 从 sync.Pool 获取 AlibabaAlicomOrderPreauthorizeCreateAPIResponse
func GetAlibabaAlicomOrderPreauthorizeCreateAPIResponse() *AlibabaAlicomOrderPreauthorizeCreateAPIResponse {
	return poolAlibabaAlicomOrderPreauthorizeCreateAPIResponse.Get().(*AlibabaAlicomOrderPreauthorizeCreateAPIResponse)
}

// ReleaseAlibabaAlicomOrderPreauthorizeCreateAPIResponse 将 AlibabaAlicomOrderPreauthorizeCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomOrderPreauthorizeCreateAPIResponse(v *AlibabaAlicomOrderPreauthorizeCreateAPIResponse) {
	v.Reset()
	poolAlibabaAlicomOrderPreauthorizeCreateAPIResponse.Put(v)
}

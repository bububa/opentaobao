package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractOrderCheckuserimeiAPIResponse 金融购机验证设备号 API返回值
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
type AlibabaInteractOrderCheckuserimeiAPIResponse struct {
	model.CommonResponse
	AlibabaInteractOrderCheckuserimeiAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractOrderCheckuserimeiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractOrderCheckuserimeiAPIResponseModel).Reset()
}

// AlibabaInteractOrderCheckuserimeiAPIResponseModel is 金融购机验证设备号 成功返回结果
type AlibabaInteractOrderCheckuserimeiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_order_checkuserimei_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	SuccessStatus bool `json:"success_status,omitempty" xml:"success_status,omitempty"`
	// 响应数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractOrderCheckuserimeiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Message = ""
	m.SuccessStatus = false
	m.Data = false
}

var poolAlibabaInteractOrderCheckuserimeiAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractOrderCheckuserimeiAPIResponse)
	},
}

// GetAlibabaInteractOrderCheckuserimeiAPIResponse 从 sync.Pool 获取 AlibabaInteractOrderCheckuserimeiAPIResponse
func GetAlibabaInteractOrderCheckuserimeiAPIResponse() *AlibabaInteractOrderCheckuserimeiAPIResponse {
	return poolAlibabaInteractOrderCheckuserimeiAPIResponse.Get().(*AlibabaInteractOrderCheckuserimeiAPIResponse)
}

// ReleaseAlibabaInteractOrderCheckuserimeiAPIResponse 将 AlibabaInteractOrderCheckuserimeiAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractOrderCheckuserimeiAPIResponse(v *AlibabaInteractOrderCheckuserimeiAPIResponse) {
	v.Reset()
	poolAlibabaInteractOrderCheckuserimeiAPIResponse.Put(v)
}

package globalvirtual

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGlobalVirtualSendcodeAPIResponse 国际虚拟商品发码服务 API返回值
// alibaba.global.virtual.sendcode
//
// global virtual send code service
type AlibabaGlobalVirtualSendcodeAPIResponse struct {
	model.CommonResponse
	AlibabaGlobalVirtualSendcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaGlobalVirtualSendcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaGlobalVirtualSendcodeAPIResponseModel).Reset()
}

// AlibabaGlobalVirtualSendcodeAPIResponseModel is 国际虚拟商品发码服务 成功返回结果
type AlibabaGlobalVirtualSendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_global_virtual_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result describe
	Result *AlibabaGlobalVirtualSendcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaGlobalVirtualSendcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaGlobalVirtualSendcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaGlobalVirtualSendcodeAPIResponse)
	},
}

// GetAlibabaGlobalVirtualSendcodeAPIResponse 从 sync.Pool 获取 AlibabaGlobalVirtualSendcodeAPIResponse
func GetAlibabaGlobalVirtualSendcodeAPIResponse() *AlibabaGlobalVirtualSendcodeAPIResponse {
	return poolAlibabaGlobalVirtualSendcodeAPIResponse.Get().(*AlibabaGlobalVirtualSendcodeAPIResponse)
}

// ReleaseAlibabaGlobalVirtualSendcodeAPIResponse 将 AlibabaGlobalVirtualSendcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaGlobalVirtualSendcodeAPIResponse(v *AlibabaGlobalVirtualSendcodeAPIResponse) {
	v.Reset()
	poolAlibabaGlobalVirtualSendcodeAPIResponse.Put(v)
}

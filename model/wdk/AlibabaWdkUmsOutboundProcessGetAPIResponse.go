package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsOutboundProcessGetAPIResponse 出库业务UMS异步处理结果返回 API返回值
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
type AlibabaWdkUmsOutboundProcessGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsOutboundProcessGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkUmsOutboundProcessGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkUmsOutboundProcessGetAPIResponseModel).Reset()
}

// AlibabaWdkUmsOutboundProcessGetAPIResponseModel is 出库业务UMS异步处理结果返回 成功返回结果
type AlibabaWdkUmsOutboundProcessGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_outbound_process_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkUmsOutboundProcessGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkUmsOutboundProcessGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsOutboundProcessGetAPIResponse)
	},
}

// GetAlibabaWdkUmsOutboundProcessGetAPIResponse 从 sync.Pool 获取 AlibabaWdkUmsOutboundProcessGetAPIResponse
func GetAlibabaWdkUmsOutboundProcessGetAPIResponse() *AlibabaWdkUmsOutboundProcessGetAPIResponse {
	return poolAlibabaWdkUmsOutboundProcessGetAPIResponse.Get().(*AlibabaWdkUmsOutboundProcessGetAPIResponse)
}

// ReleaseAlibabaWdkUmsOutboundProcessGetAPIResponse 将 AlibabaWdkUmsOutboundProcessGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkUmsOutboundProcessGetAPIResponse(v *AlibabaWdkUmsOutboundProcessGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkUmsOutboundProcessGetAPIResponse.Put(v)
}

package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsInboundAPIResponse 入库-ERP下发单 API返回值
// alibaba.wdk.ums.inbound
//
// 入库-ERP下发单
type AlibabaWdkUmsInboundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsInboundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkUmsInboundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkUmsInboundAPIResponseModel).Reset()
}

// AlibabaWdkUmsInboundAPIResponseModel is 入库-ERP下发单 成功返回结果
type AlibabaWdkUmsInboundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_inbound_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkUmsInboundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkUmsInboundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsInboundAPIResponse)
	},
}

// GetAlibabaWdkUmsInboundAPIResponse 从 sync.Pool 获取 AlibabaWdkUmsInboundAPIResponse
func GetAlibabaWdkUmsInboundAPIResponse() *AlibabaWdkUmsInboundAPIResponse {
	return poolAlibabaWdkUmsInboundAPIResponse.Get().(*AlibabaWdkUmsInboundAPIResponse)
}

// ReleaseAlibabaWdkUmsInboundAPIResponse 将 AlibabaWdkUmsInboundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkUmsInboundAPIResponse(v *AlibabaWdkUmsInboundAPIResponse) {
	v.Reset()
	poolAlibabaWdkUmsInboundAPIResponse.Put(v)
}

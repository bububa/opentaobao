package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsOutboundAPIResponse 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) API返回值
// alibaba.wdk.ums.outbound
//
// 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
type AlibabaWdkUmsOutboundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsOutboundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkUmsOutboundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkUmsOutboundAPIResponseModel).Reset()
}

// AlibabaWdkUmsOutboundAPIResponseModel is 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) 成功返回结果
type AlibabaWdkUmsOutboundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_outbound_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkUmsOutboundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkUmsOutboundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsOutboundAPIResponse)
	},
}

// GetAlibabaWdkUmsOutboundAPIResponse 从 sync.Pool 获取 AlibabaWdkUmsOutboundAPIResponse
func GetAlibabaWdkUmsOutboundAPIResponse() *AlibabaWdkUmsOutboundAPIResponse {
	return poolAlibabaWdkUmsOutboundAPIResponse.Get().(*AlibabaWdkUmsOutboundAPIResponse)
}

// ReleaseAlibabaWdkUmsOutboundAPIResponse 将 AlibabaWdkUmsOutboundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkUmsOutboundAPIResponse(v *AlibabaWdkUmsOutboundAPIResponse) {
	v.Reset()
	poolAlibabaWdkUmsOutboundAPIResponse.Put(v)
}

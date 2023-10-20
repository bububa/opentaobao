package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsShiftGetAPIResponse 移库单获取 API返回值
// alibaba.wdk.ums.shift.get
//
// 移库单获取
type AlibabaWdkUmsShiftGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsShiftGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkUmsShiftGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkUmsShiftGetAPIResponseModel).Reset()
}

// AlibabaWdkUmsShiftGetAPIResponseModel is 移库单获取 成功返回结果
type AlibabaWdkUmsShiftGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_shift_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkUmsShiftGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkUmsShiftGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkUmsShiftGetAPIResponse)
	},
}

// GetAlibabaWdkUmsShiftGetAPIResponse 从 sync.Pool 获取 AlibabaWdkUmsShiftGetAPIResponse
func GetAlibabaWdkUmsShiftGetAPIResponse() *AlibabaWdkUmsShiftGetAPIResponse {
	return poolAlibabaWdkUmsShiftGetAPIResponse.Get().(*AlibabaWdkUmsShiftGetAPIResponse)
}

// ReleaseAlibabaWdkUmsShiftGetAPIResponse 将 AlibabaWdkUmsShiftGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkUmsShiftGetAPIResponse(v *AlibabaWdkUmsShiftGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkUmsShiftGetAPIResponse.Put(v)
}

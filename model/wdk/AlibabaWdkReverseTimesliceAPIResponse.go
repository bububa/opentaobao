package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseTimesliceAPIResponse 逆向取货时间片查询 API返回值
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
type AlibabaWdkReverseTimesliceAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseTimesliceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkReverseTimesliceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkReverseTimesliceAPIResponseModel).Reset()
}

// AlibabaWdkReverseTimesliceAPIResponseModel is 逆向取货时间片查询 成功返回结果
type AlibabaWdkReverseTimesliceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_timeslice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkReverseTimesliceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkReverseTimesliceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseTimesliceAPIResponse)
	},
}

// GetAlibabaWdkReverseTimesliceAPIResponse 从 sync.Pool 获取 AlibabaWdkReverseTimesliceAPIResponse
func GetAlibabaWdkReverseTimesliceAPIResponse() *AlibabaWdkReverseTimesliceAPIResponse {
	return poolAlibabaWdkReverseTimesliceAPIResponse.Get().(*AlibabaWdkReverseTimesliceAPIResponse)
}

// ReleaseAlibabaWdkReverseTimesliceAPIResponse 将 AlibabaWdkReverseTimesliceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkReverseTimesliceAPIResponse(v *AlibabaWdkReverseTimesliceAPIResponse) {
	v.Reset()
	poolAlibabaWdkReverseTimesliceAPIResponse.Put(v)
}

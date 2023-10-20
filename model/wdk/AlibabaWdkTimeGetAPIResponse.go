package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTimeGetAPIResponse 获得当前系统时间 API返回值
// alibaba.wdk.time.get
//
// 获得当前系统时间
type AlibabaWdkTimeGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTimeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTimeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTimeGetAPIResponseModel).Reset()
}

// AlibabaWdkTimeGetAPIResponseModel is 获得当前系统时间 成功返回结果
type AlibabaWdkTimeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_time_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// date
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// dateTime
	DateTime int64 `json:"date_time,omitempty" xml:"date_time,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTimeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Date = ""
	m.DateTime = 0
}

var poolAlibabaWdkTimeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTimeGetAPIResponse)
	},
}

// GetAlibabaWdkTimeGetAPIResponse 从 sync.Pool 获取 AlibabaWdkTimeGetAPIResponse
func GetAlibabaWdkTimeGetAPIResponse() *AlibabaWdkTimeGetAPIResponse {
	return poolAlibabaWdkTimeGetAPIResponse.Get().(*AlibabaWdkTimeGetAPIResponse)
}

// ReleaseAlibabaWdkTimeGetAPIResponse 将 AlibabaWdkTimeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTimeGetAPIResponse(v *AlibabaWdkTimeGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkTimeGetAPIResponse.Put(v)
}

package wdk

import (
	"encoding/xml"

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

// AlibabaWdkTimeGetAPIResponseModel is 获得当前系统时间 成功返回结果
type AlibabaWdkTimeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_time_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// dateTime
	DateTime int64 `json:"date_time,omitempty" xml:"date_time,omitempty"`
	// date
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}

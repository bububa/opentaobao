package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponse 是否支持云回看新SDK API返回值
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
type AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponse struct {
	model.CommonResponse
	AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponseModel
}

// AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponseModel is 是否支持云回看新SDK 成功返回结果
type AlibabaailabstmallgeniesdkdeviceissupportsdkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_sdk_device_issupportsdk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaailabstmallgeniesdkdeviceissupportsdkResult `json:"result,omitempty" xml:"result,omitempty"`
}

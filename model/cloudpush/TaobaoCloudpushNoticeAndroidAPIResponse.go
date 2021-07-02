package cloudpush

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudpushNoticeAndroidAPIResponse 百川云推送发送通知给android API返回值
// taobao.cloudpush.notice.android
//
// 百川云推送发送通知给android
type TaobaoCloudpushNoticeAndroidAPIResponse struct {
	model.CommonResponse
	TaobaoCloudpushNoticeAndroidAPIResponseModel
}

// TaobaoCloudpushNoticeAndroidAPIResponseModel is 百川云推送发送通知给android 成功返回结果
type TaobaoCloudpushNoticeAndroidAPIResponseModel struct {
	XMLName xml.Name `xml:"cloudpush_notice_android_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 请求的错误代码.
	RequestErrorCode int64 `json:"request_error_code,omitempty" xml:"request_error_code,omitempty"`
	// 请求的错误信息.
	RequestErrorMsg string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`
}

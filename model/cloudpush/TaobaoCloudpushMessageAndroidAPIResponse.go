package cloudpush

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCloudpushMessageAndroidAPIResponse
百川云推送发送消息给android API返回值
taobao.cloudpush.message.android

百川用户使用云推送发送消息给android */
type TaobaoCloudpushMessageAndroidAPIResponse struct {
	model.CommonResponse
	TaobaoCloudpushMessageAndroidAPIResponseModel
}

// TaobaoCloudpushMessageAndroidAPIResponseModel is 百川云推送发送消息给android 成功返回结果
type TaobaoCloudpushMessageAndroidAPIResponseModel struct {
	XMLName xml.Name `xml:"cloudpush_message_android_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求是否成功.
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 若请求失败，则返回对应的error code.
	RequestErrorCode int64 `json:"request_error_code,omitempty" xml:"request_error_code,omitempty"`
	// 请求失败后返回的错误信息.
	RequestErrorMsg string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`
}

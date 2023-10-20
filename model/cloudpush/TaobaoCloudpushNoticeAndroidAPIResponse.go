package cloudpush

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoCloudpushNoticeAndroidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCloudpushNoticeAndroidAPIResponseModel).Reset()
}

// TaobaoCloudpushNoticeAndroidAPIResponseModel is 百川云推送发送通知给android 成功返回结果
type TaobaoCloudpushNoticeAndroidAPIResponseModel struct {
	XMLName xml.Name `xml:"cloudpush_notice_android_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求的错误信息.
	RequestErrorMsg string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`
	// 请求的错误代码.
	RequestErrorCode int64 `json:"request_error_code,omitempty" xml:"request_error_code,omitempty"`
	// 请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCloudpushNoticeAndroidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RequestErrorMsg = ""
	m.RequestErrorCode = 0
	m.IsSuccess = false
}

var poolTaobaoCloudpushNoticeAndroidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCloudpushNoticeAndroidAPIResponse)
	},
}

// GetTaobaoCloudpushNoticeAndroidAPIResponse 从 sync.Pool 获取 TaobaoCloudpushNoticeAndroidAPIResponse
func GetTaobaoCloudpushNoticeAndroidAPIResponse() *TaobaoCloudpushNoticeAndroidAPIResponse {
	return poolTaobaoCloudpushNoticeAndroidAPIResponse.Get().(*TaobaoCloudpushNoticeAndroidAPIResponse)
}

// ReleaseTaobaoCloudpushNoticeAndroidAPIResponse 将 TaobaoCloudpushNoticeAndroidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCloudpushNoticeAndroidAPIResponse(v *TaobaoCloudpushNoticeAndroidAPIResponse) {
	v.Reset()
	poolTaobaoCloudpushNoticeAndroidAPIResponse.Put(v)
}

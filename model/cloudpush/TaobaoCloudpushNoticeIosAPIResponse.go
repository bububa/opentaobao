package cloudpush

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudpushNoticeIosAPIResponse 推送通知给ios设备 API返回值
// taobao.cloudpush.notice.ios
//
// 推送通知给ios设备
type TaobaoCloudpushNoticeIosAPIResponse struct {
	model.CommonResponse
	TaobaoCloudpushNoticeIosAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCloudpushNoticeIosAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCloudpushNoticeIosAPIResponseModel).Reset()
}

// TaobaoCloudpushNoticeIosAPIResponseModel is 推送通知给ios设备 成功返回结果
type TaobaoCloudpushNoticeIosAPIResponseModel struct {
	XMLName xml.Name `xml:"cloudpush_notice_ios_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求产生的错误信息.
	RequestErrorMsg string `json:"request_error_msg,omitempty" xml:"request_error_msg,omitempty"`
	// 请求错误时产生的错误代码.
	RequestErrorCode int64 `json:"request_error_code,omitempty" xml:"request_error_code,omitempty"`
	// 请求是否成功.
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCloudpushNoticeIosAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RequestErrorMsg = ""
	m.RequestErrorCode = 0
	m.IsSuccess = false
}

var poolTaobaoCloudpushNoticeIosAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCloudpushNoticeIosAPIResponse)
	},
}

// GetTaobaoCloudpushNoticeIosAPIResponse 从 sync.Pool 获取 TaobaoCloudpushNoticeIosAPIResponse
func GetTaobaoCloudpushNoticeIosAPIResponse() *TaobaoCloudpushNoticeIosAPIResponse {
	return poolTaobaoCloudpushNoticeIosAPIResponse.Get().(*TaobaoCloudpushNoticeIosAPIResponse)
}

// ReleaseTaobaoCloudpushNoticeIosAPIResponse 将 TaobaoCloudpushNoticeIosAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCloudpushNoticeIosAPIResponse(v *TaobaoCloudpushNoticeIosAPIResponse) {
	v.Reset()
	poolTaobaoCloudpushNoticeIosAPIResponse.Put(v)
}

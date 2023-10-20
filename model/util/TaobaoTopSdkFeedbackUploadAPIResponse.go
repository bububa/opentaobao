package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSdkFeedbackUploadAPIResponse sdk信息回调 API返回值
// taobao.top.sdk.feedback.upload
//
// sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaoTopSdkFeedbackUploadAPIResponse struct {
	model.CommonResponse
	TaobaoTopSdkFeedbackUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopSdkFeedbackUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopSdkFeedbackUploadAPIResponseModel).Reset()
}

// TaobaoTopSdkFeedbackUploadAPIResponseModel is sdk信息回调 成功返回结果
type TaobaoTopSdkFeedbackUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"top_sdk_feedback_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 控制回传间隔（单位：秒）
	UploadInterval int64 `json:"upload_interval,omitempty" xml:"upload_interval,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopSdkFeedbackUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UploadInterval = 0
}

var poolTaobaoTopSdkFeedbackUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopSdkFeedbackUploadAPIResponse)
	},
}

// GetTaobaoTopSdkFeedbackUploadAPIResponse 从 sync.Pool 获取 TaobaoTopSdkFeedbackUploadAPIResponse
func GetTaobaoTopSdkFeedbackUploadAPIResponse() *TaobaoTopSdkFeedbackUploadAPIResponse {
	return poolTaobaoTopSdkFeedbackUploadAPIResponse.Get().(*TaobaoTopSdkFeedbackUploadAPIResponse)
}

// ReleaseTaobaoTopSdkFeedbackUploadAPIResponse 将 TaobaoTopSdkFeedbackUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopSdkFeedbackUploadAPIResponse(v *TaobaoTopSdkFeedbackUploadAPIResponse) {
	v.Reset()
	poolTaobaoTopSdkFeedbackUploadAPIResponse.Put(v)
}

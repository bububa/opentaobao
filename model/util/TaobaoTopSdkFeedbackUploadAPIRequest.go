package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopSdkFeedbackUploadAPIRequest
sdk信息回调 API请求
taobao.top.sdk.feedback.upload

sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性 */
type TaobaoTopSdkFeedbackUploadAPIRequest struct {
	model.Params
	// 1、回传加密信息
	_type string
	// 具体内容，json形式
	_content string
}

// NewTaobaoTopSdkFeedbackUploadRequest 初始化TaobaoTopSdkFeedbackUploadAPIRequest对象
func NewTaobaoTopSdkFeedbackUploadRequest() *TaobaoTopSdkFeedbackUploadAPIRequest {
	return &TaobaoTopSdkFeedbackUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetApiMethodName() string {
	return "taobao.top.sdk.feedback.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Type Setter
// 1、回传加密信息
func (r *TaobaoTopSdkFeedbackUploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetType() string {
	return r._type
}

// Set is Content Setter
// 具体内容，json形式
func (r *TaobaoTopSdkFeedbackUploadAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// Get Content Getter
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetContent() string {
	return r._content
}

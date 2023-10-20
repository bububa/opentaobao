package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopsdkfeedbackuploadAPIRequest sdk信息回调 API请求
// taobao.top.sdk.feedback.upload
//
// sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaotopsdkfeedbackuploadAPIRequest struct {
	model.Params
	// 1、回传加密信息
	_type string
	// 具体内容，json形式
	_content string
}

// NewTaobaotopsdkfeedbackuploadRequest 初始化TaobaotopsdkfeedbackuploadAPIRequest对象
func NewTaobaotopsdkfeedbackuploadRequest() *TaobaotopsdkfeedbackuploadAPIRequest {
	return &TaobaotopsdkfeedbackuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopsdkfeedbackuploadAPIRequest) GetApiMethodName() string {
	return "taobao.top.sdk.feedback.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopsdkfeedbackuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopsdkfeedbackuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 1、回传加密信息
func (r *TaobaotopsdkfeedbackuploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaotopsdkfeedbackuploadAPIRequest) GetType() string {
	return r._type
}

// SetContent is Content Setter
// 具体内容，json形式
func (r *TaobaotopsdkfeedbackuploadAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaotopsdkfeedbackuploadAPIRequest) GetContent() string {
	return r._content
}

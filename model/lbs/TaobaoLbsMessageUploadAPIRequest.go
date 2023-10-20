package lbs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolbsmessageuploadAPIRequest lbs数据采集 API请求
// taobao.lbs.message.upload
//
// lbs数据采集
type TaobaolbsmessageuploadAPIRequest struct {
	model.Params
	// 消息TOPIC
	_topic string
	// 消息体 json结构
	_body string
}

// NewTaobaolbsmessageuploadRequest 初始化TaobaolbsmessageuploadAPIRequest对象
func NewTaobaolbsmessageuploadRequest() *TaobaolbsmessageuploadAPIRequest {
	return &TaobaolbsmessageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolbsmessageuploadAPIRequest) GetApiMethodName() string {
	return "taobao.lbs.message.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolbsmessageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolbsmessageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopic is Topic Setter
// 消息TOPIC
func (r *TaobaolbsmessageuploadAPIRequest) SetTopic(_topic string) error {
	r._topic = _topic
	r.Set("topic", _topic)
	return nil
}

// GetTopic Topic Getter
func (r TaobaolbsmessageuploadAPIRequest) GetTopic() string {
	return r._topic
}

// SetBody is Body Setter
// 消息体 json结构
func (r *TaobaolbsmessageuploadAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r TaobaolbsmessageuploadAPIRequest) GetBody() string {
	return r._body
}

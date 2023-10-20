package lbs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLbsMessageUploadAPIRequest lbs数据采集 API请求
// taobao.lbs.message.upload
//
// lbs数据采集
type TaobaoLbsMessageUploadAPIRequest struct {
	model.Params
	// 消息TOPIC
	_topic string
	// 消息体 json结构
	_body string
}

// NewTaobaoLbsMessageUploadRequest 初始化TaobaoLbsMessageUploadAPIRequest对象
func NewTaobaoLbsMessageUploadRequest() *TaobaoLbsMessageUploadAPIRequest {
	return &TaobaoLbsMessageUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLbsMessageUploadAPIRequest) Reset() {
	r._topic = ""
	r._body = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLbsMessageUploadAPIRequest) GetApiMethodName() string {
	return "taobao.lbs.message.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLbsMessageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLbsMessageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopic is Topic Setter
// 消息TOPIC
func (r *TaobaoLbsMessageUploadAPIRequest) SetTopic(_topic string) error {
	r._topic = _topic
	r.Set("topic", _topic)
	return nil
}

// GetTopic Topic Getter
func (r TaobaoLbsMessageUploadAPIRequest) GetTopic() string {
	return r._topic
}

// SetBody is Body Setter
// 消息体 json结构
func (r *TaobaoLbsMessageUploadAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r TaobaoLbsMessageUploadAPIRequest) GetBody() string {
	return r._body
}

var poolTaobaoLbsMessageUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLbsMessageUploadRequest()
	},
}

// GetTaobaoLbsMessageUploadRequest 从 sync.Pool 获取 TaobaoLbsMessageUploadAPIRequest
func GetTaobaoLbsMessageUploadAPIRequest() *TaobaoLbsMessageUploadAPIRequest {
	return poolTaobaoLbsMessageUploadAPIRequest.Get().(*TaobaoLbsMessageUploadAPIRequest)
}

// ReleaseTaobaoLbsMessageUploadAPIRequest 将 TaobaoLbsMessageUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoLbsMessageUploadAPIRequest(v *TaobaoLbsMessageUploadAPIRequest) {
	v.Reset()
	poolTaobaoLbsMessageUploadAPIRequest.Put(v)
}

package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocloudpushmessageandroidAPIRequest 百川云推送发送消息给android API请求
// taobao.cloudpush.message.android
//
// 百川用户使用云推送发送消息给android
type TaobaocloudpushmessageandroidAPIRequest struct {
	model.Params
	// 发送的消息内容.
	_body string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
}

// NewTaobaocloudpushmessageandroidRequest 初始化TaobaocloudpushmessageandroidAPIRequest对象
func NewTaobaocloudpushmessageandroidRequest() *TaobaocloudpushmessageandroidAPIRequest {
	return &TaobaocloudpushmessageandroidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocloudpushmessageandroidAPIRequest) GetApiMethodName() string {
	return "taobao.cloudpush.message.android"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocloudpushmessageandroidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocloudpushmessageandroidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 发送的消息内容.
func (r *TaobaocloudpushmessageandroidAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r TaobaocloudpushmessageandroidAPIRequest) GetBody() string {
	return r._body
}

// SetTarget is Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaocloudpushmessageandroidAPIRequest) SetTarget(_target string) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// GetTarget Target Getter
func (r TaobaocloudpushmessageandroidAPIRequest) GetTarget() string {
	return r._target
}

// SetTargetValue is TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaocloudpushmessageandroidAPIRequest) SetTargetValue(_targetValue string) error {
	r._targetValue = _targetValue
	r.Set("target_value", _targetValue)
	return nil
}

// GetTargetValue TargetValue Getter
func (r TaobaocloudpushmessageandroidAPIRequest) GetTargetValue() string {
	return r._targetValue
}

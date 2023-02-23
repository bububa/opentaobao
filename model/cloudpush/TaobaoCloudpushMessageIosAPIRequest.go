package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudpushMessageIosAPIRequest 百川云推送发送消息给ios API请求
// taobao.cloudpush.message.ios
//
// 百川云推送发送消息给iOS设备.
type TaobaoCloudpushMessageIosAPIRequest struct {
	model.Params
	// 发送的消息内容.
	_body string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
}

// NewTaobaoCloudpushMessageIosRequest 初始化TaobaoCloudpushMessageIosAPIRequest对象
func NewTaobaoCloudpushMessageIosRequest() *TaobaoCloudpushMessageIosAPIRequest {
	return &TaobaoCloudpushMessageIosAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushMessageIosAPIRequest) GetApiMethodName() string {
	return "taobao.cloudpush.message.ios"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushMessageIosAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCloudpushMessageIosAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 发送的消息内容.
func (r *TaobaoCloudpushMessageIosAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r TaobaoCloudpushMessageIosAPIRequest) GetBody() string {
	return r._body
}

// SetTarget is Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushMessageIosAPIRequest) SetTarget(_target string) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// GetTarget Target Getter
func (r TaobaoCloudpushMessageIosAPIRequest) GetTarget() string {
	return r._target
}

// SetTargetValue is TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushMessageIosAPIRequest) SetTargetValue(_targetValue string) error {
	r._targetValue = _targetValue
	r.Set("target_value", _targetValue)
	return nil
}

// GetTargetValue TargetValue Getter
func (r TaobaoCloudpushMessageIosAPIRequest) GetTargetValue() string {
	return r._targetValue
}

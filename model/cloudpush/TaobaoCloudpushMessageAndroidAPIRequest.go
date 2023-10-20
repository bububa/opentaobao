package cloudpush

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudpushMessageAndroidAPIRequest 百川云推送发送消息给android API请求
// taobao.cloudpush.message.android
//
// 百川用户使用云推送发送消息给android
type TaobaoCloudpushMessageAndroidAPIRequest struct {
	model.Params
	// 发送的消息内容.
	_body string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
}

// NewTaobaoCloudpushMessageAndroidRequest 初始化TaobaoCloudpushMessageAndroidAPIRequest对象
func NewTaobaoCloudpushMessageAndroidRequest() *TaobaoCloudpushMessageAndroidAPIRequest {
	return &TaobaoCloudpushMessageAndroidAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCloudpushMessageAndroidAPIRequest) Reset() {
	r._body = ""
	r._target = ""
	r._targetValue = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushMessageAndroidAPIRequest) GetApiMethodName() string {
	return "taobao.cloudpush.message.android"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushMessageAndroidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCloudpushMessageAndroidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 发送的消息内容.
func (r *TaobaoCloudpushMessageAndroidAPIRequest) SetBody(_body string) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r TaobaoCloudpushMessageAndroidAPIRequest) GetBody() string {
	return r._body
}

// SetTarget is Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushMessageAndroidAPIRequest) SetTarget(_target string) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// GetTarget Target Getter
func (r TaobaoCloudpushMessageAndroidAPIRequest) GetTarget() string {
	return r._target
}

// SetTargetValue is TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushMessageAndroidAPIRequest) SetTargetValue(_targetValue string) error {
	r._targetValue = _targetValue
	r.Set("target_value", _targetValue)
	return nil
}

// GetTargetValue TargetValue Getter
func (r TaobaoCloudpushMessageAndroidAPIRequest) GetTargetValue() string {
	return r._targetValue
}

var poolTaobaoCloudpushMessageAndroidAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCloudpushMessageAndroidRequest()
	},
}

// GetTaobaoCloudpushMessageAndroidRequest 从 sync.Pool 获取 TaobaoCloudpushMessageAndroidAPIRequest
func GetTaobaoCloudpushMessageAndroidAPIRequest() *TaobaoCloudpushMessageAndroidAPIRequest {
	return poolTaobaoCloudpushMessageAndroidAPIRequest.Get().(*TaobaoCloudpushMessageAndroidAPIRequest)
}

// ReleaseTaobaoCloudpushMessageAndroidAPIRequest 将 TaobaoCloudpushMessageAndroidAPIRequest 放入 sync.Pool
func ReleaseTaobaoCloudpushMessageAndroidAPIRequest(v *TaobaoCloudpushMessageAndroidAPIRequest) {
	v.Reset()
	poolTaobaoCloudpushMessageAndroidAPIRequest.Put(v)
}

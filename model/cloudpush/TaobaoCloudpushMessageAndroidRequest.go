package cloudpush

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送消息给android API请求
taobao.cloudpush.message.android

百川用户使用云推送发送消息给android
*/
type TaobaoCloudpushMessageAndroidRequest struct {
    model.Params
    // 发送的消息内容.
    _body   string
    // 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
    _target   string
    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
    _targetValue   string
}

// 初始化TaobaoCloudpushMessageAndroidRequest对象
func NewTaobaoCloudpushMessageAndroidRequest() *TaobaoCloudpushMessageAndroidRequest{
    return &TaobaoCloudpushMessageAndroidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushMessageAndroidRequest) GetApiMethodName() string {
    return "taobao.cloudpush.message.android"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushMessageAndroidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Body Setter
// 发送的消息内容.
func (r *TaobaoCloudpushMessageAndroidRequest) SetBody(_body string) error {
    r._body = _body
    r.Set("body", _body)
    return nil
}

// Body Getter
func (r TaobaoCloudpushMessageAndroidRequest) GetBody() string {
    return r._body
}
// Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushMessageAndroidRequest) SetTarget(_target string) error {
    r._target = _target
    r.Set("target", _target)
    return nil
}

// Target Getter
func (r TaobaoCloudpushMessageAndroidRequest) GetTarget() string {
    return r._target
}
// TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushMessageAndroidRequest) SetTargetValue(_targetValue string) error {
    r._targetValue = _targetValue
    r.Set("target_value", _targetValue)
    return nil
}

// TargetValue Getter
func (r TaobaoCloudpushMessageAndroidRequest) GetTargetValue() string {
    return r._targetValue
}

package cloudpush

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送通知给ios设备 API请求
taobao.cloudpush.notice.ios

推送通知给ios设备
*/
type TaobaoCloudpushNoticeIosAPIRequest struct {
    model.Params
    // 通知摘要
    _summary   string
    // 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
    _target   string
    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
    _targetValue   string
    // iOS的通知是通过APNS中心来发送的，需要填写对应的环境信息.  DEV:表示开发环境, PRODUCT: 表示生产环境.
    _env   string
    // 提供给IOS通知的扩展属性，如角标或者声音等,注意：参数值为json
    _ext   string
}

// 初始化TaobaoCloudpushNoticeIosAPIRequest对象
func NewTaobaoCloudpushNoticeIosRequest() *TaobaoCloudpushNoticeIosAPIRequest{
    return &TaobaoCloudpushNoticeIosAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushNoticeIosAPIRequest) GetApiMethodName() string {
    return "taobao.cloudpush.notice.ios"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushNoticeIosAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Summary Setter
// 通知摘要
func (r *TaobaoCloudpushNoticeIosAPIRequest) SetSummary(_summary string) error {
    r._summary = _summary
    r.Set("summary", _summary)
    return nil
}

// Summary Getter
func (r TaobaoCloudpushNoticeIosAPIRequest) GetSummary() string {
    return r._summary
}
// Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushNoticeIosAPIRequest) SetTarget(_target string) error {
    r._target = _target
    r.Set("target", _target)
    return nil
}

// Target Getter
func (r TaobaoCloudpushNoticeIosAPIRequest) GetTarget() string {
    return r._target
}
// TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushNoticeIosAPIRequest) SetTargetValue(_targetValue string) error {
    r._targetValue = _targetValue
    r.Set("target_value", _targetValue)
    return nil
}

// TargetValue Getter
func (r TaobaoCloudpushNoticeIosAPIRequest) GetTargetValue() string {
    return r._targetValue
}
// Env Setter
// iOS的通知是通过APNS中心来发送的，需要填写对应的环境信息.  DEV:表示开发环境, PRODUCT: 表示生产环境.
func (r *TaobaoCloudpushNoticeIosAPIRequest) SetEnv(_env string) error {
    r._env = _env
    r.Set("env", _env)
    return nil
}

// Env Getter
func (r TaobaoCloudpushNoticeIosAPIRequest) GetEnv() string {
    return r._env
}
// Ext Setter
// 提供给IOS通知的扩展属性，如角标或者声音等,注意：参数值为json
func (r *TaobaoCloudpushNoticeIosAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoCloudpushNoticeIosAPIRequest) GetExt() string {
    return r._ext
}

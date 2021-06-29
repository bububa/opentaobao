package cloudpush

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送通知给android API请求
taobao.cloudpush.notice.android

百川云推送发送通知给android
*/
type TaobaoCloudpushNoticeAndroidRequest struct {
    model.Params
    // 通知摘要
    _summary   string
    // 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
    _target   string
    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
    _targetValue   string
    // 通知的标题.
    _title   string
}

// 初始化TaobaoCloudpushNoticeAndroidRequest对象
func NewTaobaoCloudpushNoticeAndroidRequest() *TaobaoCloudpushNoticeAndroidRequest{
    return &TaobaoCloudpushNoticeAndroidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushNoticeAndroidRequest) GetApiMethodName() string {
    return "taobao.cloudpush.notice.android"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushNoticeAndroidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Summary Setter
// 通知摘要
func (r *TaobaoCloudpushNoticeAndroidRequest) SetSummary(_summary string) error {
    r._summary = _summary
    r.Set("summary", _summary)
    return nil
}

// Summary Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetSummary() string {
    return r._summary
}
// Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushNoticeAndroidRequest) SetTarget(_target string) error {
    r._target = _target
    r.Set("target", _target)
    return nil
}

// Target Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetTarget() string {
    return r._target
}
// TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushNoticeAndroidRequest) SetTargetValue(_targetValue string) error {
    r._targetValue = _targetValue
    r.Set("target_value", _targetValue)
    return nil
}

// TargetValue Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetTargetValue() string {
    return r._targetValue
}
// Title Setter
// 通知的标题.
func (r *TaobaoCloudpushNoticeAndroidRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetTitle() string {
    return r._title
}

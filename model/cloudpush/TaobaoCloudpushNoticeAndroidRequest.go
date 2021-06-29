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
    summary   string
    // 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
    target   string
    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
    targetValue   string
    // 通知的标题.
    title   string
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
func (r *TaobaoCloudpushNoticeAndroidRequest) SetSummary(summary string) error {
    r.summary = summary
    r.Set("summary", summary)
    return nil
}

// Summary Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetSummary() string {
    return r.summary
}
// Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushNoticeAndroidRequest) SetTarget(target string) error {
    r.target = target
    r.Set("target", target)
    return nil
}

// Target Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetTarget() string {
    return r.target
}
// TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushNoticeAndroidRequest) SetTargetValue(targetValue string) error {
    r.targetValue = targetValue
    r.Set("target_value", targetValue)
    return nil
}

// TargetValue Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetTargetValue() string {
    return r.targetValue
}
// Title Setter
// 通知的标题.
func (r *TaobaoCloudpushNoticeAndroidRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoCloudpushNoticeAndroidRequest) GetTitle() string {
    return r.title
}

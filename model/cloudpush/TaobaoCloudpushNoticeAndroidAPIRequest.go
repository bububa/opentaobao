package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudpushNoticeAndroidAPIRequest 百川云推送发送通知给android API请求
// taobao.cloudpush.notice.android
//
// 百川云推送发送通知给android
type TaobaoCloudpushNoticeAndroidAPIRequest struct {
	model.Params
	// 通知摘要
	_summary string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
	// 通知的标题.
	_title string
}

// NewTaobaoCloudpushNoticeAndroidRequest 初始化TaobaoCloudpushNoticeAndroidAPIRequest对象
func NewTaobaoCloudpushNoticeAndroidRequest() *TaobaoCloudpushNoticeAndroidAPIRequest {
	return &TaobaoCloudpushNoticeAndroidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetApiMethodName() string {
	return "taobao.cloudpush.notice.android"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSummary is Summary Setter
// 通知摘要
func (r *TaobaoCloudpushNoticeAndroidAPIRequest) SetSummary(_summary string) error {
	r._summary = _summary
	r.Set("summary", _summary)
	return nil
}

// GetSummary Summary Getter
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetSummary() string {
	return r._summary
}

// SetTarget is Target Setter
// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
func (r *TaobaoCloudpushNoticeAndroidAPIRequest) SetTarget(_target string) error {
	r._target = _target
	r.Set("target", _target)
	return nil
}

// GetTarget Target Getter
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetTarget() string {
	return r._target
}

// SetTargetValue is TargetValue Setter
// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
func (r *TaobaoCloudpushNoticeAndroidAPIRequest) SetTargetValue(_targetValue string) error {
	r._targetValue = _targetValue
	r.Set("target_value", _targetValue)
	return nil
}

// GetTargetValue TargetValue Getter
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetTargetValue() string {
	return r._targetValue
}

// SetTitle is Title Setter
// 通知的标题.
func (r *TaobaoCloudpushNoticeAndroidAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoCloudpushNoticeAndroidAPIRequest) GetTitle() string {
	return r._title
}

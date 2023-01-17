package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJdpUserAddAPIRequest 添加数据推送用户 API请求
// taobao.jushita.jdp.user.add
//
// 提供给接入数据推送的应用添加数据推送服务的用户
type TaobaoJushitaJdpUserAddAPIRequest struct {
	model.Params
	// RDS实例名称
	_rdsName string
	// 推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。
	_historyDays int64
}

// NewTaobaoJushitaJdpUserAddRequest 初始化TaobaoJushitaJdpUserAddAPIRequest对象
func NewTaobaoJushitaJdpUserAddRequest() *TaobaoJushitaJdpUserAddAPIRequest {
	return &TaobaoJushitaJdpUserAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJdpUserAddAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJdpUserAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJdpUserAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRdsName is RdsName Setter
// RDS实例名称
func (r *TaobaoJushitaJdpUserAddAPIRequest) SetRdsName(_rdsName string) error {
	r._rdsName = _rdsName
	r.Set("rds_name", _rdsName)
	return nil
}

// GetRdsName RdsName Getter
func (r TaobaoJushitaJdpUserAddAPIRequest) GetRdsName() string {
	return r._rdsName
}

// SetHistoryDays is HistoryDays Setter
// 推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。
func (r *TaobaoJushitaJdpUserAddAPIRequest) SetHistoryDays(_historyDays int64) error {
	r._historyDays = _historyDays
	r.Set("history_days", _historyDays)
	return nil
}

// GetHistoryDays HistoryDays Getter
func (r TaobaoJushitaJdpUserAddAPIRequest) GetHistoryDays() int64 {
	return r._historyDays
}

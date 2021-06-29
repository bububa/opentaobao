package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加数据推送用户 API请求
taobao.jushita.jdp.user.add

提供给接入数据推送的应用添加数据推送服务的用户
*/
type TaobaoJushitaJdpUserAddRequest struct {
    model.Params
    // RDS实例名称
    rdsName   string
    // 推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。
    historyDays   int64
}

// 初始化TaobaoJushitaJdpUserAddRequest对象
func NewTaobaoJushitaJdpUserAddRequest() *TaobaoJushitaJdpUserAddRequest{
    return &TaobaoJushitaJdpUserAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJdpUserAddRequest) GetApiMethodName() string {
    return "taobao.jushita.jdp.user.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJdpUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RdsName Setter
// RDS实例名称
func (r *TaobaoJushitaJdpUserAddRequest) SetRdsName(rdsName string) error {
    r.rdsName = rdsName
    r.Set("rds_name", rdsName)
    return nil
}

// RdsName Getter
func (r TaobaoJushitaJdpUserAddRequest) GetRdsName() string {
    return r.rdsName
}
// HistoryDays Setter
// 推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。
func (r *TaobaoJushitaJdpUserAddRequest) SetHistoryDays(historyDays int64) error {
    r.historyDays = historyDays
    r.Set("history_days", historyDays)
    return nil
}

// HistoryDays Getter
func (r TaobaoJushitaJdpUserAddRequest) GetHistoryDays() int64 {
    return r.historyDays
}

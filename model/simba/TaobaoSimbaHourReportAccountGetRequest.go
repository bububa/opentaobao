package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户级别小时报表获取 API请求
taobao.simba.hour.report.account.get

获取账户小时实时报表数据
*/
type TaobaoSimbaHourReportAccountGetRequest struct {
    model.Params
    // 昵称
    _nick   string
    // 时间
    _theDate   string
    // 当前小时
    _hour   string
}

// 初始化TaobaoSimbaHourReportAccountGetRequest对象
func NewTaobaoSimbaHourReportAccountGetRequest() *TaobaoSimbaHourReportAccountGetRequest{
    return &TaobaoSimbaHourReportAccountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaHourReportAccountGetRequest) GetApiMethodName() string {
    return "taobao.simba.hour.report.account.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaHourReportAccountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaHourReportAccountGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaHourReportAccountGetRequest) GetNick() string {
    return r._nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportAccountGetRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaHourReportAccountGetRequest) GetTheDate() string {
    return r._theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportAccountGetRequest) SetHour(_hour string) error {
    r._hour = _hour
    r.Set("hour", _hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaHourReportAccountGetRequest) GetHour() string {
    return r._hour
}

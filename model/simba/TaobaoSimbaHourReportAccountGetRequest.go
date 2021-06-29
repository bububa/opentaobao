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
    nick   string
    // 时间
    theDate   string
    // 当前小时
    hour   string
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
func (r *TaobaoSimbaHourReportAccountGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaHourReportAccountGetRequest) GetNick() string {
    return r.nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportAccountGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaHourReportAccountGetRequest) GetTheDate() string {
    return r.theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportAccountGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaHourReportAccountGetRequest) GetHour() string {
    return r.hour
}

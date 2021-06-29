package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取账户实时报表数据 API请求
taobao.simba.rtrpt.cust.get

获取账户实时报表数据
*/
type TaobaoSimbaRtrptCustGetRequest struct {
    model.Params
    // 昵称
    nick   string
    // 日期，格式yyyy-mm-dd
    theDate   string
}

// 初始化TaobaoSimbaRtrptCustGetRequest对象
func NewTaobaoSimbaRtrptCustGetRequest() *TaobaoSimbaRtrptCustGetRequest{
    return &TaobaoSimbaRtrptCustGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCustGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.cust.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRtrptCustGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptCustGetRequest) GetNick() string {
    return r.nick
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCustGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptCustGetRequest) GetTheDate() string {
    return r.theDate
}

package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取账户实时报表数据 APIRequest
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

func NewTaobaoSimbaRtrptCustGetRequest() *TaobaoSimbaRtrptCustGetRequest{
    return &TaobaoSimbaRtrptCustGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRtrptCustGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.cust.get"
}

func (r TaobaoSimbaRtrptCustGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRtrptCustGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRtrptCustGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRtrptCustGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaRtrptCustGetRequest) GetTheDate() string {
    return r.theDate
}


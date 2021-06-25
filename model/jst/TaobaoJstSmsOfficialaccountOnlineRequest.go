package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号上线 APIRequest
taobao.jst.sms.officialaccount.online

聚石塔公众号上线
*/
type TaobaoJstSmsOfficialaccountOnlineRequest struct {
    model.Params

    // 公众号上线请求参数
    officialAccountOnlineRequest   *JstBaseRequest 

}

func NewTaobaoJstSmsOfficialaccountOnlineRequest() *TaobaoJstSmsOfficialaccountOnlineRequest{
    return &TaobaoJstSmsOfficialaccountOnlineRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsOfficialaccountOnlineRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.online"
}

func (r TaobaoJstSmsOfficialaccountOnlineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsOfficialaccountOnlineRequest) SetOfficialAccountOnlineRequest(officialAccountOnlineRequest *JstBaseRequest) error {
    r.officialAccountOnlineRequest = officialAccountOnlineRequest
    r.Set("official_account_online_request", officialAccountOnlineRequest)
    return nil
}

func (r TaobaoJstSmsOfficialaccountOnlineRequest) GetOfficialAccountOnlineRequest() *JstBaseRequest {
    return r.officialAccountOnlineRequest
}


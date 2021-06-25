package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号下线 APIRequest
taobao.jst.sms.officialaccount.offline

聚石塔公众号下线
*/
type TaobaoJstSmsOfficialaccountOfflineRequest struct {
    model.Params

    // 公众号下线请求
    officialAccountOffline   *JstBaseRequest 

}

func NewTaobaoJstSmsOfficialaccountOfflineRequest() *TaobaoJstSmsOfficialaccountOfflineRequest{
    return &TaobaoJstSmsOfficialaccountOfflineRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsOfficialaccountOfflineRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.offline"
}

func (r TaobaoJstSmsOfficialaccountOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsOfficialaccountOfflineRequest) SetOfficialAccountOffline(officialAccountOffline *JstBaseRequest) error {
    r.officialAccountOffline = officialAccountOffline
    r.Set("official_account_offline", officialAccountOffline)
    return nil
}

func (r TaobaoJstSmsOfficialaccountOfflineRequest) GetOfficialAccountOffline() *JstBaseRequest {
    return r.officialAccountOffline
}


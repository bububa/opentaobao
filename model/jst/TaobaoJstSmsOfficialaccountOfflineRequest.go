package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号下线 API请求
taobao.jst.sms.officialaccount.offline

聚石塔公众号下线
*/
type TaobaoJstSmsOfficialaccountOfflineRequest struct {
    model.Params
    // 公众号下线请求
    officialAccountOffline   *JstBaseRequest
}

// 初始化TaobaoJstSmsOfficialaccountOfflineRequest对象
func NewTaobaoJstSmsOfficialaccountOfflineRequest() *TaobaoJstSmsOfficialaccountOfflineRequest{
    return &TaobaoJstSmsOfficialaccountOfflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountOfflineRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.offline"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountOfflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAccountOffline Setter
// 公众号下线请求
func (r *TaobaoJstSmsOfficialaccountOfflineRequest) SetOfficialAccountOffline(officialAccountOffline *JstBaseRequest) error {
    r.officialAccountOffline = officialAccountOffline
    r.Set("official_account_offline", officialAccountOffline)
    return nil
}

// OfficialAccountOffline Getter
func (r TaobaoJstSmsOfficialaccountOfflineRequest) GetOfficialAccountOffline() *JstBaseRequest {
    return r.officialAccountOffline
}

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
type TaobaoJstSmsOfficialaccountOfflineAPIRequest struct {
    model.Params
    // 公众号下线请求
    _officialAccountOffline   *JstBaseRequest
}

// 初始化TaobaoJstSmsOfficialaccountOfflineAPIRequest对象
func NewTaobaoJstSmsOfficialaccountOfflineRequest() *TaobaoJstSmsOfficialaccountOfflineAPIRequest{
    return &TaobaoJstSmsOfficialaccountOfflineAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountOfflineAPIRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.offline"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountOfflineAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAccountOffline Setter
// 公众号下线请求
func (r *TaobaoJstSmsOfficialaccountOfflineAPIRequest) SetOfficialAccountOffline(_officialAccountOffline *JstBaseRequest) error {
    r._officialAccountOffline = _officialAccountOffline
    r.Set("official_account_offline", _officialAccountOffline)
    return nil
}

// OfficialAccountOffline Getter
func (r TaobaoJstSmsOfficialaccountOfflineAPIRequest) GetOfficialAccountOffline() *JstBaseRequest {
    return r._officialAccountOffline
}

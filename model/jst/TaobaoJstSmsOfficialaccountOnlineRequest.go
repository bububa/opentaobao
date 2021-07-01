package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号上线 API请求
taobao.jst.sms.officialaccount.online

聚石塔公众号上线
*/
type TaobaoJstSmsOfficialaccountOnlineAPIRequest struct {
    model.Params
    // 公众号上线请求参数
    _officialAccountOnlineRequest   *JstBaseRequest
}

// 初始化TaobaoJstSmsOfficialaccountOnlineAPIRequest对象
func NewTaobaoJstSmsOfficialaccountOnlineRequest() *TaobaoJstSmsOfficialaccountOnlineAPIRequest{
    return &TaobaoJstSmsOfficialaccountOnlineAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountOnlineAPIRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.online"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountOnlineAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAccountOnlineRequest Setter
// 公众号上线请求参数
func (r *TaobaoJstSmsOfficialaccountOnlineAPIRequest) SetOfficialAccountOnlineRequest(_officialAccountOnlineRequest *JstBaseRequest) error {
    r._officialAccountOnlineRequest = _officialAccountOnlineRequest
    r.Set("official_account_online_request", _officialAccountOnlineRequest)
    return nil
}

// OfficialAccountOnlineRequest Getter
func (r TaobaoJstSmsOfficialaccountOnlineAPIRequest) GetOfficialAccountOnlineRequest() *JstBaseRequest {
    return r._officialAccountOnlineRequest
}

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
type TaobaoJstSmsOfficialaccountOnlineRequest struct {
    model.Params
    // 公众号上线请求参数
    _officialAccountOnlineRequest   *JstBaseRequest
}

// 初始化TaobaoJstSmsOfficialaccountOnlineRequest对象
func NewTaobaoJstSmsOfficialaccountOnlineRequest() *TaobaoJstSmsOfficialaccountOnlineRequest{
    return &TaobaoJstSmsOfficialaccountOnlineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountOnlineRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.online"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountOnlineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAccountOnlineRequest Setter
// 公众号上线请求参数
func (r *TaobaoJstSmsOfficialaccountOnlineRequest) SetOfficialAccountOnlineRequest(_officialAccountOnlineRequest *JstBaseRequest) error {
    r._officialAccountOnlineRequest = _officialAccountOnlineRequest
    r.Set("official_account_online_request", _officialAccountOnlineRequest)
    return nil
}

// OfficialAccountOnlineRequest Getter
func (r TaobaoJstSmsOfficialaccountOnlineRequest) GetOfficialAccountOnlineRequest() *JstBaseRequest {
    return r._officialAccountOnlineRequest
}

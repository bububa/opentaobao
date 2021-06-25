package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号状态查询 APIRequest
taobao.jst.sms.status.query

聚石塔公众号状态查询
*/
type TaobaoJstSmsStatusQueryRequest struct {
    model.Params

    // 公众号状态信息查询请求
    officialAccountStatusQueryRequest   *JstBaseRequest 

}

func NewTaobaoJstSmsStatusQueryRequest() *TaobaoJstSmsStatusQueryRequest{
    return &TaobaoJstSmsStatusQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsStatusQueryRequest) GetApiMethodName() string {
    return "taobao.jst.sms.status.query"
}

func (r TaobaoJstSmsStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsStatusQueryRequest) SetOfficialAccountStatusQueryRequest(officialAccountStatusQueryRequest *JstBaseRequest) error {
    r.officialAccountStatusQueryRequest = officialAccountStatusQueryRequest
    r.Set("official_account_status_query_request", officialAccountStatusQueryRequest)
    return nil
}

func (r TaobaoJstSmsStatusQueryRequest) GetOfficialAccountStatusQueryRequest() *JstBaseRequest {
    return r.officialAccountStatusQueryRequest
}


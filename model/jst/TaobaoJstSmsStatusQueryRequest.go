package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号状态查询 API请求
taobao.jst.sms.status.query

聚石塔公众号状态查询
*/
type TaobaoJstSmsStatusQueryRequest struct {
    model.Params
    // 公众号状态信息查询请求
    officialAccountStatusQueryRequest   *JstBaseRequest
}

// 初始化TaobaoJstSmsStatusQueryRequest对象
func NewTaobaoJstSmsStatusQueryRequest() *TaobaoJstSmsStatusQueryRequest{
    return &TaobaoJstSmsStatusQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsStatusQueryRequest) GetApiMethodName() string {
    return "taobao.jst.sms.status.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAccountStatusQueryRequest Setter
// 公众号状态信息查询请求
func (r *TaobaoJstSmsStatusQueryRequest) SetOfficialAccountStatusQueryRequest(officialAccountStatusQueryRequest *JstBaseRequest) error {
    r.officialAccountStatusQueryRequest = officialAccountStatusQueryRequest
    r.Set("official_account_status_query_request", officialAccountStatusQueryRequest)
    return nil
}

// OfficialAccountStatusQueryRequest Getter
func (r TaobaoJstSmsStatusQueryRequest) GetOfficialAccountStatusQueryRequest() *JstBaseRequest {
    return r.officialAccountStatusQueryRequest
}

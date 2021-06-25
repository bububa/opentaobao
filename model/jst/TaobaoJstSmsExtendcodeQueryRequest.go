package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔扩展码查询 APIRequest
taobao.jst.sms.extendcode.query

聚石塔扩展码查询
*/
type TaobaoJstSmsExtendcodeQueryRequest struct {
    model.Params

    // 扩展码查询请求
    extendCodeQueryRequest   *ExtendCodeQueryRequest 

}

func NewTaobaoJstSmsExtendcodeQueryRequest() *TaobaoJstSmsExtendcodeQueryRequest{
    return &TaobaoJstSmsExtendcodeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsExtendcodeQueryRequest) GetApiMethodName() string {
    return "taobao.jst.sms.extendcode.query"
}

func (r TaobaoJstSmsExtendcodeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsExtendcodeQueryRequest) SetExtendCodeQueryRequest(extendCodeQueryRequest *ExtendCodeQueryRequest) error {
    r.extendCodeQueryRequest = extendCodeQueryRequest
    r.Set("extend_code_query_request", extendCodeQueryRequest)
    return nil
}

func (r TaobaoJstSmsExtendcodeQueryRequest) GetExtendCodeQueryRequest() *ExtendCodeQueryRequest {
    return r.extendCodeQueryRequest
}


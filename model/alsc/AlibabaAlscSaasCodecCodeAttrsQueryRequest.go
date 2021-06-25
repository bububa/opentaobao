package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码业务属性查询 APIRequest
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询
*/
type AlibabaAlscSaasCodecCodeAttrsQueryRequest struct {
    model.Params

    // 请求入参
    queryCodeRequest   *QueryCodeBizAttrRequest 

}

func NewAlibabaAlscSaasCodecCodeAttrsQueryRequest() *AlibabaAlscSaasCodecCodeAttrsQueryRequest{
    return &AlibabaAlscSaasCodecCodeAttrsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscSaasCodecCodeAttrsQueryRequest) GetApiMethodName() string {
    return "alibaba.alsc.saas.codec.code.attrs.query"
}

func (r AlibabaAlscSaasCodecCodeAttrsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscSaasCodecCodeAttrsQueryRequest) SetQueryCodeRequest(queryCodeRequest *QueryCodeBizAttrRequest) error {
    r.queryCodeRequest = queryCodeRequest
    r.Set("query_code_request", queryCodeRequest)
    return nil
}

func (r AlibabaAlscSaasCodecCodeAttrsQueryRequest) GetQueryCodeRequest() *QueryCodeBizAttrRequest {
    return r.queryCodeRequest
}


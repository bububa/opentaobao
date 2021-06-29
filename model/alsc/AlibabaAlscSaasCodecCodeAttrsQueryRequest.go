package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码业务属性查询 API请求
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询
*/
type AlibabaAlscSaasCodecCodeAttrsQueryRequest struct {
    model.Params
    // 请求入参
    queryCodeRequest   *QueryCodeBizAttrRequest
}

// 初始化AlibabaAlscSaasCodecCodeAttrsQueryRequest对象
func NewAlibabaAlscSaasCodecCodeAttrsQueryRequest() *AlibabaAlscSaasCodecCodeAttrsQueryRequest{
    return &AlibabaAlscSaasCodecCodeAttrsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscSaasCodecCodeAttrsQueryRequest) GetApiMethodName() string {
    return "alibaba.alsc.saas.codec.code.attrs.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscSaasCodecCodeAttrsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryCodeRequest Setter
// 请求入参
func (r *AlibabaAlscSaasCodecCodeAttrsQueryRequest) SetQueryCodeRequest(queryCodeRequest *QueryCodeBizAttrRequest) error {
    r.queryCodeRequest = queryCodeRequest
    r.Set("query_code_request", queryCodeRequest)
    return nil
}

// QueryCodeRequest Getter
func (r AlibabaAlscSaasCodecCodeAttrsQueryRequest) GetQueryCodeRequest() *QueryCodeBizAttrRequest {
    return r.queryCodeRequest
}

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
type AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest struct {
    model.Params
    // 请求入参
    _queryCodeRequest   *QueryCodeBizAttrRequest
}

// 初始化AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest对象
func NewAlibabaAlscSaasCodecCodeAttrsQueryRequest() *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest{
    return &AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.saas.codec.code.attrs.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryCodeRequest Setter
// 请求入参
func (r *AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) SetQueryCodeRequest(_queryCodeRequest *QueryCodeBizAttrRequest) error {
    r._queryCodeRequest = _queryCodeRequest
    r.Set("query_code_request", _queryCodeRequest)
    return nil
}

// QueryCodeRequest Getter
func (r AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest) GetQueryCodeRequest() *QueryCodeBizAttrRequest {
    return r._queryCodeRequest
}

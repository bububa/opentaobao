package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔扩展码查询 API请求
taobao.jst.sms.extendcode.query

聚石塔扩展码查询
*/
type TaobaoJstSmsExtendcodeQueryRequest struct {
    model.Params
    // 扩展码查询请求
    extendCodeQueryRequest   *ExtendCodeQueryRequest
}

// 初始化TaobaoJstSmsExtendcodeQueryRequest对象
func NewTaobaoJstSmsExtendcodeQueryRequest() *TaobaoJstSmsExtendcodeQueryRequest{
    return &TaobaoJstSmsExtendcodeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsExtendcodeQueryRequest) GetApiMethodName() string {
    return "taobao.jst.sms.extendcode.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsExtendcodeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendCodeQueryRequest Setter
// 扩展码查询请求
func (r *TaobaoJstSmsExtendcodeQueryRequest) SetExtendCodeQueryRequest(extendCodeQueryRequest *ExtendCodeQueryRequest) error {
    r.extendCodeQueryRequest = extendCodeQueryRequest
    r.Set("extend_code_query_request", extendCodeQueryRequest)
    return nil
}

// ExtendCodeQueryRequest Getter
func (r TaobaoJstSmsExtendcodeQueryRequest) GetExtendCodeQueryRequest() *ExtendCodeQueryRequest {
    return r.extendCodeQueryRequest
}

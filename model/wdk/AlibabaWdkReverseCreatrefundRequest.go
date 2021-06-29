package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向提交 API请求
alibaba.wdk.reverse.creatrefund

逆向申请提交
*/
type AlibabaWdkReverseCreatrefundRequest struct {
    model.Params
    // CreateReverseReq
    paramCreateReverseReq   *CreateReverseReq
}

// 初始化AlibabaWdkReverseCreatrefundRequest对象
func NewAlibabaWdkReverseCreatrefundRequest() *AlibabaWdkReverseCreatrefundRequest{
    return &AlibabaWdkReverseCreatrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseCreatrefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.creatrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseCreatrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCreateReverseReq Setter
// CreateReverseReq
func (r *AlibabaWdkReverseCreatrefundRequest) SetParamCreateReverseReq(paramCreateReverseReq *CreateReverseReq) error {
    r.paramCreateReverseReq = paramCreateReverseReq
    r.Set("param_create_reverse_req", paramCreateReverseReq)
    return nil
}

// ParamCreateReverseReq Getter
func (r AlibabaWdkReverseCreatrefundRequest) GetParamCreateReverseReq() *CreateReverseReq {
    return r.paramCreateReverseReq
}

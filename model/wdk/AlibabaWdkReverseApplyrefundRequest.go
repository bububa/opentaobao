package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向申请接口 API请求
alibaba.wdk.reverse.applyrefund

逆向渲染
*/
type AlibabaWdkReverseApplyrefundRequest struct {
    model.Params
    // 入参
    _paramApplyReverseReq   *ApplyReverseReq
}

// 初始化AlibabaWdkReverseApplyrefundRequest对象
func NewAlibabaWdkReverseApplyrefundRequest() *AlibabaWdkReverseApplyrefundRequest{
    return &AlibabaWdkReverseApplyrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseApplyrefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.applyrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseApplyrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamApplyReverseReq Setter
// 入参
func (r *AlibabaWdkReverseApplyrefundRequest) SetParamApplyReverseReq(_paramApplyReverseReq *ApplyReverseReq) error {
    r._paramApplyReverseReq = _paramApplyReverseReq
    r.Set("param_apply_reverse_req", _paramApplyReverseReq)
    return nil
}

// ParamApplyReverseReq Getter
func (r AlibabaWdkReverseApplyrefundRequest) GetParamApplyReverseReq() *ApplyReverseReq {
    return r._paramApplyReverseReq
}

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
type AlibabaWdkReverseApplyrefundAPIRequest struct {
    model.Params
    // 入参
    _paramApplyReverseReq   *ApplyReverseReq
}

// 初始化AlibabaWdkReverseApplyrefundAPIRequest对象
func NewAlibabaWdkReverseApplyrefundRequest() *AlibabaWdkReverseApplyrefundAPIRequest{
    return &AlibabaWdkReverseApplyrefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.applyrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamApplyReverseReq Setter
// 入参
func (r *AlibabaWdkReverseApplyrefundAPIRequest) SetParamApplyReverseReq(_paramApplyReverseReq *ApplyReverseReq) error {
    r._paramApplyReverseReq = _paramApplyReverseReq
    r.Set("param_apply_reverse_req", _paramApplyReverseReq)
    return nil
}

// ParamApplyReverseReq Getter
func (r AlibabaWdkReverseApplyrefundAPIRequest) GetParamApplyReverseReq() *ApplyReverseReq {
    return r._paramApplyReverseReq
}

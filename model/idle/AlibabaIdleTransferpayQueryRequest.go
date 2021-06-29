package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼转账结果查询 API请求
alibaba.idle.transferpay.query

商家业务 转账支付的结果查询
*/
type AlibabaIdleTransferpayQueryRequest struct {
    model.Params
    // 入参
    param   *PayQueryRequest
}

// 初始化AlibabaIdleTransferpayQueryRequest对象
func NewAlibabaIdleTransferpayQueryRequest() *AlibabaIdleTransferpayQueryRequest{
    return &AlibabaIdleTransferpayQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleTransferpayQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.transferpay.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleTransferpayQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaIdleTransferpayQueryRequest) SetParam(param *PayQueryRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaIdleTransferpayQueryRequest) GetParam() *PayQueryRequest {
    return r.param
}

package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户取消 API请求
alibaba.ele.fengniao.cancel.merchant

商户取消配送
*/
type AlibabaEleFengniaoCancelMerchantRequest struct {
    model.Params
    // 参数param
    _param   *Param
}

// 初始化AlibabaEleFengniaoCancelMerchantRequest对象
func NewAlibabaEleFengniaoCancelMerchantRequest() *AlibabaEleFengniaoCancelMerchantRequest{
    return &AlibabaEleFengniaoCancelMerchantRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCancelMerchantRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.cancel.merchant"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCancelMerchantRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数param
func (r *AlibabaEleFengniaoCancelMerchantRequest) SetParam(_param *Param) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoCancelMerchantRequest) GetParam() *Param {
    return r._param
}

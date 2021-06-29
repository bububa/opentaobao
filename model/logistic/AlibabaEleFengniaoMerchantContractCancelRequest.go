package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟商户解约接口 API请求
alibaba.ele.fengniao.merchant.contract.cancel

通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
*/
type AlibabaEleFengniaoMerchantContractCancelRequest struct {
    model.Params
    // 系统自动生成
    _param   *Param
}

// 初始化AlibabaEleFengniaoMerchantContractCancelRequest对象
func NewAlibabaEleFengniaoMerchantContractCancelRequest() *AlibabaEleFengniaoMerchantContractCancelRequest{
    return &AlibabaEleFengniaoMerchantContractCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoMerchantContractCancelRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.merchant.contract.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoMerchantContractCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoMerchantContractCancelRequest) SetParam(_param *Param) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoMerchantContractCancelRequest) GetParam() *Param {
    return r._param
}

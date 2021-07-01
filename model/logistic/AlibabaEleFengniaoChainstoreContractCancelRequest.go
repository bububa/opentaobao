package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店解约接口 API请求
alibaba.ele.fengniao.chainstore.contract.cancel

调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
*/
type AlibabaEleFengniaoChainstoreContractCancelAPIRequest struct {
    model.Params
    // 系统自动生成
    _param   *AlibabaEleFengniaoChainstoreContractCancelData
}

// 初始化AlibabaEleFengniaoChainstoreContractCancelAPIRequest对象
func NewAlibabaEleFengniaoChainstoreContractCancelRequest() *AlibabaEleFengniaoChainstoreContractCancelAPIRequest{
    return &AlibabaEleFengniaoChainstoreContractCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.contract.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoChainstoreContractCancelAPIRequest) SetParam(_param *AlibabaEleFengniaoChainstoreContractCancelData) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetParam() *AlibabaEleFengniaoChainstoreContractCancelData {
    return r._param
}

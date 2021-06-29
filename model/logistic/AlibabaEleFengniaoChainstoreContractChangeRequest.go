package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店改签合同接口 API请求
alibaba.ele.fengniao.chainstore.contract.change

通过调用接口，门店改签物流服务包
*/
type AlibabaEleFengniaoChainstoreContractChangeRequest struct {
    model.Params
    // 系统自动生成
    _param   *Param
}

// 初始化AlibabaEleFengniaoChainstoreContractChangeRequest对象
func NewAlibabaEleFengniaoChainstoreContractChangeRequest() *AlibabaEleFengniaoChainstoreContractChangeRequest{
    return &AlibabaEleFengniaoChainstoreContractChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreContractChangeRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.contract.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreContractChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoChainstoreContractChangeRequest) SetParam(_param *Param) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoChainstoreContractChangeRequest) GetParam() *Param {
    return r._param
}

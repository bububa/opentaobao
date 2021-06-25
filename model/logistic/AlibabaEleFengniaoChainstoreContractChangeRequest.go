package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店改签合同接口 APIRequest
alibaba.ele.fengniao.chainstore.contract.change

通过调用接口，门店改签物流服务包
*/
type AlibabaEleFengniaoChainstoreContractChangeRequest struct {
    model.Params

    // 系统自动生成
    param   *Param 

}

func NewAlibabaEleFengniaoChainstoreContractChangeRequest() *AlibabaEleFengniaoChainstoreContractChangeRequest{
    return &AlibabaEleFengniaoChainstoreContractChangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoChainstoreContractChangeRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.contract.change"
}

func (r AlibabaEleFengniaoChainstoreContractChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoChainstoreContractChangeRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoChainstoreContractChangeRequest) GetParam() *Param {
    return r.param
}


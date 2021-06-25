package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店解约接口 APIRequest
alibaba.ele.fengniao.chainstore.contract.cancel

调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
*/
type AlibabaEleFengniaoChainstoreContractCancelRequest struct {
    model.Params

    // 系统自动生成
    param   *AlibabaEleFengniaoChainstoreContractCancelData 

}

func NewAlibabaEleFengniaoChainstoreContractCancelRequest() *AlibabaEleFengniaoChainstoreContractCancelRequest{
    return &AlibabaEleFengniaoChainstoreContractCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoChainstoreContractCancelRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.contract.cancel"
}

func (r AlibabaEleFengniaoChainstoreContractCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoChainstoreContractCancelRequest) SetParam(param *AlibabaEleFengniaoChainstoreContractCancelData) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoChainstoreContractCancelRequest) GetParam() *AlibabaEleFengniaoChainstoreContractCancelData {
    return r.param
}


package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店解约接口 APIResponse
alibaba.ele.fengniao.chainstore.contract.cancel

调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
*/
type AlibabaEleFengniaoChainstoreContractCancelAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoChainstoreContractCancelResponse `json:"alibaba_ele_fengniao_chainstore_contract_cancel_response,omitempty"`
}

type AlibabaEleFengniaoChainstoreContractCancelResponse struct {

    // msg
    Message   string `json:"message,omitempty"`

}

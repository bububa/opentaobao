package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店改签合同接口 APIResponse
alibaba.ele.fengniao.chainstore.contract.change

通过调用接口，门店改签物流服务包
*/
type AlibabaEleFengniaoChainstoreContractChangeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaEleFengniaoChainstoreContractChangeResponse `json:"alibaba_ele_fengniao_chainstore_contract_change_response,omitempty"` 
    AlibabaEleFengniaoChainstoreContractChangeResponse
}

/* model for simplify = false
type AlibabaEleFengniaoChainstoreContractChangeResponse struct {

    // msg
    
    Message   string `json:"message,omitempty"`
    

}
*/

type AlibabaEleFengniaoChainstoreContractChangeResponse struct {

    // msg
    Message   string `json:"message,omitempty"`

}

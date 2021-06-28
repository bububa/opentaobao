package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店改签合同接口 APIResponse
alibaba.ele.fengniao.chainstore.contract.change

通过调用接口，门店改签物流服务包
*/
type AlibabaEleFengniaoChainstoreContractChangeAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoChainstoreContractChangeResponse
}

type AlibabaEleFengniaoChainstoreContractChangeResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_contract_change_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // msg
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}

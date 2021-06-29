package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟商户解约接口 APIResponse
alibaba.ele.fengniao.merchant.contract.cancel

通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
*/
type AlibabaEleFengniaoMerchantContractCancelAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoMerchantContractCancelResponse
}

type AlibabaEleFengniaoMerchantContractCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_merchant_contract_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}

package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟商户解约接口 APIResponse
alibaba.ele.fengniao.merchant.contract.cancel

通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
*/
type AlibabaEleFengniaoMerchantContractCancelAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoMerchantContractCancelResponse `json:"alibaba_ele_fengniao_merchant_contract_cancel_response,omitempty"`
}

type AlibabaEleFengniaoMerchantContractCancelResponse struct {

    // 出参
    Message   string `json:"message,omitempty"`

}

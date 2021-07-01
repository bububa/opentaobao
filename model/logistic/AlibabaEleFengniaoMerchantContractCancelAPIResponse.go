package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟商户解约接口 API返回值 
alibaba.ele.fengniao.merchant.contract.cancel

通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
*/
type AlibabaEleFengniaoMerchantContractCancelAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoMerchantContractCancelAPIResponseModel
}

// 蜂鸟商户解约接口 成功返回结果
type AlibabaEleFengniaoMerchantContractCancelAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_merchant_contract_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}

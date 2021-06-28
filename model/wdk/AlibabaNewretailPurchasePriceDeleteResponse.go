package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存 商户删除采购价 APIResponse
alibaba.newretail.purchase.price.delete

共享库存 商户删除采购价
*/
type AlibabaNewretailPurchasePriceDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaNewretailPurchasePriceDeleteResponse
}

type AlibabaNewretailPurchasePriceDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_newretail_purchase_price_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 拆单结果对象
    
    Result   *TopBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

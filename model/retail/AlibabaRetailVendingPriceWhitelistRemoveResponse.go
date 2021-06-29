package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格管控白名单去除 APIResponse
alibaba.retail.vending.price.whitelist.remove

商家价格管控白名单去除
*/
type AlibabaRetailVendingPriceWhitelistRemoveAPIResponse struct {
    model.CommonResponse
    AlibabaRetailVendingPriceWhitelistRemoveResponse
}

type AlibabaRetailVendingPriceWhitelistRemoveResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_vending_price_whitelist_remove_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaRetailVendingPriceWhitelistRemoveResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}

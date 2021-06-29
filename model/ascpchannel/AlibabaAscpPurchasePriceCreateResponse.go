package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ascp采购价写入接口 APIResponse
alibaba.ascp.purchase.price.create

供应链平台采购价创建或修改接口
*/
type AlibabaAscpPurchasePriceCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpPurchasePriceCreateResponse
}

type AlibabaAscpPurchasePriceCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_purchase_price_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`

    
}

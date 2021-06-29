package pur

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步产品 APIResponse
alibaba.pur.product.sync

同步产品
*/
type AlibabaPurProductSyncAPIResponse struct {
    model.CommonResponse
    AlibabaPurProductSyncResponse
}

type AlibabaPurProductSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_pur_product_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取url的出参
    
    Result   *ActionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

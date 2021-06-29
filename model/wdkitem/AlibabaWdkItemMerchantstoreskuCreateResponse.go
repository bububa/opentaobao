package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品信息新建 APIResponse
alibaba.wdk.item.merchantstoresku.create

门店商品信息新建
*/
type AlibabaWdkItemMerchantstoreskuCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemMerchantstoreskuCreateResponse
}

type AlibabaWdkItemMerchantstoreskuCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_merchantstoresku_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaWdkItemMerchantstoreskuCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销涉及到的商品的库存同步接口 APIResponse
alibaba.wdk.bm.stock.publish

用于操作sku的库存
*/
type AlibabaWdkBmStockPublishAPIResponse struct {
    model.CommonResponse
    AlibabaWdkBmStockPublishResponse
}

type AlibabaWdkBmStockPublishResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_bm_stock_publish_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *BmResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

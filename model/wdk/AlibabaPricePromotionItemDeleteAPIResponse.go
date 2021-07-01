package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除档期 API返回值 
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品
*/
type AlibabaPricePromotionItemDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaPricePromotionItemDeleteAPIResponseModel
}

// 批量删除档期 成功返回结果
type AlibabaPricePromotionItemDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_price_promotion_item_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaPricePromotionItemDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

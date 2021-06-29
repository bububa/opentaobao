package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品优惠价格查询 APIResponse
tmall.popupstore.item.discount.price

商品优惠价格查询
*/
type TmallPopupstoreItemDiscountPriceAPIResponse struct {
    model.CommonResponse
    TmallPopupstoreItemDiscountPriceResponse
}

type TmallPopupstoreItemDiscountPriceResponse struct {
    XMLName xml.Name `xml:"tmall_popupstore_item_discount_price_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参对象
    
    ResultDto   *TmallPopupstoreItemDiscountPriceResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`

    
}

package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
限时打折详情查询 APIResponse
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。
*/
type TaobaoPromotionLimitdiscountDetailGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_limitdiscount_detail_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 限时打折对应的商品详情列表。
    
    ItemDiscountDetailList   []LimitDiscountDetail `json:"item_discount_detail_list,omitempty" xml:"
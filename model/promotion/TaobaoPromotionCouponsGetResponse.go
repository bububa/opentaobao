package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家优惠券 APIResponse
taobao.promotion.coupons.get

查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量
*/
type TaobaoPromotionCouponsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_coupons_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询的总数量
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"
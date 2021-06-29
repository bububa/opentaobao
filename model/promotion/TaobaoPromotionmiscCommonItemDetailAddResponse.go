package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建通用单品优惠详情 API返回值 
taobao.promotionmisc.common.item.detail.add

创建通用单品优惠详情。
1、使用此接口在指定的优惠活动下创建参与的商品的优惠信息，如还未创建活动，需要先使用接口taobao.promotionmisc.common.item.activity.add创建优惠活动；
2、同一卖家同一活动下的优惠详情数量限制为150个，超过限制需先调用taobao.promotionmisc.common.item.detail.delete接口删除无用的详情后才可再创建新的优惠详情；
3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能创建
*/
type TaobaoPromotionmiscCommonItemDetailAddAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscCommonItemDetailAddResponse
}

// 创建通用单品优惠详情 成功返回结果
type TaobaoPromotionmiscCommonItemDetailAddResponse struct {
    XMLName xml.Name `xml:"promotionmisc_common_item_detail_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否创建成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 优惠详情ID
    DetailId   int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
}

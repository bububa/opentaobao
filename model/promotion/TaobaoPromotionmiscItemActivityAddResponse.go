package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建无条件单品优惠活动 APIResponse
taobao.promotionmisc.item.activity.add

创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。<br/>2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。<br/>3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
type TaobaoPromotionmiscItemActivityAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscItemActivityAddResponse `json:"taobao_promotionmisc_item_activity_add_response,omitempty"`
}

type TaobaoPromotionmiscItemActivityAddResponse struct {

    // 是否保存成功。
    IsSuccess   bool `json:"is_success,omitempty"`

    // 活动id。
    ActivityId   int64 `json:"activity_id,omitempty"`

}

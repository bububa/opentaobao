package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除通用单品优惠详情 APIResponse
taobao.promotionmisc.common.item.detail.delete

删除通用单品优惠详情。
*/
type TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscCommonItemDetailDeleteResponse `json:"taobao_promotionmisc_common_item_detail_delete_response,omitempty"`
}

type TaobaoPromotionmiscCommonItemDetailDeleteResponse struct {

    // 是否删除成功
    IsSuccess   bool `json:"is_success,omitempty"`

}

package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改通用单品优惠活动 APIResponse
taobao.promotionmisc.common.item.activity.update

修改通用单品优惠活动。
1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、使用该接口时需要把未做修改的字段值也传入
*/
type TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscCommonItemActivityUpdateResponse `json:"promotionmisc_common_item_activity_update_response,omitempty"` 
    TaobaoPromotionmiscCommonItemActivityUpdateResponse
}

/* model for simplify = false
type TaobaoPromotionmiscCommonItemActivityUpdateResponse struct {

    // 是否修改成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoPromotionmiscCommonItemActivityUpdateResponse struct {

    // 是否修改成功
    IsSuccess   bool `json:"is_success,omitempty"`

}

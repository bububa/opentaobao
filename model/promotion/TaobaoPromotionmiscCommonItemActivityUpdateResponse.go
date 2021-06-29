package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改通用单品优惠活动 API返回值 
taobao.promotionmisc.common.item.activity.update

修改通用单品优惠活动。
1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、使用该接口时需要把未做修改的字段值也传入
*/
type TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscCommonItemActivityUpdateResponse
}

// 修改通用单品优惠活动 成功返回结果
type TaobaoPromotionmiscCommonItemActivityUpdateResponse struct {
    XMLName xml.Name `xml:"promotionmisc_common_item_activity_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否修改成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityUpdateAPIResponse 修改无条件单品优惠活动 API返回值
// taobao.promotionmisc.item.activity.update
//
// 修改无条件单品优惠活动。&lt;br/&gt;1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 &lt;br/&gt;2、使用该接口时需要同时把未做修改的字段值也传入。 &lt;br/&gt;&lt;br/&gt;3、该接口受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
type TaobaoPromotionmiscItemActivityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscItemActivityUpdateAPIResponseModel
}

// TaobaoPromotionmiscItemActivityUpdateAPIResponseModel is 修改无条件单品优惠活动 成功返回结果
type TaobaoPromotionmiscItemActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_item_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

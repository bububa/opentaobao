package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest
修改通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.update

修改通用单品优惠活动。
1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
2、使用该接口时需要把未做修改的字段值也传入 */
type TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 活动名称，不能超过32字符
	_name string
	// 活动描述，不能超过100字符
	_description string
	// 活动开始时间
	_startTime string
	// 活动结束时间
	_endTime string
	// 是否指定人群标签
	_isUserTag bool
	// 用户标签。当is_user_tag为true时，该值才有意义。
	_userTag string
}

// New

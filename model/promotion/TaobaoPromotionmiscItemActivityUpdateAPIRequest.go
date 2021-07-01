package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscItemActivityUpdateAPIRequest
修改无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.update

修改无条件单品优惠活动。<br/>1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 <br/>2、使用该接口时需要同时把未做修改的字段值也传入。 <br/><br/>3、该接口受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。 */
type TaobaoPromotionmiscItemActivityUpdateAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
	// 活动名称。
	_name string
	// 活动范围：0表示全部参与； 1表示部分商品参与。
	_participateRange int64
	// 活动开始时间。
	_startTime string
	// 活动结束时间。
	_endTime string
	// 是否指定用户标签。
	_isUserTag bool
	// 用户标签。当is_user_tag为true时，该值才有意义。
	_userTag string
	// 是否有减钱行为。
	_isDecreaseMoney bool
	// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
	_decreaseAmount int64
	// 是否有打折行为。
	_isDiscount bool
	// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
	_discountRate int64
}

// New

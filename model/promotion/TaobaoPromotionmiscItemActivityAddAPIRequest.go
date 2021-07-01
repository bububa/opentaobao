package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscItemActivityAddAPIRequest
创建无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.add

创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。<br/>2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。<br/>3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。 */
type TaobaoPromotionmiscItemActivityAddAPIRequest struct {
	model.Params
	// 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
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

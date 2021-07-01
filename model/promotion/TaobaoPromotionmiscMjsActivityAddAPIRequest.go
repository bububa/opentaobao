package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscMjsActivityAddAPIRequest
创建满就送活动 API请求
taobao.promotionmisc.mjs.activity.add

创建满就送活动。<br/>1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。 2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。 3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。 */
type TaobaoPromotionmiscMjsActivityAddAPIRequest struct {
	model.Params
	// 活动名称。
	_name string
	// 活动类型： 1表示商品级别的活动；2表示店铺级别的活动
	_type int64
	// 活动范围：0表示全部参与； 1表示部分商品参与。
	_participateRange int64
	// 活动开始时间。
	_startTime string
	// 活动结束时间。
	_endTime string
	// 是否有满元条件。
	_isAmountOver bool
	// 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
	_totalPrice int64
	// 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。默认为false。
	_isAmountMultiple bool
	// 是否有满件条件。
	_isItemCountOver bool
	// 满多少件。当is_item_count_over为true时，该值才有意义。
	_itemCount int64
	// 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。 默认为false。
	_isItemMultiple bool
	// 是否有店铺会员等级条件。
	_isShopMember bool
	// 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
	_shopMemberLevel int64
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
	// 是否有送礼品行为。
	_isSendGift bool
	// 礼品名称。当is_send_gift为true时，该值才有意义。
	_giftName string
	// 礼品id，当is_send_gift为true时，该值才有意义。 1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
	_giftId int64
	// 商品详情的url，当is_send_gift为true时，该值才有效。
	_giftUrl string
	// 是否有免邮行为。
	_isFreePost bool
	// 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
	_excludeArea string
}

// New

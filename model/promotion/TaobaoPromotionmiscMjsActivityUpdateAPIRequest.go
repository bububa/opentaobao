package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityUpdateAPIRequest 修改满就送活动 API请求
// taobao.promotionmisc.mjs.activity.update
//
// 修改满就送活动。 <br/>1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 <br/>2、使用该接口时需要同时把未做修改的字段值也传入。 <br/>3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
type TaobaoPromotionmiscMjsActivityUpdateAPIRequest struct {
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

// NewTaobaoPromotionmiscMjsActivityUpdateRequest 初始化TaobaoPromotionmiscMjsActivityUpdateAPIRequest对象
func NewTaobaoPromotionmiscMjsActivityUpdateRequest() *TaobaoPromotionmiscMjsActivityUpdateAPIRequest {
	return &TaobaoPromotionmiscMjsActivityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// Set is Name Setter
// 活动名称。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetName() string {
	return r._name
}

// Set is ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetParticipateRange(_participateRange int64) error {
	r._participateRange = _participateRange
	r.Set("participate_range", _participateRange)
	return nil
}

// Get ParticipateRange Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetParticipateRange() int64 {
	return r._participateRange
}

// Set is StartTime Setter
// 活动开始时间。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 活动结束时间。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is IsAmountOver Setter
// 是否有满元条件。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsAmountOver(_isAmountOver bool) error {
	r._isAmountOver = _isAmountOver
	r.Set("is_amount_over", _isAmountOver)
	return nil
}

// Get IsAmountOver Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsAmountOver() bool {
	return r._isAmountOver
}

// Set is TotalPrice Setter
// 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetTotalPrice(_totalPrice int64) error {
	r._totalPrice = _totalPrice
	r.Set("total_price", _totalPrice)
	return nil
}

// Get TotalPrice Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetTotalPrice() int64 {
	return r._totalPrice
}

// Set is IsAmountMultiple Setter
// 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。默认为false。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsAmountMultiple(_isAmountMultiple bool) error {
	r._isAmountMultiple = _isAmountMultiple
	r.Set("is_amount_multiple", _isAmountMultiple)
	return nil
}

// Get IsAmountMultiple Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsAmountMultiple() bool {
	return r._isAmountMultiple
}

// Set is IsItemCountOver Setter
// 是否有满件条件。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsItemCountOver(_isItemCountOver bool) error {
	r._isItemCountOver = _isItemCountOver
	r.Set("is_item_count_over", _isItemCountOver)
	return nil
}

// Get IsItemCountOver Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsItemCountOver() bool {
	return r._isItemCountOver
}

// Set is ItemCount Setter
// 满多少件。当is_item_count_over为true时，该值才有意义。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetItemCount(_itemCount int64) error {
	r._itemCount = _itemCount
	r.Set("item_count", _itemCount)
	return nil
}

// Get ItemCount Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetItemCount() int64 {
	return r._itemCount
}

// Set is IsItemMultiple Setter
// 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。 默认为false。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsItemMultiple(_isItemMultiple bool) error {
	r._isItemMultiple = _isItemMultiple
	r.Set("is_item_multiple", _isItemMultiple)
	return nil
}

// Get IsItemMultiple Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsItemMultiple() bool {
	return r._isItemMultiple
}

// Set is IsShopMember Setter
// 是否有店铺会员等级条件。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsShopMember(_isShopMember bool) error {
	r._isShopMember = _isShopMember
	r.Set("is_shop_member", _isShopMember)
	return nil
}

// Get IsShopMember Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsShopMember() bool {
	return r._isShopMember
}

// Set is ShopMemberLevel Setter
// 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetShopMemberLevel(_shopMemberLevel int64) error {
	r._shopMemberLevel = _shopMemberLevel
	r.Set("shop_member_level", _shopMemberLevel)
	return nil
}

// Get ShopMemberLevel Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetShopMemberLevel() int64 {
	return r._shopMemberLevel
}

// Set is IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsUserTag(_isUserTag bool) error {
	r._isUserTag = _isUserTag
	r.Set("is_user_tag", _isUserTag)
	return nil
}

// Get IsUserTag Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsUserTag() bool {
	return r._isUserTag
}

// Set is UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetUserTag(_userTag string) error {
	r._userTag = _userTag
	r.Set("user_tag", _userTag)
	return nil
}

// Get UserTag Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetUserTag() string {
	return r._userTag
}

// Set is IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsDecreaseMoney(_isDecreaseMoney bool) error {
	r._isDecreaseMoney = _isDecreaseMoney
	r.Set("is_decrease_money", _isDecreaseMoney)
	return nil
}

// Get IsDecreaseMoney Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsDecreaseMoney() bool {
	return r._isDecreaseMoney
}

// Set is DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetDecreaseAmount(_decreaseAmount int64) error {
	r._decreaseAmount = _decreaseAmount
	r.Set("decrease_amount", _decreaseAmount)
	return nil
}

// Get DecreaseAmount Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetDecreaseAmount() int64 {
	return r._decreaseAmount
}

// Set is IsDiscount Setter
// 是否有打折行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsDiscount(_isDiscount bool) error {
	r._isDiscount = _isDiscount
	r.Set("is_discount", _isDiscount)
	return nil
}

// Get IsDiscount Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsDiscount() bool {
	return r._isDiscount
}

// Set is DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetDiscountRate(_discountRate int64) error {
	r._discountRate = _discountRate
	r.Set("discount_rate", _discountRate)
	return nil
}

// Get DiscountRate Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetDiscountRate() int64 {
	return r._discountRate
}

// Set is IsSendGift Setter
// 是否有送礼品行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsSendGift(_isSendGift bool) error {
	r._isSendGift = _isSendGift
	r.Set("is_send_gift", _isSendGift)
	return nil
}

// Get IsSendGift Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsSendGift() bool {
	return r._isSendGift
}

// Set is GiftName Setter
// 礼品名称。当is_send_gift为true时，该值才有意义。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetGiftName(_giftName string) error {
	r._giftName = _giftName
	r.Set("gift_name", _giftName)
	return nil
}

// Get GiftName Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetGiftName() string {
	return r._giftName
}

// Set is GiftId Setter
// 礼品id，当is_send_gift为true时，该值才有意义。 1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetGiftId(_giftId int64) error {
	r._giftId = _giftId
	r.Set("gift_id", _giftId)
	return nil
}

// Get GiftId Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetGiftId() int64 {
	return r._giftId
}

// Set is GiftUrl Setter
// 商品详情的url，当is_send_gift为true时，该值才有效。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetGiftUrl(_giftUrl string) error {
	r._giftUrl = _giftUrl
	r.Set("gift_url", _giftUrl)
	return nil
}

// Get GiftUrl Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetGiftUrl() string {
	return r._giftUrl
}

// Set is IsFreePost Setter
// 是否有免邮行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetIsFreePost(_isFreePost bool) error {
	r._isFreePost = _isFreePost
	r.Set("is_free_post", _isFreePost)
	return nil
}

// Get IsFreePost Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetIsFreePost() bool {
	return r._isFreePost
}

// Set is ExcludeArea Setter
// 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
func (r *TaobaoPromotionmiscMjsActivityUpdateAPIRequest) SetExcludeArea(_excludeArea string) error {
	r._excludeArea = _excludeArea
	r.Set("exclude_area", _excludeArea)
	return nil
}

// Get ExcludeArea Getter
func (r TaobaoPromotionmiscMjsActivityUpdateAPIRequest) GetExcludeArea() string {
	return r._excludeArea
}

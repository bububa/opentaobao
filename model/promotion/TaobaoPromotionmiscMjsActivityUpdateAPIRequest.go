package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscmjsactivityupdateAPIRequest 修改满就送活动 API请求
// taobao.promotionmisc.mjs.activity.update
//
// 修改满就送活动。 &lt;br/&gt;1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 &lt;br/&gt;2、使用该接口时需要同时把未做修改的字段值也传入。 &lt;br/&gt;3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
type TaobaopromotionmiscmjsactivityupdateAPIRequest struct {
	model.Params
	// 活动名称。
	_name string
	// 活动开始时间。
	_startTime string
	// 活动结束时间。
	_endTime string
	// 用户标签。当is_user_tag为true时，该值才有意义。
	_userTag string
	// 礼品名称。当is_send_gift为true时，该值才有意义。
	_giftName string
	// 商品详情的url，当is_send_gift为true时，该值才有效。
	_giftUrl string
	// 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
	_excludeArea string
	// 活动id。
	_activityId int64
	// 活动范围：0表示全部参与； 1表示部分商品参与。
	_participateRange int64
	// 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
	_totalPrice int64
	// 满多少件。当is_item_count_over为true时，该值才有意义。
	_itemCount int64
	// 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
	_shopMemberLevel int64
	// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
	_decreaseAmount int64
	// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
	_discountRate int64
	// 礼品id，当is_send_gift为true时，该值才有意义。 1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
	_giftId int64
	// 是否有满元条件。
	_isAmountOver bool
	// 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。默认为false。
	_isAmountMultiple bool
	// 是否有满件条件。
	_isItemCountOver bool
	// 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。 默认为false。
	_isItemMultiple bool
	// 是否有店铺会员等级条件。
	_isShopMember bool
	// 是否指定用户标签。
	_isUserTag bool
	// 是否有减钱行为。
	_isDecreaseMoney bool
	// 是否有打折行为。
	_isDiscount bool
	// 是否有送礼品行为。
	_isSendGift bool
	// 是否有免邮行为。
	_isFreePost bool
}

// NewTaobaopromotionmiscmjsactivityupdateRequest 初始化TaobaopromotionmiscmjsactivityupdateAPIRequest对象
func NewTaobaopromotionmiscmjsactivityupdateRequest() *TaobaopromotionmiscmjsactivityupdateAPIRequest {
	return &TaobaopromotionmiscmjsactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 活动名称。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetName() string {
	return r._name
}

// SetStartTime is StartTime Setter
// 活动开始时间。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetUserTag is UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetUserTag(_userTag string) error {
	r._userTag = _userTag
	r.Set("user_tag", _userTag)
	return nil
}

// GetUserTag UserTag Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetUserTag() string {
	return r._userTag
}

// SetGiftName is GiftName Setter
// 礼品名称。当is_send_gift为true时，该值才有意义。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetGiftName(_giftName string) error {
	r._giftName = _giftName
	r.Set("gift_name", _giftName)
	return nil
}

// GetGiftName GiftName Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetGiftName() string {
	return r._giftName
}

// SetGiftUrl is GiftUrl Setter
// 商品详情的url，当is_send_gift为true时，该值才有效。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetGiftUrl(_giftUrl string) error {
	r._giftUrl = _giftUrl
	r.Set("gift_url", _giftUrl)
	return nil
}

// GetGiftUrl GiftUrl Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetGiftUrl() string {
	return r._giftUrl
}

// SetExcludeArea is ExcludeArea Setter
// 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetExcludeArea(_excludeArea string) error {
	r._excludeArea = _excludeArea
	r.Set("exclude_area", _excludeArea)
	return nil
}

// GetExcludeArea ExcludeArea Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetExcludeArea() string {
	return r._excludeArea
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetParticipateRange is ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetParticipateRange(_participateRange int64) error {
	r._participateRange = _participateRange
	r.Set("participate_range", _participateRange)
	return nil
}

// GetParticipateRange ParticipateRange Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetParticipateRange() int64 {
	return r._participateRange
}

// SetTotalPrice is TotalPrice Setter
// 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetTotalPrice(_totalPrice int64) error {
	r._totalPrice = _totalPrice
	r.Set("total_price", _totalPrice)
	return nil
}

// GetTotalPrice TotalPrice Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetTotalPrice() int64 {
	return r._totalPrice
}

// SetItemCount is ItemCount Setter
// 满多少件。当is_item_count_over为true时，该值才有意义。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetItemCount(_itemCount int64) error {
	r._itemCount = _itemCount
	r.Set("item_count", _itemCount)
	return nil
}

// GetItemCount ItemCount Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetItemCount() int64 {
	return r._itemCount
}

// SetShopMemberLevel is ShopMemberLevel Setter
// 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetShopMemberLevel(_shopMemberLevel int64) error {
	r._shopMemberLevel = _shopMemberLevel
	r.Set("shop_member_level", _shopMemberLevel)
	return nil
}

// GetShopMemberLevel ShopMemberLevel Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetShopMemberLevel() int64 {
	return r._shopMemberLevel
}

// SetDecreaseAmount is DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetDecreaseAmount(_decreaseAmount int64) error {
	r._decreaseAmount = _decreaseAmount
	r.Set("decrease_amount", _decreaseAmount)
	return nil
}

// GetDecreaseAmount DecreaseAmount Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetDecreaseAmount() int64 {
	return r._decreaseAmount
}

// SetDiscountRate is DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetDiscountRate(_discountRate int64) error {
	r._discountRate = _discountRate
	r.Set("discount_rate", _discountRate)
	return nil
}

// GetDiscountRate DiscountRate Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetDiscountRate() int64 {
	return r._discountRate
}

// SetGiftId is GiftId Setter
// 礼品id，当is_send_gift为true时，该值才有意义。 1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetGiftId(_giftId int64) error {
	r._giftId = _giftId
	r.Set("gift_id", _giftId)
	return nil
}

// GetGiftId GiftId Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetGiftId() int64 {
	return r._giftId
}

// SetIsAmountOver is IsAmountOver Setter
// 是否有满元条件。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsAmountOver(_isAmountOver bool) error {
	r._isAmountOver = _isAmountOver
	r.Set("is_amount_over", _isAmountOver)
	return nil
}

// GetIsAmountOver IsAmountOver Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsAmountOver() bool {
	return r._isAmountOver
}

// SetIsAmountMultiple is IsAmountMultiple Setter
// 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。默认为false。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsAmountMultiple(_isAmountMultiple bool) error {
	r._isAmountMultiple = _isAmountMultiple
	r.Set("is_amount_multiple", _isAmountMultiple)
	return nil
}

// GetIsAmountMultiple IsAmountMultiple Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsAmountMultiple() bool {
	return r._isAmountMultiple
}

// SetIsItemCountOver is IsItemCountOver Setter
// 是否有满件条件。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsItemCountOver(_isItemCountOver bool) error {
	r._isItemCountOver = _isItemCountOver
	r.Set("is_item_count_over", _isItemCountOver)
	return nil
}

// GetIsItemCountOver IsItemCountOver Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsItemCountOver() bool {
	return r._isItemCountOver
}

// SetIsItemMultiple is IsItemMultiple Setter
// 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。 默认为false。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsItemMultiple(_isItemMultiple bool) error {
	r._isItemMultiple = _isItemMultiple
	r.Set("is_item_multiple", _isItemMultiple)
	return nil
}

// GetIsItemMultiple IsItemMultiple Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsItemMultiple() bool {
	return r._isItemMultiple
}

// SetIsShopMember is IsShopMember Setter
// 是否有店铺会员等级条件。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsShopMember(_isShopMember bool) error {
	r._isShopMember = _isShopMember
	r.Set("is_shop_member", _isShopMember)
	return nil
}

// GetIsShopMember IsShopMember Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsShopMember() bool {
	return r._isShopMember
}

// SetIsUserTag is IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsUserTag(_isUserTag bool) error {
	r._isUserTag = _isUserTag
	r.Set("is_user_tag", _isUserTag)
	return nil
}

// GetIsUserTag IsUserTag Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsUserTag() bool {
	return r._isUserTag
}

// SetIsDecreaseMoney is IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsDecreaseMoney(_isDecreaseMoney bool) error {
	r._isDecreaseMoney = _isDecreaseMoney
	r.Set("is_decrease_money", _isDecreaseMoney)
	return nil
}

// GetIsDecreaseMoney IsDecreaseMoney Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsDecreaseMoney() bool {
	return r._isDecreaseMoney
}

// SetIsDiscount is IsDiscount Setter
// 是否有打折行为。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsDiscount(_isDiscount bool) error {
	r._isDiscount = _isDiscount
	r.Set("is_discount", _isDiscount)
	return nil
}

// GetIsDiscount IsDiscount Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsDiscount() bool {
	return r._isDiscount
}

// SetIsSendGift is IsSendGift Setter
// 是否有送礼品行为。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsSendGift(_isSendGift bool) error {
	r._isSendGift = _isSendGift
	r.Set("is_send_gift", _isSendGift)
	return nil
}

// GetIsSendGift IsSendGift Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsSendGift() bool {
	return r._isSendGift
}

// SetIsFreePost is IsFreePost Setter
// 是否有免邮行为。
func (r *TaobaopromotionmiscmjsactivityupdateAPIRequest) SetIsFreePost(_isFreePost bool) error {
	r._isFreePost = _isFreePost
	r.Set("is_free_post", _isFreePost)
	return nil
}

// GetIsFreePost IsFreePost Getter
func (r TaobaopromotionmiscmjsactivityupdateAPIRequest) GetIsFreePost() bool {
	return r._isFreePost
}

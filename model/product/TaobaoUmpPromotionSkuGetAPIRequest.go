package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpPromotionSkuGetAPIRequest 商品优惠详情查询 API请求
// taobao.ump.promotion.sku.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TaobaoUmpPromotionSkuGetAPIRequest struct {
	model.Params
	// skuId
	_skuList string
	// 渠道
	_channelKey string
	// 买家id
	_buyerId string
	// 商品id
	_itemId int64
}

// NewTaobaoUmpPromotionSkuGetRequest 初始化TaobaoUmpPromotionSkuGetAPIRequest对象
func NewTaobaoUmpPromotionSkuGetRequest() *TaobaoUmpPromotionSkuGetAPIRequest {
	return &TaobaoUmpPromotionSkuGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSkuList is SkuList Setter
// skuId
func (r *TaobaoUmpPromotionSkuGetAPIRequest) SetSkuList(_skuList string) error {
	r._skuList = _skuList
	r.Set("sku_list", _skuList)
	return nil
}

// GetSkuList SkuList Getter
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetSkuList() string {
	return r._skuList
}

// SetChannelKey is ChannelKey Setter
// 渠道
func (r *TaobaoUmpPromotionSkuGetAPIRequest) SetChannelKey(_channelKey string) error {
	r._channelKey = _channelKey
	r.Set("channel_key", _channelKey)
	return nil
}

// GetChannelKey ChannelKey Getter
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetChannelKey() string {
	return r._channelKey
}

// SetBuyerId is BuyerId Setter
// 买家id
func (r *TaobaoUmpPromotionSkuGetAPIRequest) SetBuyerId(_buyerId string) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetBuyerId() string {
	return r._buyerId
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoUmpPromotionSkuGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

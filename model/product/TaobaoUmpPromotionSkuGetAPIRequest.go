package product

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpPromotionSkuGetAPIRequest) Reset() {
	r._skuList = ""
	r._channelKey = ""
	r._buyerId = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpPromotionSkuGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoUmpPromotionSkuGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpPromotionSkuGetRequest()
	},
}

// GetTaobaoUmpPromotionSkuGetRequest 从 sync.Pool 获取 TaobaoUmpPromotionSkuGetAPIRequest
func GetTaobaoUmpPromotionSkuGetAPIRequest() *TaobaoUmpPromotionSkuGetAPIRequest {
	return poolTaobaoUmpPromotionSkuGetAPIRequest.Get().(*TaobaoUmpPromotionSkuGetAPIRequest)
}

// ReleaseTaobaoUmpPromotionSkuGetAPIRequest 将 TaobaoUmpPromotionSkuGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpPromotionSkuGetAPIRequest(v *TaobaoUmpPromotionSkuGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpPromotionSkuGetAPIRequest.Put(v)
}

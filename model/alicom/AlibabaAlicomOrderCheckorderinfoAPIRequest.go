package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomordercheckorderinfoAPIRequest 金融购机订单信息校验 API请求
// alibaba.alicom.order.checkorderinfo
//
// 金融购机订单信息校验
type AlibabaalicomordercheckorderinfoAPIRequest struct {
	model.Params
	// 商家的店铺名
	_shopName string
	// 买家昵称
	_userNick string
	// 订单状态1: 等待买家付款   2:等待卖家发货   4:已退款  6: 交易成功
	_tradeStatus int64
	// 淘宝交易订单id
	_bizOrderId int64
	// 宝贝id
	_itemId int64
}

// NewAlibabaalicomordercheckorderinfoRequest 初始化AlibabaalicomordercheckorderinfoAPIRequest对象
func NewAlibabaalicomordercheckorderinfoRequest() *AlibabaalicomordercheckorderinfoAPIRequest {
	return &AlibabaalicomordercheckorderinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.checkorderinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopName is ShopName Setter
// 商家的店铺名
func (r *AlibabaalicomordercheckorderinfoAPIRequest) SetShopName(_shopName string) error {
	r._shopName = _shopName
	r.Set("shop_name", _shopName)
	return nil
}

// GetShopName ShopName Getter
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetShopName() string {
	return r._shopName
}

// SetUserNick is UserNick Setter
// 买家昵称
func (r *AlibabaalicomordercheckorderinfoAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetTradeStatus is TradeStatus Setter
// 订单状态1: 等待买家付款   2:等待卖家发货   4:已退款  6: 交易成功
func (r *AlibabaalicomordercheckorderinfoAPIRequest) SetTradeStatus(_tradeStatus int64) error {
	r._tradeStatus = _tradeStatus
	r.Set("trade_status", _tradeStatus)
	return nil
}

// GetTradeStatus TradeStatus Getter
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetTradeStatus() int64 {
	return r._tradeStatus
}

// SetBizOrderId is BizOrderId Setter
// 淘宝交易订单id
func (r *AlibabaalicomordercheckorderinfoAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetItemId is ItemId Setter
// 宝贝id
func (r *AlibabaalicomordercheckorderinfoAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaalicomordercheckorderinfoAPIRequest) GetItemId() int64 {
	return r._itemId
}

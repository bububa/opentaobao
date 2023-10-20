package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest 1688分销铺货到淘宝关系添加 API请求
// alibaba.fenxiao.cbutotaobao.relation.add
//
// 1688分销铺货到淘宝关系添加
type AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest struct {
	model.Params
	// 淘宝商品id
	_itemId int64
	// 1688买家id
	_buyerId int64
	// 1688供应商id
	_supplierId int64
	// 1688商品id
	_offerId int64
}

// NewAlibabaFenxiaoCbutotaobaoRelationAddRequest 初始化AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest对象
func NewAlibabaFenxiaoCbutotaobaoRelationAddRequest() *AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest {
	return &AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetApiMethodName() string {
	return "alibaba.fenxiao.cbutotaobao.relation.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 淘宝商品id
func (r *AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetBuyerId is BuyerId Setter
// 1688买家id
func (r *AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}

// SetSupplierId is SupplierId Setter
// 1688供应商id
func (r *AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) SetSupplierId(_supplierId int64) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetSupplierId() int64 {
	return r._supplierId
}

// SetOfferId is OfferId Setter
// 1688商品id
func (r *AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlibabaFenxiaoCbutotaobaoRelationAddAPIRequest) GetOfferId() int64 {
	return r._offerId
}

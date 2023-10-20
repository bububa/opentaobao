package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoOrderQrcodeAPIRequest 根据商品id列表获取可扫描下单二维码 API请求
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
type TmallAliautoOrderQrcodeAPIRequest struct {
	model.Params
	// 商品id和数量组合列表
	_itemAndSkuNumList string
	// 门店自定义编码
	_outerShopId string
	// 工单id
	_receiptId int64
}

// NewTmallAliautoOrderQrcodeRequest 初始化TmallAliautoOrderQrcodeAPIRequest对象
func NewTmallAliautoOrderQrcodeRequest() *TmallAliautoOrderQrcodeAPIRequest {
	return &TmallAliautoOrderQrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoOrderQrcodeAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.order.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoOrderQrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoOrderQrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemAndSkuNumList is ItemAndSkuNumList Setter
// 商品id和数量组合列表
func (r *TmallAliautoOrderQrcodeAPIRequest) SetItemAndSkuNumList(_itemAndSkuNumList string) error {
	r._itemAndSkuNumList = _itemAndSkuNumList
	r.Set("item_and_sku_num_list", _itemAndSkuNumList)
	return nil
}

// GetItemAndSkuNumList ItemAndSkuNumList Getter
func (r TmallAliautoOrderQrcodeAPIRequest) GetItemAndSkuNumList() string {
	return r._itemAndSkuNumList
}

// SetOuterShopId is OuterShopId Setter
// 门店自定义编码
func (r *TmallAliautoOrderQrcodeAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r TmallAliautoOrderQrcodeAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

// SetReceiptId is ReceiptId Setter
// 工单id
func (r *TmallAliautoOrderQrcodeAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallAliautoOrderQrcodeAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}

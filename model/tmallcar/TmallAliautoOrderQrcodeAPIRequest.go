package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoorderqrcodeAPIRequest 根据商品id列表获取可扫描下单二维码 API请求
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
type TmallaliautoorderqrcodeAPIRequest struct {
	model.Params
	// 商品id和数量组合列表
	_itemAndSkuNumList string
	// 门店自定义编码
	_outerShopId string
	// 工单id
	_receiptId int64
}

// NewTmallaliautoorderqrcodeRequest 初始化TmallaliautoorderqrcodeAPIRequest对象
func NewTmallaliautoorderqrcodeRequest() *TmallaliautoorderqrcodeAPIRequest {
	return &TmallaliautoorderqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoorderqrcodeAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.order.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoorderqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoorderqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemAndSkuNumList is ItemAndSkuNumList Setter
// 商品id和数量组合列表
func (r *TmallaliautoorderqrcodeAPIRequest) SetItemAndSkuNumList(_itemAndSkuNumList string) error {
	r._itemAndSkuNumList = _itemAndSkuNumList
	r.Set("item_and_sku_num_list", _itemAndSkuNumList)
	return nil
}

// GetItemAndSkuNumList ItemAndSkuNumList Getter
func (r TmallaliautoorderqrcodeAPIRequest) GetItemAndSkuNumList() string {
	return r._itemAndSkuNumList
}

// SetOuterShopId is OuterShopId Setter
// 门店自定义编码
func (r *TmallaliautoorderqrcodeAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r TmallaliautoorderqrcodeAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

// SetReceiptId is ReceiptId Setter
// 工单id
func (r *TmallaliautoorderqrcodeAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallaliautoorderqrcodeAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}

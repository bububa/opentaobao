package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitembarcodeupdateAPIRequest 更新商品条形码信息 API请求
// taobao.item.barcode.update
//
// 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoitembarcodeupdateAPIRequest struct {
	model.Params
	// 商品条形码，如果不用更新，可选择不填
	_itemBarcode string
	// 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
	_skuIds string
	// SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
	_skuBarcodes string
	// 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
	_src string
	// 被更新商品的ID
	_itemId int64
	// 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
	_isforce bool
}

// NewTaobaoitembarcodeupdateRequest 初始化TaobaoitembarcodeupdateAPIRequest对象
func NewTaobaoitembarcodeupdateRequest() *TaobaoitembarcodeupdateAPIRequest {
	return &TaobaoitembarcodeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitembarcodeupdateAPIRequest) GetApiMethodName() string {
	return "taobao.item.barcode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitembarcodeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitembarcodeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemBarcode is ItemBarcode Setter
// 商品条形码，如果不用更新，可选择不填
func (r *TaobaoitembarcodeupdateAPIRequest) SetItemBarcode(_itemBarcode string) error {
	r._itemBarcode = _itemBarcode
	r.Set("item_barcode", _itemBarcode)
	return nil
}

// GetItemBarcode ItemBarcode Getter
func (r TaobaoitembarcodeupdateAPIRequest) GetItemBarcode() string {
	return r._itemBarcode
}

// SetSkuIds is SkuIds Setter
// 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
func (r *TaobaoitembarcodeupdateAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// GetSkuIds SkuIds Getter
func (r TaobaoitembarcodeupdateAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// SetSkuBarcodes is SkuBarcodes Setter
// SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
func (r *TaobaoitembarcodeupdateAPIRequest) SetSkuBarcodes(_skuBarcodes string) error {
	r._skuBarcodes = _skuBarcodes
	r.Set("sku_barcodes", _skuBarcodes)
	return nil
}

// GetSkuBarcodes SkuBarcodes Getter
func (r TaobaoitembarcodeupdateAPIRequest) GetSkuBarcodes() string {
	return r._skuBarcodes
}

// SetSrc is Src Setter
// 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
func (r *TaobaoitembarcodeupdateAPIRequest) SetSrc(_src string) error {
	r._src = _src
	r.Set("src", _src)
	return nil
}

// GetSrc Src Getter
func (r TaobaoitembarcodeupdateAPIRequest) GetSrc() string {
	return r._src
}

// SetItemId is ItemId Setter
// 被更新商品的ID
func (r *TaobaoitembarcodeupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoitembarcodeupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetIsforce is Isforce Setter
// 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
func (r *TaobaoitembarcodeupdateAPIRequest) SetIsforce(_isforce bool) error {
	r._isforce = _isforce
	r.Set("isforce", _isforce)
	return nil
}

// GetIsforce Isforce Getter
func (r TaobaoitembarcodeupdateAPIRequest) GetIsforce() bool {
	return r._isforce
}

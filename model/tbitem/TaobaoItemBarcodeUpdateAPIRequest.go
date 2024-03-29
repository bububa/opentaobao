package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemBarcodeUpdateAPIRequest 更新商品条形码信息 API请求
// taobao.item.barcode.update
//
// 通过该接口，将商品以及SKU上得条形码信息补全
type TaobaoItemBarcodeUpdateAPIRequest struct {
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

// NewTaobaoItemBarcodeUpdateRequest 初始化TaobaoItemBarcodeUpdateAPIRequest对象
func NewTaobaoItemBarcodeUpdateRequest() *TaobaoItemBarcodeUpdateAPIRequest {
	return &TaobaoItemBarcodeUpdateAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemBarcodeUpdateAPIRequest) Reset() {
	r._itemBarcode = ""
	r._skuIds = ""
	r._skuBarcodes = ""
	r._src = ""
	r._itemId = 0
	r._isforce = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemBarcodeUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.item.barcode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemBarcodeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemBarcodeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemBarcode is ItemBarcode Setter
// 商品条形码，如果不用更新，可选择不填
func (r *TaobaoItemBarcodeUpdateAPIRequest) SetItemBarcode(_itemBarcode string) error {
	r._itemBarcode = _itemBarcode
	r.Set("item_barcode", _itemBarcode)
	return nil
}

// GetItemBarcode ItemBarcode Getter
func (r TaobaoItemBarcodeUpdateAPIRequest) GetItemBarcode() string {
	return r._itemBarcode
}

// SetSkuIds is SkuIds Setter
// 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
func (r *TaobaoItemBarcodeUpdateAPIRequest) SetSkuIds(_skuIds string) error {
	r._skuIds = _skuIds
	r.Set("sku_ids", _skuIds)
	return nil
}

// GetSkuIds SkuIds Getter
func (r TaobaoItemBarcodeUpdateAPIRequest) GetSkuIds() string {
	return r._skuIds
}

// SetSkuBarcodes is SkuBarcodes Setter
// SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
func (r *TaobaoItemBarcodeUpdateAPIRequest) SetSkuBarcodes(_skuBarcodes string) error {
	r._skuBarcodes = _skuBarcodes
	r.Set("sku_barcodes", _skuBarcodes)
	return nil
}

// GetSkuBarcodes SkuBarcodes Getter
func (r TaobaoItemBarcodeUpdateAPIRequest) GetSkuBarcodes() string {
	return r._skuBarcodes
}

// SetSrc is Src Setter
// 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
func (r *TaobaoItemBarcodeUpdateAPIRequest) SetSrc(_src string) error {
	r._src = _src
	r.Set("src", _src)
	return nil
}

// GetSrc Src Getter
func (r TaobaoItemBarcodeUpdateAPIRequest) GetSrc() string {
	return r._src
}

// SetItemId is ItemId Setter
// 被更新商品的ID
func (r *TaobaoItemBarcodeUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoItemBarcodeUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetIsforce is Isforce Setter
// 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
func (r *TaobaoItemBarcodeUpdateAPIRequest) SetIsforce(_isforce bool) error {
	r._isforce = _isforce
	r.Set("isforce", _isforce)
	return nil
}

// GetIsforce Isforce Getter
func (r TaobaoItemBarcodeUpdateAPIRequest) GetIsforce() bool {
	return r._isforce
}

var poolTaobaoItemBarcodeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemBarcodeUpdateRequest()
	},
}

// GetTaobaoItemBarcodeUpdateRequest 从 sync.Pool 获取 TaobaoItemBarcodeUpdateAPIRequest
func GetTaobaoItemBarcodeUpdateAPIRequest() *TaobaoItemBarcodeUpdateAPIRequest {
	return poolTaobaoItemBarcodeUpdateAPIRequest.Get().(*TaobaoItemBarcodeUpdateAPIRequest)
}

// ReleaseTaobaoItemBarcodeUpdateAPIRequest 将 TaobaoItemBarcodeUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemBarcodeUpdateAPIRequest(v *TaobaoItemBarcodeUpdateAPIRequest) {
	v.Reset()
	poolTaobaoItemBarcodeUpdateAPIRequest.Put(v)
}

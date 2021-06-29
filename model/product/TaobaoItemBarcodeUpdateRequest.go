package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品条形码信息 API请求
taobao.item.barcode.update

通过该接口，将商品以及SKU上得条形码信息补全
*/
type TaobaoItemBarcodeUpdateRequest struct {
    model.Params
    // 被更新商品的ID
    _itemId   int64
    // 商品条形码，如果不用更新，可选择不填
    _itemBarcode   string
    // 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
    _skuIds   string
    // SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
    _skuBarcodes   string
    // 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
    _isforce   bool
    // 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
    _src   string
}

// 初始化TaobaoItemBarcodeUpdateRequest对象
func NewTaobaoItemBarcodeUpdateRequest() *TaobaoItemBarcodeUpdateRequest{
    return &TaobaoItemBarcodeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemBarcodeUpdateRequest) GetApiMethodName() string {
    return "taobao.item.barcode.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemBarcodeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 被更新商品的ID
func (r *TaobaoItemBarcodeUpdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoItemBarcodeUpdateRequest) GetItemId() int64 {
    return r._itemId
}
// ItemBarcode Setter
// 商品条形码，如果不用更新，可选择不填
func (r *TaobaoItemBarcodeUpdateRequest) SetItemBarcode(_itemBarcode string) error {
    r._itemBarcode = _itemBarcode
    r.Set("item_barcode", _itemBarcode)
    return nil
}

// ItemBarcode Getter
func (r TaobaoItemBarcodeUpdateRequest) GetItemBarcode() string {
    return r._itemBarcode
}
// SkuIds Setter
// 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
func (r *TaobaoItemBarcodeUpdateRequest) SetSkuIds(_skuIds string) error {
    r._skuIds = _skuIds
    r.Set("sku_ids", _skuIds)
    return nil
}

// SkuIds Getter
func (r TaobaoItemBarcodeUpdateRequest) GetSkuIds() string {
    return r._skuIds
}
// SkuBarcodes Setter
// SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
func (r *TaobaoItemBarcodeUpdateRequest) SetSkuBarcodes(_skuBarcodes string) error {
    r._skuBarcodes = _skuBarcodes
    r.Set("sku_barcodes", _skuBarcodes)
    return nil
}

// SkuBarcodes Getter
func (r TaobaoItemBarcodeUpdateRequest) GetSkuBarcodes() string {
    return r._skuBarcodes
}
// Isforce Setter
// 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
func (r *TaobaoItemBarcodeUpdateRequest) SetIsforce(_isforce bool) error {
    r._isforce = _isforce
    r.Set("isforce", _isforce)
    return nil
}

// Isforce Getter
func (r TaobaoItemBarcodeUpdateRequest) GetIsforce() bool {
    return r._isforce
}
// Src Setter
// 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
func (r *TaobaoItemBarcodeUpdateRequest) SetSrc(_src string) error {
    r._src = _src
    r.Set("src", _src)
    return nil
}

// Src Getter
func (r TaobaoItemBarcodeUpdateRequest) GetSrc() string {
    return r._src
}

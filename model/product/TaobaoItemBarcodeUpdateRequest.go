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
    itemId   int64
    // 商品条形码，如果不用更新，可选择不填
    itemBarcode   string
    // 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
    skuIds   string
    // SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
    skuBarcodes   string
    // 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
    isforce   bool
    // 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
    src   string
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
func (r *TaobaoItemBarcodeUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoItemBarcodeUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// ItemBarcode Setter
// 商品条形码，如果不用更新，可选择不填
func (r *TaobaoItemBarcodeUpdateRequest) SetItemBarcode(itemBarcode string) error {
    r.itemBarcode = itemBarcode
    r.Set("item_barcode", itemBarcode)
    return nil
}

// ItemBarcode Getter
func (r TaobaoItemBarcodeUpdateRequest) GetItemBarcode() string {
    return r.itemBarcode
}
// SkuIds Setter
// 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
func (r *TaobaoItemBarcodeUpdateRequest) SetSkuIds(skuIds string) error {
    r.skuIds = skuIds
    r.Set("sku_ids", skuIds)
    return nil
}

// SkuIds Getter
func (r TaobaoItemBarcodeUpdateRequest) GetSkuIds() string {
    return r.skuIds
}
// SkuBarcodes Setter
// SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
func (r *TaobaoItemBarcodeUpdateRequest) SetSkuBarcodes(skuBarcodes string) error {
    r.skuBarcodes = skuBarcodes
    r.Set("sku_barcodes", skuBarcodes)
    return nil
}

// SkuBarcodes Getter
func (r TaobaoItemBarcodeUpdateRequest) GetSkuBarcodes() string {
    return r.skuBarcodes
}
// Isforce Setter
// 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
func (r *TaobaoItemBarcodeUpdateRequest) SetIsforce(isforce bool) error {
    r.isforce = isforce
    r.Set("isforce", isforce)
    return nil
}

// Isforce Getter
func (r TaobaoItemBarcodeUpdateRequest) GetIsforce() bool {
    return r.isforce
}
// Src Setter
// 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
func (r *TaobaoItemBarcodeUpdateRequest) SetSrc(src string) error {
    r.src = src
    r.Set("src", src)
    return nil
}

// Src Getter
func (r TaobaoItemBarcodeUpdateRequest) GetSrc() string {
    return r.src
}

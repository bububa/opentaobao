package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品条形码信息 APIRequest
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

func NewTaobaoItemBarcodeUpdateRequest() *TaobaoItemBarcodeUpdateRequest{
    return &TaobaoItemBarcodeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemBarcodeUpdateRequest) GetApiMethodName() string {
    return "taobao.item.barcode.update"
}

func (r TaobaoItemBarcodeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemBarcodeUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoItemBarcodeUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoItemBarcodeUpdateRequest) SetItemBarcode(itemBarcode string) error {
    r.itemBarcode = itemBarcode
    r.Set("item_barcode", itemBarcode)
    return nil
}

func (r TaobaoItemBarcodeUpdateRequest) GetItemBarcode() string {
    return r.itemBarcode
}

func (r *TaobaoItemBarcodeUpdateRequest) SetSkuIds(skuIds string) error {
    r.skuIds = skuIds
    r.Set("sku_ids", skuIds)
    return nil
}

func (r TaobaoItemBarcodeUpdateRequest) GetSkuIds() string {
    return r.skuIds
}

func (r *TaobaoItemBarcodeUpdateRequest) SetSkuBarcodes(skuBarcodes string) error {
    r.skuBarcodes = skuBarcodes
    r.Set("sku_barcodes", skuBarcodes)
    return nil
}

func (r TaobaoItemBarcodeUpdateRequest) GetSkuBarcodes() string {
    return r.skuBarcodes
}

func (r *TaobaoItemBarcodeUpdateRequest) SetIsforce(isforce bool) error {
    r.isforce = isforce
    r.Set("isforce", isforce)
    return nil
}

func (r TaobaoItemBarcodeUpdateRequest) GetIsforce() bool {
    return r.isforce
}

func (r *TaobaoItemBarcodeUpdateRequest) SetSrc(src string) error {
    r.src = src
    r.Set("src", src)
    return nil
}

func (r TaobaoItemBarcodeUpdateRequest) GetSrc() string {
    return r.src
}


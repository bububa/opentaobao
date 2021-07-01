package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布类目预测 API请求
alibaba.item.category.predict

<font color='red'>商品发布类目预测接口，预测匹配的结果存在一定误差，需要商家二次确认，避免类目配置错误产生其他影响。</font>
*/
type AlibabaItemCategoryPredictAPIRequest struct {
    model.Params
    // 商品主图链接，最多5张，传入完整URL
    _images   []string
    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    _market   string
    // 商品条码
    _barcode   string
    // 商品条码图片
    _barcodeImage   string
    // 商品介绍
    _itemDesc   string
}

// 初始化AlibabaItemCategoryPredictAPIRequest对象
func NewAlibabaItemCategoryPredictRequest() *AlibabaItemCategoryPredictAPIRequest{
    return &AlibabaItemCategoryPredictAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemCategoryPredictAPIRequest) GetApiMethodName() string {
    return "alibaba.item.category.predict"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemCategoryPredictAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Images Setter
// 商品主图链接，最多5张，传入完整URL
func (r *AlibabaItemCategoryPredictAPIRequest) SetImages(_images []string) error {
    r._images = _images
    r.Set("images", _images)
    return nil
}

// Images Getter
func (r AlibabaItemCategoryPredictAPIRequest) GetImages() []string {
    return r._images
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemCategoryPredictAPIRequest) SetMarket(_market string) error {
    r._market = _market
    r.Set("market", _market)
    return nil
}

// Market Getter
func (r AlibabaItemCategoryPredictAPIRequest) GetMarket() string {
    return r._market
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemCategoryPredictAPIRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemCategoryPredictAPIRequest) GetBarcode() string {
    return r._barcode
}
// BarcodeImage Setter
// 商品条码图片
func (r *AlibabaItemCategoryPredictAPIRequest) SetBarcodeImage(_barcodeImage string) error {
    r._barcodeImage = _barcodeImage
    r.Set("barcode_image", _barcodeImage)
    return nil
}

// BarcodeImage Getter
func (r AlibabaItemCategoryPredictAPIRequest) GetBarcodeImage() string {
    return r._barcodeImage
}
// ItemDesc Setter
// 商品介绍
func (r *AlibabaItemCategoryPredictAPIRequest) SetItemDesc(_itemDesc string) error {
    r._itemDesc = _itemDesc
    r.Set("item_desc", _itemDesc)
    return nil
}

// ItemDesc Getter
func (r AlibabaItemCategoryPredictAPIRequest) GetItemDesc() string {
    return r._itemDesc
}

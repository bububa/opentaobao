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
type AlibabaItemCategoryPredictRequest struct {
    model.Params
    // 商品主图链接，最多5张，传入完整URL
    images   []string
    // 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
    market   string
    // 商品条码
    barcode   string
    // 商品条码图片
    barcodeImage   string
    // 商品介绍
    itemDesc   string
}

// 初始化AlibabaItemCategoryPredictRequest对象
func NewAlibabaItemCategoryPredictRequest() *AlibabaItemCategoryPredictRequest{
    return &AlibabaItemCategoryPredictRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemCategoryPredictRequest) GetApiMethodName() string {
    return "alibaba.item.category.predict"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemCategoryPredictRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Images Setter
// 商品主图链接，最多5张，传入完整URL
func (r *AlibabaItemCategoryPredictRequest) SetImages(images []string) error {
    r.images = images
    r.Set("images", images)
    return nil
}

// Images Getter
func (r AlibabaItemCategoryPredictRequest) GetImages() []string {
    return r.images
}
// Market Setter
// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
func (r *AlibabaItemCategoryPredictRequest) SetMarket(market string) error {
    r.market = market
    r.Set("market", market)
    return nil
}

// Market Getter
func (r AlibabaItemCategoryPredictRequest) GetMarket() string {
    return r.market
}
// Barcode Setter
// 商品条码
func (r *AlibabaItemCategoryPredictRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

// Barcode Getter
func (r AlibabaItemCategoryPredictRequest) GetBarcode() string {
    return r.barcode
}
// BarcodeImage Setter
// 商品条码图片
func (r *AlibabaItemCategoryPredictRequest) SetBarcodeImage(barcodeImage string) error {
    r.barcodeImage = barcodeImage
    r.Set("barcode_image", barcodeImage)
    return nil
}

// BarcodeImage Getter
func (r AlibabaItemCategoryPredictRequest) GetBarcodeImage() string {
    return r.barcodeImage
}
// ItemDesc Setter
// 商品介绍
func (r *AlibabaItemCategoryPredictRequest) SetItemDesc(itemDesc string) error {
    r.itemDesc = itemDesc
    r.Set("item_desc", itemDesc)
    return nil
}

// ItemDesc Getter
func (r AlibabaItemCategoryPredictRequest) GetItemDesc() string {
    return r.itemDesc
}

package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心价格修改 API请求
alibaba.wdk.item.price.update

修改门店商品售价和会员价
*/
type AlibabaWdkItemPriceUpdateRequest struct {
    model.Params
    // 盒马门店id
    storeId   string
    // 商品编码
    skuCode   string
    // 价格，单位是分
    skuPrice   int64
    // 会员价格，单位是分，且不能大于价格
    skuMemberPrice   int64
}

// 初始化AlibabaWdkItemPriceUpdateRequest对象
func NewAlibabaWdkItemPriceUpdateRequest() *AlibabaWdkItemPriceUpdateRequest{
    return &AlibabaWdkItemPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemPriceUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.price.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 盒马门店id
func (r *AlibabaWdkItemPriceUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetStoreId() string {
    return r.storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// SkuPrice Setter
// 价格，单位是分
func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuPrice(skuPrice int64) error {
    r.skuPrice = skuPrice
    r.Set("sku_price", skuPrice)
    return nil
}

// SkuPrice Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetSkuPrice() int64 {
    return r.skuPrice
}
// SkuMemberPrice Setter
// 会员价格，单位是分，且不能大于价格
func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuMemberPrice(skuMemberPrice int64) error {
    r.skuMemberPrice = skuMemberPrice
    r.Set("sku_member_price", skuMemberPrice)
    return nil
}

// SkuMemberPrice Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetSkuMemberPrice() int64 {
    return r.skuMemberPrice
}

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
    _storeId   string
    // 商品编码
    _skuCode   string
    // 价格，单位是分
    _skuPrice   int64
    // 会员价格，单位是分，且不能大于价格
    _skuMemberPrice   int64
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
func (r *AlibabaWdkItemPriceUpdateRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetStoreId() string {
    return r._storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetSkuCode() string {
    return r._skuCode
}
// SkuPrice Setter
// 价格，单位是分
func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuPrice(_skuPrice int64) error {
    r._skuPrice = _skuPrice
    r.Set("sku_price", _skuPrice)
    return nil
}

// SkuPrice Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetSkuPrice() int64 {
    return r._skuPrice
}
// SkuMemberPrice Setter
// 会员价格，单位是分，且不能大于价格
func (r *AlibabaWdkItemPriceUpdateRequest) SetSkuMemberPrice(_skuMemberPrice int64) error {
    r._skuMemberPrice = _skuMemberPrice
    r.Set("sku_member_price", _skuMemberPrice)
    return nil
}

// SkuMemberPrice Getter
func (r AlibabaWdkItemPriceUpdateRequest) GetSkuMemberPrice() int64 {
    return r._skuMemberPrice
}

package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品售价会员价修改 API请求
alibaba.wdk.item.memberprice.update

商品售价会员价修改
*/
type AlibabaWdkItemMemberpriceUpdateRequest struct {
    model.Params
    // 门店ID
    storeId   string
    // 商品编码
    skuCode   string
    // 售价，单位分，售价会员价至少填一个
    skuPrice   int64
    // 会员价，单位分
    skuMemberPrice   int64
    // 是否清空会员价
    cleanSkuMemberPrice   bool
    // 时间戳
    timeStamp   int64
}

// 初始化AlibabaWdkItemMemberpriceUpdateRequest对象
func NewAlibabaWdkItemMemberpriceUpdateRequest() *AlibabaWdkItemMemberpriceUpdateRequest{
    return &AlibabaWdkItemMemberpriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.memberprice.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetStoreId() string {
    return r.storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// SkuPrice Setter
// 售价，单位分，售价会员价至少填一个
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetSkuPrice(skuPrice int64) error {
    r.skuPrice = skuPrice
    r.Set("sku_price", skuPrice)
    return nil
}

// SkuPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetSkuPrice() int64 {
    return r.skuPrice
}
// SkuMemberPrice Setter
// 会员价，单位分
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetSkuMemberPrice(skuMemberPrice int64) error {
    r.skuMemberPrice = skuMemberPrice
    r.Set("sku_member_price", skuMemberPrice)
    return nil
}

// SkuMemberPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetSkuMemberPrice() int64 {
    return r.skuMemberPrice
}
// CleanSkuMemberPrice Setter
// 是否清空会员价
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetCleanSkuMemberPrice(cleanSkuMemberPrice bool) error {
    r.cleanSkuMemberPrice = cleanSkuMemberPrice
    r.Set("clean_sku_member_price", cleanSkuMemberPrice)
    return nil
}

// CleanSkuMemberPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetCleanSkuMemberPrice() bool {
    return r.cleanSkuMemberPrice
}
// TimeStamp Setter
// 时间戳
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetTimeStamp(timeStamp int64) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetTimeStamp() int64 {
    return r.timeStamp
}

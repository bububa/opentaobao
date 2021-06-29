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
    _storeId   string
    // 商品编码
    _skuCode   string
    // 售价，单位分，售价会员价至少填一个
    _skuPrice   int64
    // 会员价，单位分
    _skuMemberPrice   int64
    // 是否清空会员价
    _cleanSkuMemberPrice   bool
    // 时间戳
    _timeStamp   int64
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
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetStoreId() string {
    return r._storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetSkuCode() string {
    return r._skuCode
}
// SkuPrice Setter
// 售价，单位分，售价会员价至少填一个
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetSkuPrice(_skuPrice int64) error {
    r._skuPrice = _skuPrice
    r.Set("sku_price", _skuPrice)
    return nil
}

// SkuPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetSkuPrice() int64 {
    return r._skuPrice
}
// SkuMemberPrice Setter
// 会员价，单位分
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetSkuMemberPrice(_skuMemberPrice int64) error {
    r._skuMemberPrice = _skuMemberPrice
    r.Set("sku_member_price", _skuMemberPrice)
    return nil
}

// SkuMemberPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetSkuMemberPrice() int64 {
    return r._skuMemberPrice
}
// CleanSkuMemberPrice Setter
// 是否清空会员价
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetCleanSkuMemberPrice(_cleanSkuMemberPrice bool) error {
    r._cleanSkuMemberPrice = _cleanSkuMemberPrice
    r.Set("clean_sku_member_price", _cleanSkuMemberPrice)
    return nil
}

// CleanSkuMemberPrice Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetCleanSkuMemberPrice() bool {
    return r._cleanSkuMemberPrice
}
// TimeStamp Setter
// 时间戳
func (r *AlibabaWdkItemMemberpriceUpdateRequest) SetTimeStamp(_timeStamp int64) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaWdkItemMemberpriceUpdateRequest) GetTimeStamp() int64 {
    return r._timeStamp
}

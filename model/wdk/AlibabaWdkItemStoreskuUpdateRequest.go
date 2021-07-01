package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心门店商品修改 API请求
alibaba.wdk.item.storesku.update

五道口商品中心门店商品修改
*/
type AlibabaWdkItemStoreskuUpdateAPIRequest struct {
    model.Params
    // 盒马门店id
    _storeId   string
    // 商品编码
    _skuCode   string
    // 1-可售   0-不可售
    _saleFlag   int64
}

// 初始化AlibabaWdkItemStoreskuUpdateAPIRequest对象
func NewAlibabaWdkItemStoreskuUpdateRequest() *AlibabaWdkItemStoreskuUpdateAPIRequest{
    return &AlibabaWdkItemStoreskuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskuUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storesku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 盒马门店id
func (r *AlibabaWdkItemStoreskuUpdateAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemStoreskuUpdateAPIRequest) GetStoreId() string {
    return r._storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemStoreskuUpdateAPIRequest) SetSkuCode(_skuCode string) error {
    r._skuCode = _skuCode
    r.Set("sku_code", _skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemStoreskuUpdateAPIRequest) GetSkuCode() string {
    return r._skuCode
}
// SaleFlag Setter
// 1-可售   0-不可售
func (r *AlibabaWdkItemStoreskuUpdateAPIRequest) SetSaleFlag(_saleFlag int64) error {
    r._saleFlag = _saleFlag
    r.Set("sale_flag", _saleFlag)
    return nil
}

// SaleFlag Getter
func (r AlibabaWdkItemStoreskuUpdateAPIRequest) GetSaleFlag() int64 {
    return r._saleFlag
}

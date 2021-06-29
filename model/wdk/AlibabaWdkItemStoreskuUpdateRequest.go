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
type AlibabaWdkItemStoreskuUpdateRequest struct {
    model.Params
    // 盒马门店id
    storeId   string
    // 商品编码
    skuCode   string
    // 1-可售   0-不可售
    saleFlag   int64
}

// 初始化AlibabaWdkItemStoreskuUpdateRequest对象
func NewAlibabaWdkItemStoreskuUpdateRequest() *AlibabaWdkItemStoreskuUpdateRequest{
    return &AlibabaWdkItemStoreskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storesku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 盒马门店id
func (r *AlibabaWdkItemStoreskuUpdateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkItemStoreskuUpdateRequest) GetStoreId() string {
    return r.storeId
}
// SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemStoreskuUpdateRequest) SetSkuCode(skuCode string) error {
    r.skuCode = skuCode
    r.Set("sku_code", skuCode)
    return nil
}

// SkuCode Getter
func (r AlibabaWdkItemStoreskuUpdateRequest) GetSkuCode() string {
    return r.skuCode
}
// SaleFlag Setter
// 1-可售   0-不可售
func (r *AlibabaWdkItemStoreskuUpdateRequest) SetSaleFlag(saleFlag int64) error {
    r.saleFlag = saleFlag
    r.Set("sale_flag", saleFlag)
    return nil
}

// SaleFlag Getter
func (r AlibabaWdkItemStoreskuUpdateRequest) GetSaleFlag() int64 {
    return r.saleFlag
}

package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量游标方式查询接口 API请求
alibaba.wdk.sku.scroll.query

通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
*/
type AlibabaWdkSkuScrollQueryRequest struct {
    model.Params
    // 商家类目编码
    _merchantCatCode   string
    // 门店编码
    _ouCode   string
    // 游标：第一次请求不用填写，否则请填写上一次请求返回的值，直到获取到足够的数据
    _scrollId   string
    // 英文逗号分隔的商品编码，最多20个。如果配合门店字段使用，直接非游标方式返回商品数据
    _skuCodes   string
}

// 初始化AlibabaWdkSkuScrollQueryRequest对象
func NewAlibabaWdkSkuScrollQueryRequest() *AlibabaWdkSkuScrollQueryRequest{
    return &AlibabaWdkSkuScrollQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuScrollQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.scroll.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuScrollQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCatCode Setter
// 商家类目编码
func (r *AlibabaWdkSkuScrollQueryRequest) SetMerchantCatCode(_merchantCatCode string) error {
    r._merchantCatCode = _merchantCatCode
    r.Set("merchant_cat_code", _merchantCatCode)
    return nil
}

// MerchantCatCode Getter
func (r AlibabaWdkSkuScrollQueryRequest) GetMerchantCatCode() string {
    return r._merchantCatCode
}
// OuCode Setter
// 门店编码
func (r *AlibabaWdkSkuScrollQueryRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaWdkSkuScrollQueryRequest) GetOuCode() string {
    return r._ouCode
}
// ScrollId Setter
// 游标：第一次请求不用填写，否则请填写上一次请求返回的值，直到获取到足够的数据
func (r *AlibabaWdkSkuScrollQueryRequest) SetScrollId(_scrollId string) error {
    r._scrollId = _scrollId
    r.Set("scroll_id", _scrollId)
    return nil
}

// ScrollId Getter
func (r AlibabaWdkSkuScrollQueryRequest) GetScrollId() string {
    return r._scrollId
}
// SkuCodes Setter
// 英文逗号分隔的商品编码，最多20个。如果配合门店字段使用，直接非游标方式返回商品数据
func (r *AlibabaWdkSkuScrollQueryRequest) SetSkuCodes(_skuCodes string) error {
    r._skuCodes = _skuCodes
    r.Set("sku_codes", _skuCodes)
    return nil
}

// SkuCodes Getter
func (r AlibabaWdkSkuScrollQueryRequest) GetSkuCodes() string {
    return r._skuCodes
}

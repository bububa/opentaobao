package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id获取商品属性 API请求
alibaba.alihealth.tracecodeseller.product.attr.search

根据商品id获取商品属性
*/
type AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest struct {
    model.Params
    // 企业id
    _entInfoId   int64
    // 货品id
    _tracUserProductInfoId   int64
}

// 初始化AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerProductAttrSearchRequest() *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest{
    return &AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.product.attr.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
// TracUserProductInfoId Setter
// 货品id
func (r *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) SetTracUserProductInfoId(_tracUserProductInfoId int64) error {
    r._tracUserProductInfoId = _tracUserProductInfoId
    r.Set("trac_user_product_info_id", _tracUserProductInfoId)
    return nil
}

// TracUserProductInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetTracUserProductInfoId() int64 {
    return r._tracUserProductInfoId
}

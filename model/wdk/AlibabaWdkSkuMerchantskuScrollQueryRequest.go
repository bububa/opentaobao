package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品批量查询接口 API请求
alibaba.wdk.sku.merchantsku.scroll.query

提供主档商品数据接口查询
*/
type AlibabaWdkSkuMerchantskuScrollQueryRequest struct {
    model.Params
    // HM
    _orgNo   string
    // 第一次为null，后面从结果中获取
    _scrollId   string
}

// 初始化AlibabaWdkSkuMerchantskuScrollQueryRequest对象
func NewAlibabaWdkSkuMerchantskuScrollQueryRequest() *AlibabaWdkSkuMerchantskuScrollQueryRequest{
    return &AlibabaWdkSkuMerchantskuScrollQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.merchantsku.scroll.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrgNo Setter
// HM
func (r *AlibabaWdkSkuMerchantskuScrollQueryRequest) SetOrgNo(_orgNo string) error {
    r._orgNo = _orgNo
    r.Set("org_no", _orgNo)
    return nil
}

// OrgNo Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetOrgNo() string {
    return r._orgNo
}
// ScrollId Setter
// 第一次为null，后面从结果中获取
func (r *AlibabaWdkSkuMerchantskuScrollQueryRequest) SetScrollId(_scrollId string) error {
    r._scrollId = _scrollId
    r.Set("scroll_id", _scrollId)
    return nil
}

// ScrollId Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetScrollId() string {
    return r._scrollId
}

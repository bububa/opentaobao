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
    orgNo   string
    // 第一次为null，后面从结果中获取
    scrollId   string
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
func (r *AlibabaWdkSkuMerchantskuScrollQueryRequest) SetOrgNo(orgNo string) error {
    r.orgNo = orgNo
    r.Set("org_no", orgNo)
    return nil
}

// OrgNo Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetOrgNo() string {
    return r.orgNo
}
// ScrollId Setter
// 第一次为null，后面从结果中获取
func (r *AlibabaWdkSkuMerchantskuScrollQueryRequest) SetScrollId(scrollId string) error {
    r.scrollId = scrollId
    r.Set("scroll_id", scrollId)
    return nil
}

// ScrollId Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetScrollId() string {
    return r.scrollId
}

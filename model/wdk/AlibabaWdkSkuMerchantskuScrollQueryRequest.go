package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品批量查询接口 APIRequest
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

func NewAlibabaWdkSkuMerchantskuScrollQueryRequest() *AlibabaWdkSkuMerchantskuScrollQueryRequest{
    return &AlibabaWdkSkuMerchantskuScrollQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.merchantsku.scroll.query"
}

func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuMerchantskuScrollQueryRequest) SetOrgNo(orgNo string) error {
    r.orgNo = orgNo
    r.Set("org_no", orgNo)
    return nil
}

func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetOrgNo() string {
    return r.orgNo
}

func (r *AlibabaWdkSkuMerchantskuScrollQueryRequest) SetScrollId(scrollId string) error {
    r.scrollId = scrollId
    r.Set("scroll_id", scrollId)
    return nil
}

func (r AlibabaWdkSkuMerchantskuScrollQueryRequest) GetScrollId() string {
    return r.scrollId
}


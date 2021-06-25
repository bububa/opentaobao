package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建单品特价活动【同城零售】 APIRequest
alibaba.retail.marketing.itemdiscount.activity.create

同城零售单品特价活动创建
*/
type AlibabaRetailMarketingItemdiscountActivityCreateRequest struct {
    model.Params

    // 创建活动参数
    param   *ItemDiscountActivityOperateRequest 

}

func NewAlibabaRetailMarketingItemdiscountActivityCreateRequest() *AlibabaRetailMarketingItemdiscountActivityCreateRequest{
    return &AlibabaRetailMarketingItemdiscountActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItemdiscountActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.create"
}

func (r AlibabaRetailMarketingItemdiscountActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItemdiscountActivityCreateRequest) SetParam(param *ItemDiscountActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItemdiscountActivityCreateRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r.param
}


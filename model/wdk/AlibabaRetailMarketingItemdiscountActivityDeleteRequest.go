package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品特价活动【同城零售】 APIRequest
alibaba.retail.marketing.itemdiscount.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingItemdiscountActivityDeleteRequest struct {
    model.Params

    // 删除活动参数
    param   *ItemDiscountActivityOperateRequest 

}

func NewAlibabaRetailMarketingItemdiscountActivityDeleteRequest() *AlibabaRetailMarketingItemdiscountActivityDeleteRequest{
    return &AlibabaRetailMarketingItemdiscountActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItemdiscountActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.delete"
}

func (r AlibabaRetailMarketingItemdiscountActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItemdiscountActivityDeleteRequest) SetParam(param *ItemDiscountActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItemdiscountActivityDeleteRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r.param
}


package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品特价活动【同城零售】 APIRequest
alibaba.retail.marketing.itemdiscount.activity.update

同城零售单品特价活动更新
*/
type AlibabaRetailMarketingItemdiscountActivityUpdateRequest struct {
    model.Params

    // 创建活动参数
    param   *ItemDiscountActivityOperateRequest 

}

func NewAlibabaRetailMarketingItemdiscountActivityUpdateRequest() *AlibabaRetailMarketingItemdiscountActivityUpdateRequest{
    return &AlibabaRetailMarketingItemdiscountActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItemdiscountActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.update"
}

func (r AlibabaRetailMarketingItemdiscountActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItemdiscountActivityUpdateRequest) SetParam(param *ItemDiscountActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItemdiscountActivityUpdateRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r.param
}


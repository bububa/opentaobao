package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除特价活动商品 APIRequest
alibaba.retail.marketing.itemdiscount.activity.sku.delete

删除活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest struct {
    model.Params

    // 添加活动商品参数
    param   *ItemDiscountActivityElementOperateRequest 

}

func NewAlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest() *AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest{
    return &AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.sku.delete"
}

func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) SetParam(param *ItemDiscountActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItemdiscountActivitySkuDeleteRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
    return r.param
}


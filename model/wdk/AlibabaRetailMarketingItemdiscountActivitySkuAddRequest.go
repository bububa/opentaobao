package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
特价活动新增商品 APIRequest
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuAddRequest struct {
    model.Params

    // 添加活动商品参数
    param   *ItemDiscountActivityElementOperateRequest 

}

func NewAlibabaRetailMarketingItemdiscountActivitySkuAddRequest() *AlibabaRetailMarketingItemdiscountActivitySkuAddRequest{
    return &AlibabaRetailMarketingItemdiscountActivitySkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itemdiscount.activity.sku.add"
}

func (r AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) SetParam(param *ItemDiscountActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItemdiscountActivitySkuAddRequest) GetParam() *ItemDiscountActivityElementOperateRequest {
    return r.param
}


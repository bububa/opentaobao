package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动商品【同城零售】 APIRequest
alibaba.retail.marketing.itempool.activity.sku.delete

删除商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuDeleteRequest struct {
    model.Params

    // 入参
    param   *ItemPoolActivityElementOperateRequest 

}

func NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest() *AlibabaRetailMarketingItempoolActivitySkuDeleteRequest{
    return &AlibabaRetailMarketingItempoolActivitySkuDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.sku.delete"
}

func (r AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) SetParam(param *ItemPoolActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivitySkuDeleteRequest) GetParam() *ItemPoolActivityElementOperateRequest {
    return r.param
}


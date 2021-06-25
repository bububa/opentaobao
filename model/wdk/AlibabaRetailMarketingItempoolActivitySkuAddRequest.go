package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池活动新增商品 APIRequest
alibaba.retail.marketing.itempool.activity.sku.add

新增或更新商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuAddRequest struct {
    model.Params

    // 入参
    param   *ItemPoolActivityElementOperateRequest 

}

func NewAlibabaRetailMarketingItempoolActivitySkuAddRequest() *AlibabaRetailMarketingItempoolActivitySkuAddRequest{
    return &AlibabaRetailMarketingItempoolActivitySkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItempoolActivitySkuAddRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.sku.add"
}

func (r AlibabaRetailMarketingItempoolActivitySkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItempoolActivitySkuAddRequest) SetParam(param *ItemPoolActivityElementOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivitySkuAddRequest) GetParam() *ItemPoolActivityElementOperateRequest {
    return r.param
}


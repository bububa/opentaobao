package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商品池活动【同城零售】 APIRequest
alibaba.retail.marketing.itempool.activity.create

同城零售商品池活动创建
*/
type AlibabaRetailMarketingItempoolActivityCreateRequest struct {
    model.Params

    // 创建商品池活动参数
    param   *ItemPoolActivityOperateRequest 

}

func NewAlibabaRetailMarketingItempoolActivityCreateRequest() *AlibabaRetailMarketingItempoolActivityCreateRequest{
    return &AlibabaRetailMarketingItempoolActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItempoolActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.create"
}

func (r AlibabaRetailMarketingItempoolActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItempoolActivityCreateRequest) SetParam(param *ItemPoolActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivityCreateRequest) GetParam() *ItemPoolActivityOperateRequest {
    return r.param
}


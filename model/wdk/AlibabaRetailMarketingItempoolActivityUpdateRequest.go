package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品池活动【同城零售】 APIRequest
alibaba.retail.marketing.itempool.activity.update

同城零售商品池活动更新
*/
type AlibabaRetailMarketingItempoolActivityUpdateRequest struct {
    model.Params

    // 更新商品池活动参数
    param   *ItemPoolActivityOperateRequest 

}

func NewAlibabaRetailMarketingItempoolActivityUpdateRequest() *AlibabaRetailMarketingItempoolActivityUpdateRequest{
    return &AlibabaRetailMarketingItempoolActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingItempoolActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.itempool.activity.update"
}

func (r AlibabaRetailMarketingItempoolActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingItempoolActivityUpdateRequest) SetParam(param *ItemPoolActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingItempoolActivityUpdateRequest) GetParam() *ItemPoolActivityOperateRequest {
    return r.param
}


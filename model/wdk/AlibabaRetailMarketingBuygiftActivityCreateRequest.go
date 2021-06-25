package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建买赠活动 APIRequest
alibaba.retail.marketing.buygift.activity.create

同城供给买赠活动创建
*/
type AlibabaRetailMarketingBuygiftActivityCreateRequest struct {
    model.Params

    // 创建活动参数
    param   *BuyGiftActivityOperateRequest 

}

func NewAlibabaRetailMarketingBuygiftActivityCreateRequest() *AlibabaRetailMarketingBuygiftActivityCreateRequest{
    return &AlibabaRetailMarketingBuygiftActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingBuygiftActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.create"
}

func (r AlibabaRetailMarketingBuygiftActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingBuygiftActivityCreateRequest) SetParam(param *BuyGiftActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingBuygiftActivityCreateRequest) GetParam() *BuyGiftActivityOperateRequest {
    return r.param
}


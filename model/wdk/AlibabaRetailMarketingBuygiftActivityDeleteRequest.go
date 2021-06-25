package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动 APIRequest
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingBuygiftActivityDeleteRequest struct {
    model.Params

    // 删除活动参数
    param   *ItemDiscountActivityOperateRequest 

}

func NewAlibabaRetailMarketingBuygiftActivityDeleteRequest() *AlibabaRetailMarketingBuygiftActivityDeleteRequest{
    return &AlibabaRetailMarketingBuygiftActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingBuygiftActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.delete"
}

func (r AlibabaRetailMarketingBuygiftActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingBuygiftActivityDeleteRequest) SetParam(param *ItemDiscountActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingBuygiftActivityDeleteRequest) GetParam() *ItemDiscountActivityOperateRequest {
    return r.param
}


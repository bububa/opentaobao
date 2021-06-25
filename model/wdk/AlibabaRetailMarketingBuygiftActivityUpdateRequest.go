package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品买赠活动 APIRequest
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新
*/
type AlibabaRetailMarketingBuygiftActivityUpdateRequest struct {
    model.Params

    // 更新单品买赠活动参数
    param   *BuyGiftActivityOperateRequest 

}

func NewAlibabaRetailMarketingBuygiftActivityUpdateRequest() *AlibabaRetailMarketingBuygiftActivityUpdateRequest{
    return &AlibabaRetailMarketingBuygiftActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingBuygiftActivityUpdateRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.update"
}

func (r AlibabaRetailMarketingBuygiftActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingBuygiftActivityUpdateRequest) SetParam(param *BuyGiftActivityOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingBuygiftActivityUpdateRequest) GetParam() *BuyGiftActivityOperateRequest {
    return r.param
}


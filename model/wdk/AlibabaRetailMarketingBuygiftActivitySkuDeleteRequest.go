package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动商品 APIRequest
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest struct {
    model.Params

    // 删除买赠活动商品参数
    param   *BuyGiftActivitySkuOperateRequest 

}

func NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest() *AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest{
    return &AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.sku.delete"
}

func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) SetParam(param *BuyGiftActivitySkuOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
    return r.param
}


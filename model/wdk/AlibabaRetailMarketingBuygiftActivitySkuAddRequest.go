package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加单品买赠活动商品 APIRequest
alibaba.retail.marketing.buygift.activity.sku.add

新增或更新单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuAddRequest struct {
    model.Params

    // 添加活动商品参数
    param   *BuyGiftActivitySkuOperateRequest 

}

func NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest() *AlibabaRetailMarketingBuygiftActivitySkuAddRequest{
    return &AlibabaRetailMarketingBuygiftActivitySkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailMarketingBuygiftActivitySkuAddRequest) GetApiMethodName() string {
    return "alibaba.retail.marketing.buygift.activity.sku.add"
}

func (r AlibabaRetailMarketingBuygiftActivitySkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailMarketingBuygiftActivitySkuAddRequest) SetParam(param *BuyGiftActivitySkuOperateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaRetailMarketingBuygiftActivitySkuAddRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
    return r.param
}


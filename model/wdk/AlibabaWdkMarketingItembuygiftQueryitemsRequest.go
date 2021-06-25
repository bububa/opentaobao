package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动下的商品 APIRequest
alibaba.wdk.marketing.itembuygift.queryitems

查询买赠活动下的商品
*/
type AlibabaWdkMarketingItembuygiftQueryitemsRequest struct {
    model.Params

    // 查询入参
    param   *ActivitySkuQuery 

}

func NewAlibabaWdkMarketingItembuygiftQueryitemsRequest() *AlibabaWdkMarketingItembuygiftQueryitemsRequest{
    return &AlibabaWdkMarketingItembuygiftQueryitemsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItembuygiftQueryitemsRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.queryitems"
}

func (r AlibabaWdkMarketingItembuygiftQueryitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItembuygiftQueryitemsRequest) SetParam(param *ActivitySkuQuery) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftQueryitemsRequest) GetParam() *ActivitySkuQuery {
    return r.param
}


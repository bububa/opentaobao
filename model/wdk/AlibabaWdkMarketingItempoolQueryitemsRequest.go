package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品池活动下的商品 APIRequest
alibaba.wdk.marketing.itempool.queryitems

查询商品池活动下面的商品
*/
type AlibabaWdkMarketingItempoolQueryitemsRequest struct {
    model.Params

    // 查询入参
    param   *ActivitySkuQuery 

}

func NewAlibabaWdkMarketingItempoolQueryitemsRequest() *AlibabaWdkMarketingItempoolQueryitemsRequest{
    return &AlibabaWdkMarketingItempoolQueryitemsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolQueryitemsRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.queryitems"
}

func (r AlibabaWdkMarketingItempoolQueryitemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolQueryitemsRequest) SetParam(param *ActivitySkuQuery) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItempoolQueryitemsRequest) GetParam() *ActivitySkuQuery {
    return r.param
}


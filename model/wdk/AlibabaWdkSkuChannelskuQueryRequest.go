package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商品 APIRequest
alibaba.wdk.sku.channelsku.query

查询渠道商品
*/
type AlibabaWdkSkuChannelskuQueryRequest struct {
    model.Params

    // 查询渠道商品的入参
    param   *ChannelSkuQueryDo 

}

func NewAlibabaWdkSkuChannelskuQueryRequest() *AlibabaWdkSkuChannelskuQueryRequest{
    return &AlibabaWdkSkuChannelskuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuChannelskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.query"
}

func (r AlibabaWdkSkuChannelskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuChannelskuQueryRequest) SetParam(param *ChannelSkuQueryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuChannelskuQueryRequest) GetParam() *ChannelSkuQueryDo {
    return r.param
}


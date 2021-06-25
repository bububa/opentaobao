package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新渠道商品 APIRequest
alibaba.wdk.sku.channelsku.update

批量更新渠道商品，商家通过Top接入
*/
type AlibabaWdkSkuChannelskuUpdateRequest struct {
    model.Params

    // 请求参数
    paramList   []ChannelSkuDo 

}

func NewAlibabaWdkSkuChannelskuUpdateRequest() *AlibabaWdkSkuChannelskuUpdateRequest{
    return &AlibabaWdkSkuChannelskuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuChannelskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.update"
}

func (r AlibabaWdkSkuChannelskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuChannelskuUpdateRequest) SetParamList(paramList []ChannelSkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaWdkSkuChannelskuUpdateRequest) GetParamList() []ChannelSkuDo {
    return r.paramList
}


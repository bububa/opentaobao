package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增渠道商品 APIRequest
alibaba.wdk.sku.channelsku.add

盒马帮1期新增渠道商品
*/
type AlibabaWdkSkuChannelskuAddRequest struct {
    model.Params

    // 入参模型
    chSkuDOList   []ChannelSkuDo 

}

func NewAlibabaWdkSkuChannelskuAddRequest() *AlibabaWdkSkuChannelskuAddRequest{
    return &AlibabaWdkSkuChannelskuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuChannelskuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.add"
}

func (r AlibabaWdkSkuChannelskuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuChannelskuAddRequest) SetChSkuDOList(chSkuDOList []ChannelSkuDo) error {
    r.chSkuDOList = chSkuDOList
    r.Set("ch_sku_d_o_list", chSkuDOList)
    return nil
}

func (r AlibabaWdkSkuChannelskuAddRequest) GetChSkuDOList() []ChannelSkuDo {
    return r.chSkuDOList
}


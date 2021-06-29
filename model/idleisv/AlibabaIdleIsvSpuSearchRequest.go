package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
spu搜索接口 APIRequest
alibaba.idle.isv.spu.search

搜索的品牌和型号，供服务商进行选择
*/
type AlibabaIdleIsvSpuSearchRequest struct {
    model.Params

    // 闲鱼渠道类目的id
    channelCatId   string 

    // 搜索的文本
    searchText   string 

}

func NewAlibabaIdleIsvSpuSearchRequest() *AlibabaIdleIsvSpuSearchRequest{
    return &AlibabaIdleIsvSpuSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvSpuSearchRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.spu.search"
}

func (r AlibabaIdleIsvSpuSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvSpuSearchRequest) SetChannelCatId(channelCatId string) error {
    r.channelCatId = channelCatId
    r.Set("channel_cat_id", channelCatId)
    return nil
}

func (r AlibabaIdleIsvSpuSearchRequest) GetChannelCatId() string {
    return r.channelCatId
}

func (r *AlibabaIdleIsvSpuSearchRequest) SetSearchText(searchText string) error {
    r.searchText = searchText
    r.Set("search_text", searchText)
    return nil
}

func (r AlibabaIdleIsvSpuSearchRequest) GetSearchText() string {
    return r.searchText
}


package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
spu搜索接口 API请求
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

// 初始化AlibabaIdleIsvSpuSearchRequest对象
func NewAlibabaIdleIsvSpuSearchRequest() *AlibabaIdleIsvSpuSearchRequest{
    return &AlibabaIdleIsvSpuSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvSpuSearchRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.spu.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvSpuSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCatId Setter
// 闲鱼渠道类目的id
func (r *AlibabaIdleIsvSpuSearchRequest) SetChannelCatId(channelCatId string) error {
    r.channelCatId = channelCatId
    r.Set("channel_cat_id", channelCatId)
    return nil
}

// ChannelCatId Getter
func (r AlibabaIdleIsvSpuSearchRequest) GetChannelCatId() string {
    return r.channelCatId
}
// SearchText Setter
// 搜索的文本
func (r *AlibabaIdleIsvSpuSearchRequest) SetSearchText(searchText string) error {
    r.searchText = searchText
    r.Set("search_text", searchText)
    return nil
}

// SearchText Getter
func (r AlibabaIdleIsvSpuSearchRequest) GetSearchText() string {
    return r.searchText
}

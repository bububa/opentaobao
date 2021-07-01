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
type AlibabaIdleIsvSpuSearchAPIRequest struct {
    model.Params
    // 闲鱼渠道类目的id
    _channelCatId   string
    // 搜索的文本
    _searchText   string
}

// 初始化AlibabaIdleIsvSpuSearchAPIRequest对象
func NewAlibabaIdleIsvSpuSearchRequest() *AlibabaIdleIsvSpuSearchAPIRequest{
    return &AlibabaIdleIsvSpuSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvSpuSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.spu.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvSpuSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCatId Setter
// 闲鱼渠道类目的id
func (r *AlibabaIdleIsvSpuSearchAPIRequest) SetChannelCatId(_channelCatId string) error {
    r._channelCatId = _channelCatId
    r.Set("channel_cat_id", _channelCatId)
    return nil
}

// ChannelCatId Getter
func (r AlibabaIdleIsvSpuSearchAPIRequest) GetChannelCatId() string {
    return r._channelCatId
}
// SearchText Setter
// 搜索的文本
func (r *AlibabaIdleIsvSpuSearchAPIRequest) SetSearchText(_searchText string) error {
    r._searchText = _searchText
    r.Set("search_text", _searchText)
    return nil
}

// SearchText Getter
func (r AlibabaIdleIsvSpuSearchAPIRequest) GetSearchText() string {
    return r._searchText
}

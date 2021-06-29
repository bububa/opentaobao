package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的推荐关键词列表 API请求
taobao.simba.keywords.recommend.get

取得一个推广组的推荐关键词列表
*/
type TaobaoSimbaKeywordsRecommendGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组ID
    _adgroupId   int64
    // 搜索量,设置此值后返回的就是大于此搜索量的词列表
    _search   int64
    // 相关度
    _pertinence   string
    // 返回的每页数据量大小,最大200
    _pageSize   int64
    // 返回的第几页数据，默认为1
    _pageNo   int64
    // 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
    _orderBy   string
}

// 初始化TaobaoSimbaKeywordsRecommendGetRequest对象
func NewTaobaoSimbaKeywordsRecommendGetRequest() *TaobaoSimbaKeywordsRecommendGetRequest{
    return &TaobaoSimbaKeywordsRecommendGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.recommend.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// Search Setter
// 搜索量,设置此值后返回的就是大于此搜索量的词列表
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetSearch(_search int64) error {
    r._search = _search
    r.Set("search", _search)
    return nil
}

// Search Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetSearch() int64 {
    return r._search
}
// Pertinence Setter
// 相关度
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetPertinence(_pertinence string) error {
    r._pertinence = _pertinence
    r.Set("pertinence", _pertinence)
    return nil
}

// Pertinence Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetPertinence() string {
    return r._pertinence
}
// PageSize Setter
// 返回的每页数据量大小,最大200
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// OrderBy Setter
// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetOrderBy(_orderBy string) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoSimbaKeywordsRecommendGetRequest) GetOrderBy() string {
    return r._orderBy
}

package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的推荐关键词列表 APIRequest
taobao.simba.keywords.recommend.get

取得一个推广组的推荐关键词列表
*/
type TaobaoSimbaKeywordsRecommendGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组ID
    adgroupId   int64 

    // 搜索量,设置此值后返回的就是大于此搜索量的词列表
    search   int64 

    // 相关度
    pertinence   string 

    // 返回的每页数据量大小,最大200
    pageSize   int64 

    // 返回的第几页数据，默认为1
    pageNo   int64 

    // 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
    orderBy   string 

}

func NewTaobaoSimbaKeywordsRecommendGetRequest() *TaobaoSimbaKeywordsRecommendGetRequest{
    return &TaobaoSimbaKeywordsRecommendGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.recommend.get"
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetSearch(search int64) error {
    r.search = search
    r.Set("search", search)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetSearch() int64 {
    return r.search
}

func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetPertinence(pertinence string) error {
    r.pertinence = pertinence
    r.Set("pertinence", pertinence)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetPertinence() string {
    return r.pertinence
}

func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaKeywordsRecommendGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r TaobaoSimbaKeywordsRecommendGetRequest) GetOrderBy() string {
    return r.orderBy
}


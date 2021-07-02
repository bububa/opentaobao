package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsRecommendGetAPIRequest 取得一个推广组的推荐关键词列表 API请求
// taobao.simba.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
type TaobaoSimbaKeywordsRecommendGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组ID
	_adgroupId int64
	// 搜索量,设置此值后返回的就是大于此搜索量的词列表
	_search int64
	// 相关度
	_pertinence string
	// 返回的每页数据量大小,最大200
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
	// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
	_orderBy string
}

// NewTaobaoSimbaKeywordsRecommendGetRequest 初始化TaobaoSimbaKeywordsRecommendGetAPIRequest对象
func NewTaobaoSimbaKeywordsRecommendGetRequest() *TaobaoSimbaKeywordsRecommendGetAPIRequest {
	return &TaobaoSimbaKeywordsRecommendGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is Search Setter
// 搜索量,设置此值后返回的就是大于此搜索量的词列表
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetSearch(_search int64) error {
	r._search = _search
	r.Set("search", _search)
	return nil
}

// Get Search Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetSearch() int64 {
	return r._search
}

// Set is Pertinence Setter
// 相关度
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetPertinence(_pertinence string) error {
	r._pertinence = _pertinence
	r.Set("pertinence", _pertinence)
	return nil
}

// Get Pertinence Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetPertinence() string {
	return r._pertinence
}

// Set is PageSize Setter
// 返回的每页数据量大小,最大200
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is OrderBy Setter
// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// Get OrderBy Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

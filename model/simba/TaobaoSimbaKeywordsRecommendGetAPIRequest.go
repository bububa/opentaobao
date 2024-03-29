package simba

import (
	"net/url"
	"sync"

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
	// 相关度
	_pertinence string
	// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
	_orderBy string
	// 推广组ID
	_adgroupId int64
	// 搜索量,设置此值后返回的就是大于此搜索量的词列表
	_search int64
	// 返回的每页数据量大小,最大200
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
}

// NewTaobaoSimbaKeywordsRecommendGetRequest 初始化TaobaoSimbaKeywordsRecommendGetAPIRequest对象
func NewTaobaoSimbaKeywordsRecommendGetRequest() *TaobaoSimbaKeywordsRecommendGetAPIRequest {
	return &TaobaoSimbaKeywordsRecommendGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) Reset() {
	r._nick = ""
	r._pertinence = ""
	r._orderBy = ""
	r._adgroupId = 0
	r._search = 0
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetNick() string {
	return r._nick
}

// SetPertinence is Pertinence Setter
// 相关度
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetPertinence(_pertinence string) error {
	r._pertinence = _pertinence
	r.Set("pertinence", _pertinence)
	return nil
}

// GetPertinence Pertinence Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetPertinence() string {
	return r._pertinence
}

// SetOrderBy is OrderBy Setter
// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetAdgroupId is AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetSearch is Search Setter
// 搜索量,设置此值后返回的就是大于此搜索量的词列表
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetSearch(_search int64) error {
	r._search = _search
	r.Set("search", _search)
	return nil
}

// GetSearch Search Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetSearch() int64 {
	return r._search
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,最大200
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaKeywordsRecommendGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaKeywordsRecommendGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaKeywordsRecommendGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordsRecommendGetRequest()
	},
}

// GetTaobaoSimbaKeywordsRecommendGetRequest 从 sync.Pool 获取 TaobaoSimbaKeywordsRecommendGetAPIRequest
func GetTaobaoSimbaKeywordsRecommendGetAPIRequest() *TaobaoSimbaKeywordsRecommendGetAPIRequest {
	return poolTaobaoSimbaKeywordsRecommendGetAPIRequest.Get().(*TaobaoSimbaKeywordsRecommendGetAPIRequest)
}

// ReleaseTaobaoSimbaKeywordsRecommendGetAPIRequest 将 TaobaoSimbaKeywordsRecommendGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordsRecommendGetAPIRequest(v *TaobaoSimbaKeywordsRecommendGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordsRecommendGetAPIRequest.Put(v)
}

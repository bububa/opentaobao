package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsrecommendgetAPIRequest 取得一个推广组的推荐关键词列表 API请求
// taobao.simba.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
type TaobaosimbakeywordsrecommendgetAPIRequest struct {
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

// NewTaobaosimbakeywordsrecommendgetRequest 初始化TaobaosimbakeywordsrecommendgetAPIRequest对象
func NewTaobaosimbakeywordsrecommendgetRequest() *TaobaosimbakeywordsrecommendgetAPIRequest {
	return &TaobaosimbakeywordsrecommendgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetNick() string {
	return r._nick
}

// SetPertinence is Pertinence Setter
// 相关度
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetPertinence(_pertinence string) error {
	r._pertinence = _pertinence
	r.Set("pertinence", _pertinence)
	return nil
}

// GetPertinence Pertinence Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetPertinence() string {
	return r._pertinence
}

// SetOrderBy is OrderBy Setter
// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetAdgroupId is AdgroupId Setter
// 推广组ID
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetSearch is Search Setter
// 搜索量,设置此值后返回的就是大于此搜索量的词列表
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetSearch(_search int64) error {
	r._search = _search
	r.Set("search", _search)
	return nil
}

// GetSearch Search Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetSearch() int64 {
	return r._search
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,最大200
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaosimbakeywordsrecommendgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbakeywordsrecommendgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

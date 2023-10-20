package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongsinfoqueryAPIRequest 搜索歌曲列表 API请求
// xiami.content.songs.info.query
//
// 多维度查询歌曲列表
type XiamicontentsongsinfoqueryAPIRequest struct {
	model.Params
	// 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
	_searchTerms []SearchTermsDto
	// tag搜索条件，tag尽量不要超过50个
	_tagOptional *SongCatsSearchDto
	// 排序,默认按照最新排序 1最新 2本周最热 3本月最热
	_orderBy int64
	// 分页信息
	_page *PagingVo
}

// NewXiamicontentsongsinfoqueryRequest 初始化XiamicontentsongsinfoqueryAPIRequest对象
func NewXiamicontentsongsinfoqueryRequest() *XiamicontentsongsinfoqueryAPIRequest {
	return &XiamicontentsongsinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentsongsinfoqueryAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentsongsinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentsongsinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchTerms is SearchTerms Setter
// 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
func (r *XiamicontentsongsinfoqueryAPIRequest) SetSearchTerms(_searchTerms []SearchTermsDto) error {
	r._searchTerms = _searchTerms
	r.Set("search_terms", _searchTerms)
	return nil
}

// GetSearchTerms SearchTerms Getter
func (r XiamicontentsongsinfoqueryAPIRequest) GetSearchTerms() []SearchTermsDto {
	return r._searchTerms
}

// SetTagOptional is TagOptional Setter
// tag搜索条件，tag尽量不要超过50个
func (r *XiamicontentsongsinfoqueryAPIRequest) SetTagOptional(_tagOptional *SongCatsSearchDto) error {
	r._tagOptional = _tagOptional
	r.Set("tag_optional", _tagOptional)
	return nil
}

// GetTagOptional TagOptional Getter
func (r XiamicontentsongsinfoqueryAPIRequest) GetTagOptional() *SongCatsSearchDto {
	return r._tagOptional
}

// SetOrderBy is OrderBy Setter
// 排序,默认按照最新排序 1最新 2本周最热 3本月最热
func (r *XiamicontentsongsinfoqueryAPIRequest) SetOrderBy(_orderBy int64) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r XiamicontentsongsinfoqueryAPIRequest) GetOrderBy() int64 {
	return r._orderBy
}

// SetPage is Page Setter
// 分页信息
func (r *XiamicontentsongsinfoqueryAPIRequest) SetPage(_page *PagingVo) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r XiamicontentsongsinfoqueryAPIRequest) GetPage() *PagingVo {
	return r._page
}

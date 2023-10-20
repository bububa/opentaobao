package xiamicontent

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMusicInfoQueryAPIRequest 搜索音乐 API请求
// xiami.content.music.info.query
//
// (批量)获取歌曲信息
type XiamiContentMusicInfoQueryAPIRequest struct {
	model.Params
	// 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
	_searchTerms *SearchTermsDto
	// tag搜索条件，tag尽量不要超过50个
	_tagOptional *SongCatsSearchDto
	// 排序,默认按照最新排序 1最新 2本周最热 3本月最热
	_orderBy int64
	// 分页信息
	_page *PagingVo
}

// NewXiamiContentMusicInfoQueryRequest 初始化XiamiContentMusicInfoQueryAPIRequest对象
func NewXiamiContentMusicInfoQueryRequest() *XiamiContentMusicInfoQueryAPIRequest {
	return &XiamiContentMusicInfoQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiContentMusicInfoQueryAPIRequest) Reset() {
	r._searchTerms = nil
	r._tagOptional = nil
	r._orderBy = 0
	r._page = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentMusicInfoQueryAPIRequest) GetApiMethodName() string {
	return "xiami.content.music.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentMusicInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentMusicInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchTerms is SearchTerms Setter
// 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
func (r *XiamiContentMusicInfoQueryAPIRequest) SetSearchTerms(_searchTerms *SearchTermsDto) error {
	r._searchTerms = _searchTerms
	r.Set("search_terms", _searchTerms)
	return nil
}

// GetSearchTerms SearchTerms Getter
func (r XiamiContentMusicInfoQueryAPIRequest) GetSearchTerms() *SearchTermsDto {
	return r._searchTerms
}

// SetTagOptional is TagOptional Setter
// tag搜索条件，tag尽量不要超过50个
func (r *XiamiContentMusicInfoQueryAPIRequest) SetTagOptional(_tagOptional *SongCatsSearchDto) error {
	r._tagOptional = _tagOptional
	r.Set("tag_optional", _tagOptional)
	return nil
}

// GetTagOptional TagOptional Getter
func (r XiamiContentMusicInfoQueryAPIRequest) GetTagOptional() *SongCatsSearchDto {
	return r._tagOptional
}

// SetOrderBy is OrderBy Setter
// 排序,默认按照最新排序 1最新 2本周最热 3本月最热
func (r *XiamiContentMusicInfoQueryAPIRequest) SetOrderBy(_orderBy int64) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r XiamiContentMusicInfoQueryAPIRequest) GetOrderBy() int64 {
	return r._orderBy
}

// SetPage is Page Setter
// 分页信息
func (r *XiamiContentMusicInfoQueryAPIRequest) SetPage(_page *PagingVo) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r XiamiContentMusicInfoQueryAPIRequest) GetPage() *PagingVo {
	return r._page
}

var poolXiamiContentMusicInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiContentMusicInfoQueryRequest()
	},
}

// GetXiamiContentMusicInfoQueryRequest 从 sync.Pool 获取 XiamiContentMusicInfoQueryAPIRequest
func GetXiamiContentMusicInfoQueryAPIRequest() *XiamiContentMusicInfoQueryAPIRequest {
	return poolXiamiContentMusicInfoQueryAPIRequest.Get().(*XiamiContentMusicInfoQueryAPIRequest)
}

// ReleaseXiamiContentMusicInfoQueryAPIRequest 将 XiamiContentMusicInfoQueryAPIRequest 放入 sync.Pool
func ReleaseXiamiContentMusicInfoQueryAPIRequest(v *XiamiContentMusicInfoQueryAPIRequest) {
	v.Reset()
	poolXiamiContentMusicInfoQueryAPIRequest.Put(v)
}

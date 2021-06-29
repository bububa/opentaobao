package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索歌曲列表 API请求
xiami.content.songs.info.query

多维度查询歌曲列表
*/
type XiamiContentSongsInfoQueryRequest struct {
    model.Params
    // 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
    _searchTerms   []SearchTermsDTO
    // tag搜索条件，tag尽量不要超过50个
    _tagOptional   *SongCatsSearchDTO
    // 排序,默认按照最新排序 1最新 2本周最热 3本月最热
    _orderBy   int64
    // 分页信息
    _page   *PagingVo
}

// 初始化XiamiContentSongsInfoQueryRequest对象
func NewXiamiContentSongsInfoQueryRequest() *XiamiContentSongsInfoQueryRequest{
    return &XiamiContentSongsInfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentSongsInfoQueryRequest) GetApiMethodName() string {
    return "xiami.content.songs.info.query"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentSongsInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SearchTerms Setter
// 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
func (r *XiamiContentSongsInfoQueryRequest) SetSearchTerms(_searchTerms []SearchTermsDTO) error {
    r._searchTerms = _searchTerms
    r.Set("search_terms", _searchTerms)
    return nil
}

// SearchTerms Getter
func (r XiamiContentSongsInfoQueryRequest) GetSearchTerms() []SearchTermsDTO {
    return r._searchTerms
}
// TagOptional Setter
// tag搜索条件，tag尽量不要超过50个
func (r *XiamiContentSongsInfoQueryRequest) SetTagOptional(_tagOptional *SongCatsSearchDTO) error {
    r._tagOptional = _tagOptional
    r.Set("tag_optional", _tagOptional)
    return nil
}

// TagOptional Getter
func (r XiamiContentSongsInfoQueryRequest) GetTagOptional() *SongCatsSearchDTO {
    return r._tagOptional
}
// OrderBy Setter
// 排序,默认按照最新排序 1最新 2本周最热 3本月最热
func (r *XiamiContentSongsInfoQueryRequest) SetOrderBy(_orderBy int64) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r XiamiContentSongsInfoQueryRequest) GetOrderBy() int64 {
    return r._orderBy
}
// Page Setter
// 分页信息
func (r *XiamiContentSongsInfoQueryRequest) SetPage(_page *PagingVo) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r XiamiContentSongsInfoQueryRequest) GetPage() *PagingVo {
    return r._page
}

package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索歌曲列表 APIRequest
xiami.content.songs.info.query

多维度查询歌曲列表
*/
type XiamiContentSongsInfoQueryRequest struct {
    model.Params

    // 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
    searchTerms   []SearchTermsDto 

    // tag搜索条件，tag尽量不要超过50个
    tagOptional   *SongCatsSearchDto 

    // 排序,默认按照最新排序 1最新 2本周最热 3本月最热
    orderBy   int64 

    // 分页信息
    page   *PagingVo 

}

func NewXiamiContentSongsInfoQueryRequest() *XiamiContentSongsInfoQueryRequest{
    return &XiamiContentSongsInfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentSongsInfoQueryRequest) GetApiMethodName() string {
    return "xiami.content.songs.info.query"
}

func (r XiamiContentSongsInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentSongsInfoQueryRequest) SetSearchTerms(searchTerms []SearchTermsDto) error {
    r.searchTerms = searchTerms
    r.Set("search_terms", searchTerms)
    return nil
}

func (r XiamiContentSongsInfoQueryRequest) GetSearchTerms() []SearchTermsDto {
    return r.searchTerms
}

func (r *XiamiContentSongsInfoQueryRequest) SetTagOptional(tagOptional *SongCatsSearchDto) error {
    r.tagOptional = tagOptional
    r.Set("tag_optional", tagOptional)
    return nil
}

func (r XiamiContentSongsInfoQueryRequest) GetTagOptional() *SongCatsSearchDto {
    return r.tagOptional
}

func (r *XiamiContentSongsInfoQueryRequest) SetOrderBy(orderBy int64) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r XiamiContentSongsInfoQueryRequest) GetOrderBy() int64 {
    return r.orderBy
}

func (r *XiamiContentSongsInfoQueryRequest) SetPage(page *PagingVo) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r XiamiContentSongsInfoQueryRequest) GetPage() *PagingVo {
    return r.page
}


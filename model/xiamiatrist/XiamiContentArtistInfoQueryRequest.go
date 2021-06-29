package xiamiatrist

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索艺人列表 APIRequest
xiami.content.artist.info.query

根据查询条件，搜索相关艺人列表
*/
type XiamiContentArtistInfoQueryRequest struct {
    model.Params

    // 性别：1男性 2女性 3乐队
    gender   int64 

    // 语种：1华语 2日本 3韩国 4欧美 5其他
    language   int64 

    // 流派: 1嘻哈(说唱),2流行，3摇滚，4布鲁斯，5爵士，6雷鬼，7世界音乐，8拉丁，9电子，10节奏布鲁斯，11实验，12轻音乐，13新世纪，14舞台 / 银幕 / 娱乐      * 15唱作人，16民谣，18金属，19中国特色，20乡村，21古典，22儿童，23有声书，24动漫，25朋克
    genre   int64 

    // 分页信息
    pageReq   *PagingVo 

}

func NewXiamiContentArtistInfoQueryRequest() *XiamiContentArtistInfoQueryRequest{
    return &XiamiContentArtistInfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentArtistInfoQueryRequest) GetApiMethodName() string {
    return "xiami.content.artist.info.query"
}

func (r XiamiContentArtistInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentArtistInfoQueryRequest) SetGender(gender int64) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

func (r XiamiContentArtistInfoQueryRequest) GetGender() int64 {
    return r.gender
}

func (r *XiamiContentArtistInfoQueryRequest) SetLanguage(language int64) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r XiamiContentArtistInfoQueryRequest) GetLanguage() int64 {
    return r.language
}

func (r *XiamiContentArtistInfoQueryRequest) SetGenre(genre int64) error {
    r.genre = genre
    r.Set("genre", genre)
    return nil
}

func (r XiamiContentArtistInfoQueryRequest) GetGenre() int64 {
    return r.genre
}

func (r *XiamiContentArtistInfoQueryRequest) SetPageReq(pageReq *PagingVo) error {
    r.pageReq = pageReq
    r.Set("page_req", pageReq)
    return nil
}

func (r XiamiContentArtistInfoQueryRequest) GetPageReq() *PagingVo {
    return r.pageReq
}


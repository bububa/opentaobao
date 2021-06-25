package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐－风格，流派专辑列表 APIRequest
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表
*/
type AlibabaXiamiApiTagGenreAlbumGetRequest struct {
    model.Params

    // 1:风格，2:流派
    type   int64 

    // 风格，流派id
    id   int64 

    // 页数
    page   int64 

    // 每页数量
    limit   int64 

}

func NewAlibabaXiamiApiTagGenreAlbumGetRequest() *AlibabaXiamiApiTagGenreAlbumGetRequest{
    return &AlibabaXiamiApiTagGenreAlbumGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.tag.genre.album.get"
}

func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetLimit() int64 {
    return r.limit
}


package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米-风格，流派歌曲 APIRequest
alibaba.xiami.api.tag.genre.song.get

虾米-风格，流派歌曲
*/
type AlibabaXiamiApiTagGenreSongGetRequest struct {
    model.Params

    // 1:风格，2：流派
    type   string 

    // 风格或流派id
    id   int64 

    // 页数
    page   int64 

    // 每页数量
    limit   int64 

}

func NewAlibabaXiamiApiTagGenreSongGetRequest() *AlibabaXiamiApiTagGenreSongGetRequest{
    return &AlibabaXiamiApiTagGenreSongGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiTagGenreSongGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.tag.genre.song.get"
}

func (r AlibabaXiamiApiTagGenreSongGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaXiamiApiTagGenreSongGetRequest) GetType() string {
    return r.type
}

func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaXiamiApiTagGenreSongGetRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaXiamiApiTagGenreSongGetRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaXiamiApiTagGenreSongGetRequest) GetLimit() int64 {
    return r.limit
}


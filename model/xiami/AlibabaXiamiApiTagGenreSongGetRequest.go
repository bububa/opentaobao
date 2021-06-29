package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米-风格，流派歌曲 API请求
alibaba.xiami.api.tag.genre.song.get

虾米-风格，流派歌曲
*/
type AlibabaXiamiApiTagGenreSongGetRequest struct {
    model.Params
    // 1:风格，2：流派
    _type   string
    // 风格或流派id
    _id   int64
    // 页数
    _page   int64
    // 每页数量
    _limit   int64
}

// 初始化AlibabaXiamiApiTagGenreSongGetRequest对象
func NewAlibabaXiamiApiTagGenreSongGetRequest() *AlibabaXiamiApiTagGenreSongGetRequest{
    return &AlibabaXiamiApiTagGenreSongGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiTagGenreSongGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.tag.genre.song.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiTagGenreSongGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 1:风格，2：流派
func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiTagGenreSongGetRequest) GetType() string {
    return r._type
}
// Id Setter
// 风格或流派id
func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiTagGenreSongGetRequest) GetId() int64 {
    return r._id
}
// Page Setter
// 页数
func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaXiamiApiTagGenreSongGetRequest) GetPage() int64 {
    return r._page
}
// Limit Setter
// 每页数量
func (r *AlibabaXiamiApiTagGenreSongGetRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiTagGenreSongGetRequest) GetLimit() int64 {
    return r._limit
}

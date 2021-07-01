package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐－风格，流派专辑列表 API请求
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表
*/
type AlibabaXiamiApiTagGenreAlbumGetAPIRequest struct {
    model.Params
    // 1:风格，2:流派
    _type   int64
    // 风格，流派id
    _id   int64
    // 页数
    _page   int64
    // 每页数量
    _limit   int64
}

// 初始化AlibabaXiamiApiTagGenreAlbumGetAPIRequest对象
func NewAlibabaXiamiApiTagGenreAlbumGetRequest() *AlibabaXiamiApiTagGenreAlbumGetAPIRequest{
    return &AlibabaXiamiApiTagGenreAlbumGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiTagGenreAlbumGetAPIRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.tag.genre.album.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiTagGenreAlbumGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 1:风格，2:流派
func (r *AlibabaXiamiApiTagGenreAlbumGetAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiTagGenreAlbumGetAPIRequest) GetType() int64 {
    return r._type
}
// Id Setter
// 风格，流派id
func (r *AlibabaXiamiApiTagGenreAlbumGetAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiTagGenreAlbumGetAPIRequest) GetId() int64 {
    return r._id
}
// Page Setter
// 页数
func (r *AlibabaXiamiApiTagGenreAlbumGetAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaXiamiApiTagGenreAlbumGetAPIRequest) GetPage() int64 {
    return r._page
}
// Limit Setter
// 每页数量
func (r *AlibabaXiamiApiTagGenreAlbumGetAPIRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiTagGenreAlbumGetAPIRequest) GetLimit() int64 {
    return r._limit
}

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

// 初始化AlibabaXiamiApiTagGenreAlbumGetRequest对象
func NewAlibabaXiamiApiTagGenreAlbumGetRequest() *AlibabaXiamiApiTagGenreAlbumGetRequest{
    return &AlibabaXiamiApiTagGenreAlbumGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.tag.genre.album.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 1:风格，2:流派
func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetType() int64 {
    return r.type
}
// Id Setter
// 风格，流派id
func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetId() int64 {
    return r.id
}
// Page Setter
// 页数
func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetPage() int64 {
    return r.page
}
// Limit Setter
// 每页数量
func (r *AlibabaXiamiApiTagGenreAlbumGetRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiTagGenreAlbumGetRequest) GetLimit() int64 {
    return r.limit
}

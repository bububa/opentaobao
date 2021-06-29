package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
艺人专辑 API请求
alibaba.xiami.api.artist.albums.get

艺人专辑
*/
type AlibabaXiamiApiArtistAlbumsGetRequest struct {
    model.Params
    // 歌曲数量
    _limit   int64
    // 页码
    _page   int64
    // 艺人id
    _id   int64
}

// 初始化AlibabaXiamiApiArtistAlbumsGetRequest对象
func NewAlibabaXiamiApiArtistAlbumsGetRequest() *AlibabaXiamiApiArtistAlbumsGetRequest{
    return &AlibabaXiamiApiArtistAlbumsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistAlbumsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.albums.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistAlbumsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Limit Setter
// 歌曲数量
func (r *AlibabaXiamiApiArtistAlbumsGetRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r AlibabaXiamiApiArtistAlbumsGetRequest) GetLimit() int64 {
    return r._limit
}
// Page Setter
// 页码
func (r *AlibabaXiamiApiArtistAlbumsGetRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaXiamiApiArtistAlbumsGetRequest) GetPage() int64 {
    return r._page
}
// Id Setter
// 艺人id
func (r *AlibabaXiamiApiArtistAlbumsGetRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiArtistAlbumsGetRequest) GetId() int64 {
    return r._id
}

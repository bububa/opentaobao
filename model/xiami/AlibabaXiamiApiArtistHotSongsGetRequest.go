package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
热门歌曲 API请求
alibaba.xiami.api.artist.hotSongs.get

热门歌曲
*/
type AlibabaXiamiApiArtistHotSongsGetRequest struct {
    model.Params
    // 艺人id
    id   int64
}

// 初始化AlibabaXiamiApiArtistHotSongsGetRequest对象
func NewAlibabaXiamiApiArtistHotSongsGetRequest() *AlibabaXiamiApiArtistHotSongsGetRequest{
    return &AlibabaXiamiApiArtistHotSongsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiArtistHotSongsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.hotSongs.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiArtistHotSongsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 艺人id
func (r *AlibabaXiamiApiArtistHotSongsGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaXiamiApiArtistHotSongsGetRequest) GetId() int64 {
    return r.id
}

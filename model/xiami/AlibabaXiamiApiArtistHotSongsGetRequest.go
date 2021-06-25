package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
热门歌曲 APIRequest
alibaba.xiami.api.artist.hotSongs.get

热门歌曲
*/
type AlibabaXiamiApiArtistHotSongsGetRequest struct {
    model.Params

    // 艺人id
    id   int64 

}

func NewAlibabaXiamiApiArtistHotSongsGetRequest() *AlibabaXiamiApiArtistHotSongsGetRequest{
    return &AlibabaXiamiApiArtistHotSongsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiArtistHotSongsGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.artist.hotSongs.get"
}

func (r AlibabaXiamiApiArtistHotSongsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiArtistHotSongsGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaXiamiApiArtistHotSongsGetRequest) GetId() int64 {
    return r.id
}


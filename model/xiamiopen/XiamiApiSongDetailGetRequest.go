package xiamiopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲详情 APIRequest
xiami.api.song.detail.get

获取歌曲详情
*/
type XiamiApiSongDetailGetRequest struct {
    model.Params

    // 歌曲id
    songIds   []int64 

}

func NewXiamiApiSongDetailGetRequest() *XiamiApiSongDetailGetRequest{
    return &XiamiApiSongDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiApiSongDetailGetRequest) GetApiMethodName() string {
    return "xiami.api.song.detail.get"
}

func (r XiamiApiSongDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiApiSongDetailGetRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

func (r XiamiApiSongDetailGetRequest) GetSongIds() []int64 {
    return r.songIds
}


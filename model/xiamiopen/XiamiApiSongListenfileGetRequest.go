package xiamiopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲试听文件 APIRequest
xiami.api.song.listenfile.get

获取歌曲试听文件
*/
type XiamiApiSongListenfileGetRequest struct {
    model.Params

    // 歌曲id
    songIds   []int64 

}

func NewXiamiApiSongListenfileGetRequest() *XiamiApiSongListenfileGetRequest{
    return &XiamiApiSongListenfileGetRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiApiSongListenfileGetRequest) GetApiMethodName() string {
    return "xiami.api.song.listenfile.get"
}

func (r XiamiApiSongListenfileGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiApiSongListenfileGetRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

func (r XiamiApiSongListenfileGetRequest) GetSongIds() []int64 {
    return r.songIds
}


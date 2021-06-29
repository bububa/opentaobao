package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲信息 APIRequest
xiami.content.songs.info.get

(批量)获取歌曲信息
*/
type XiamiContentSongsInfoGetRequest struct {
    model.Params

    // 歌曲ID
    songIds   []int64 

}

func NewXiamiContentSongsInfoGetRequest() *XiamiContentSongsInfoGetRequest{
    return &XiamiContentSongsInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentSongsInfoGetRequest) GetApiMethodName() string {
    return "xiami.content.songs.info.get"
}

func (r XiamiContentSongsInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentSongsInfoGetRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

func (r XiamiContentSongsInfoGetRequest) GetSongIds() []int64 {
    return r.songIds
}


package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲音频 APIRequest
xiami.content.songs.audio.get

获取歌曲音频
*/
type XiamiContentSongsAudioGetRequest struct {
    model.Params

    // 歌曲ID
    songIds   []int64 

}

func NewXiamiContentSongsAudioGetRequest() *XiamiContentSongsAudioGetRequest{
    return &XiamiContentSongsAudioGetRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentSongsAudioGetRequest) GetApiMethodName() string {
    return "xiami.content.songs.audio.get"
}

func (r XiamiContentSongsAudioGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentSongsAudioGetRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

func (r XiamiContentSongsAudioGetRequest) GetSongIds() []int64 {
    return r.songIds
}


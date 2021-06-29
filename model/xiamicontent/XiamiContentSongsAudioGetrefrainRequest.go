package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取副歌信息 APIRequest
xiami.content.songs.audio.getrefrain

获取歌曲音频副歌
*/
type XiamiContentSongsAudioGetrefrainRequest struct {
    model.Params

    // 歌曲ID
    songIds   []int64 

}

func NewXiamiContentSongsAudioGetrefrainRequest() *XiamiContentSongsAudioGetrefrainRequest{
    return &XiamiContentSongsAudioGetrefrainRequest{
        Params: model.NewParams(),
    }
}

func (r XiamiContentSongsAudioGetrefrainRequest) GetApiMethodName() string {
    return "xiami.content.songs.audio.getrefrain"
}

func (r XiamiContentSongsAudioGetrefrainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *XiamiContentSongsAudioGetrefrainRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

func (r XiamiContentSongsAudioGetrefrainRequest) GetSongIds() []int64 {
    return r.songIds
}


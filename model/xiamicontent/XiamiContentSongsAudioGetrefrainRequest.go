package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取副歌信息 API请求
xiami.content.songs.audio.getrefrain

获取歌曲音频副歌
*/
type XiamiContentSongsAudioGetrefrainRequest struct {
    model.Params
    // 歌曲ID
    songIds   []int64
}

// 初始化XiamiContentSongsAudioGetrefrainRequest对象
func NewXiamiContentSongsAudioGetrefrainRequest() *XiamiContentSongsAudioGetrefrainRequest{
    return &XiamiContentSongsAudioGetrefrainRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentSongsAudioGetrefrainRequest) GetApiMethodName() string {
    return "xiami.content.songs.audio.getrefrain"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentSongsAudioGetrefrainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsAudioGetrefrainRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

// SongIds Getter
func (r XiamiContentSongsAudioGetrefrainRequest) GetSongIds() []int64 {
    return r.songIds
}

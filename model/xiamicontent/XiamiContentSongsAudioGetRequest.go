package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲音频 API请求
xiami.content.songs.audio.get

获取歌曲音频
*/
type XiamiContentSongsAudioGetRequest struct {
    model.Params
    // 歌曲ID
    _songIds   []int64
}

// 初始化XiamiContentSongsAudioGetRequest对象
func NewXiamiContentSongsAudioGetRequest() *XiamiContentSongsAudioGetRequest{
    return &XiamiContentSongsAudioGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentSongsAudioGetRequest) GetApiMethodName() string {
    return "xiami.content.songs.audio.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentSongsAudioGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsAudioGetRequest) SetSongIds(_songIds []int64) error {
    r._songIds = _songIds
    r.Set("song_ids", _songIds)
    return nil
}

// SongIds Getter
func (r XiamiContentSongsAudioGetRequest) GetSongIds() []int64 {
    return r._songIds
}

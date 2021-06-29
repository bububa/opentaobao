package xiamiopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲试听文件 API请求
xiami.api.song.listenfile.get

获取歌曲试听文件
*/
type XiamiApiSongListenfileGetRequest struct {
    model.Params
    // 歌曲id
    songIds   []int64
}

// 初始化XiamiApiSongListenfileGetRequest对象
func NewXiamiApiSongListenfileGetRequest() *XiamiApiSongListenfileGetRequest{
    return &XiamiApiSongListenfileGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiApiSongListenfileGetRequest) GetApiMethodName() string {
    return "xiami.api.song.listenfile.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiApiSongListenfileGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SongIds Setter
// 歌曲id
func (r *XiamiApiSongListenfileGetRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

// SongIds Getter
func (r XiamiApiSongListenfileGetRequest) GetSongIds() []int64 {
    return r.songIds
}

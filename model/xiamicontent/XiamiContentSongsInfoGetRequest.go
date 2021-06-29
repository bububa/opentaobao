package xiamicontent

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲信息 API请求
xiami.content.songs.info.get

(批量)获取歌曲信息
*/
type XiamiContentSongsInfoGetRequest struct {
    model.Params
    // 歌曲ID
    songIds   []int64
}

// 初始化XiamiContentSongsInfoGetRequest对象
func NewXiamiContentSongsInfoGetRequest() *XiamiContentSongsInfoGetRequest{
    return &XiamiContentSongsInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentSongsInfoGetRequest) GetApiMethodName() string {
    return "xiami.content.songs.info.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentSongsInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsInfoGetRequest) SetSongIds(songIds []int64) error {
    r.songIds = songIds
    r.Set("song_ids", songIds)
    return nil
}

// SongIds Getter
func (r XiamiContentSongsInfoGetRequest) GetSongIds() []int64 {
    return r.songIds
}

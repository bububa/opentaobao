package xiamiopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取歌曲详情 API请求
xiami.api.song.detail.get

获取歌曲详情
*/
type XiamiApiSongDetailGetRequest struct {
    model.Params
    // 歌曲id
    _songIds   []int64
}

// 初始化XiamiApiSongDetailGetRequest对象
func NewXiamiApiSongDetailGetRequest() *XiamiApiSongDetailGetRequest{
    return &XiamiApiSongDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiApiSongDetailGetRequest) GetApiMethodName() string {
    return "xiami.api.song.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiApiSongDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SongIds Setter
// 歌曲id
func (r *XiamiApiSongDetailGetRequest) SetSongIds(_songIds []int64) error {
    r._songIds = _songIds
    r.Set("song_ids", _songIds)
    return nil
}

// SongIds Getter
func (r XiamiApiSongDetailGetRequest) GetSongIds() []int64 {
    return r._songIds
}

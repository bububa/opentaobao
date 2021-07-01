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
type XiamiContentSongsInfoGetAPIRequest struct {
    model.Params
    // 歌曲ID
    _songIds   []int64
}

// 初始化XiamiContentSongsInfoGetAPIRequest对象
func NewXiamiContentSongsInfoGetRequest() *XiamiContentSongsInfoGetAPIRequest{
    return &XiamiContentSongsInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r XiamiContentSongsInfoGetAPIRequest) GetApiMethodName() string {
    return "xiami.content.songs.info.get"
}

// IRequest interface 方法, 获取API参数
func (r XiamiContentSongsInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsInfoGetAPIRequest) SetSongIds(_songIds []int64) error {
    r._songIds = _songIds
    r.Set("song_ids", _songIds)
    return nil
}

// SongIds Getter
func (r XiamiContentSongsInfoGetAPIRequest) GetSongIds() []int64 {
    return r._songIds
}

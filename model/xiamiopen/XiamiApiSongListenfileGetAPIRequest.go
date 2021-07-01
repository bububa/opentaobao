package xiamiopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiApiSongListenfileGetAPIRequest
获取歌曲试听文件 API请求
xiami.api.song.listenfile.get

获取歌曲试听文件 */
type XiamiApiSongListenfileGetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// NewXiamiApiSongListenfileGetRequest 初始化XiamiApiSongListenfileGetAPIRequest对象
func NewXiamiApiSongListenfileGetRequest() *XiamiApiSongListenfileGetAPIRequest {
	return &XiamiApiSongListenfileGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiApiSongListenfileGetAPIRequest) GetApiMethodName() string {
	return "xiami.api.song.listenfile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiApiSongListenfileGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SongIds Setter
// 歌曲id
func (r *XiamiApiSongListenfileGetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// Get SongIds Getter
func (r XiamiApiSongListenfileGetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

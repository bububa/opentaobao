package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentSongsAudioGetrefrainAPIRequest
获取副歌信息 API请求
xiami.content.songs.audio.getrefrain

获取歌曲音频副歌 */
type XiamiContentSongsAudioGetrefrainAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamiContentSongsAudioGetrefrainRequest 初始化XiamiContentSongsAudioGetrefrainAPIRequest对象
func NewXiamiContentSongsAudioGetrefrainRequest() *XiamiContentSongsAudioGetrefrainAPIRequest {
	return &XiamiContentSongsAudioGetrefrainAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.audio.getrefrain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsAudioGetrefrainAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// Get SongIds Getter
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

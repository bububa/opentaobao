package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsAudioGetAPIRequest 获取歌曲音频 API请求
// xiami.content.songs.audio.get
//
// 获取歌曲音频
type XiamiContentSongsAudioGetAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamiContentSongsAudioGetRequest 初始化XiamiContentSongsAudioGetAPIRequest对象
func NewXiamiContentSongsAudioGetRequest() *XiamiContentSongsAudioGetAPIRequest {
	return &XiamiContentSongsAudioGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentSongsAudioGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.audio.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentSongsAudioGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsAudioGetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// Get SongIds Getter
func (r XiamiContentSongsAudioGetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

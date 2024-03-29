package xiamicontent

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiContentSongsAudioGetAPIRequest) Reset() {
	r._songIds = r._songIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentSongsAudioGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.audio.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentSongsAudioGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentSongsAudioGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsAudioGetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiContentSongsAudioGetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

var poolXiamiContentSongsAudioGetAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiContentSongsAudioGetRequest()
	},
}

// GetXiamiContentSongsAudioGetRequest 从 sync.Pool 获取 XiamiContentSongsAudioGetAPIRequest
func GetXiamiContentSongsAudioGetAPIRequest() *XiamiContentSongsAudioGetAPIRequest {
	return poolXiamiContentSongsAudioGetAPIRequest.Get().(*XiamiContentSongsAudioGetAPIRequest)
}

// ReleaseXiamiContentSongsAudioGetAPIRequest 将 XiamiContentSongsAudioGetAPIRequest 放入 sync.Pool
func ReleaseXiamiContentSongsAudioGetAPIRequest(v *XiamiContentSongsAudioGetAPIRequest) {
	v.Reset()
	poolXiamiContentSongsAudioGetAPIRequest.Put(v)
}

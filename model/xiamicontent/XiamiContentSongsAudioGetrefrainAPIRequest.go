package xiamicontent

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsAudioGetrefrainAPIRequest 获取副歌信息 API请求
// xiami.content.songs.audio.getrefrain
//
// 获取歌曲音频副歌
type XiamiContentSongsAudioGetrefrainAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamiContentSongsAudioGetrefrainRequest 初始化XiamiContentSongsAudioGetrefrainAPIRequest对象
func NewXiamiContentSongsAudioGetrefrainRequest() *XiamiContentSongsAudioGetrefrainAPIRequest {
	return &XiamiContentSongsAudioGetrefrainAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiContentSongsAudioGetrefrainAPIRequest) Reset() {
	r._songIds = r._songIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.audio.getrefrain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsAudioGetrefrainAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiContentSongsAudioGetrefrainAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

var poolXiamiContentSongsAudioGetrefrainAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiContentSongsAudioGetrefrainRequest()
	},
}

// GetXiamiContentSongsAudioGetrefrainRequest 从 sync.Pool 获取 XiamiContentSongsAudioGetrefrainAPIRequest
func GetXiamiContentSongsAudioGetrefrainAPIRequest() *XiamiContentSongsAudioGetrefrainAPIRequest {
	return poolXiamiContentSongsAudioGetrefrainAPIRequest.Get().(*XiamiContentSongsAudioGetrefrainAPIRequest)
}

// ReleaseXiamiContentSongsAudioGetrefrainAPIRequest 将 XiamiContentSongsAudioGetrefrainAPIRequest 放入 sync.Pool
func ReleaseXiamiContentSongsAudioGetrefrainAPIRequest(v *XiamiContentSongsAudioGetrefrainAPIRequest) {
	v.Reset()
	poolXiamiContentSongsAudioGetrefrainAPIRequest.Put(v)
}

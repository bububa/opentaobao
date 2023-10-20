package xiamicontent

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMusicInfoGetAPIRequest 获取音乐信息 API请求
// xiami.content.music.info.get
//
// (批量)获取歌曲信息
type XiamiContentMusicInfoGetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds int64
}

// NewXiamiContentMusicInfoGetRequest 初始化XiamiContentMusicInfoGetAPIRequest对象
func NewXiamiContentMusicInfoGetRequest() *XiamiContentMusicInfoGetAPIRequest {
	return &XiamiContentMusicInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiContentMusicInfoGetAPIRequest) Reset() {
	r._songIds = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentMusicInfoGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.music.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentMusicInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentMusicInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲id
func (r *XiamiContentMusicInfoGetAPIRequest) SetSongIds(_songIds int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiContentMusicInfoGetAPIRequest) GetSongIds() int64 {
	return r._songIds
}

var poolXiamiContentMusicInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiContentMusicInfoGetRequest()
	},
}

// GetXiamiContentMusicInfoGetRequest 从 sync.Pool 获取 XiamiContentMusicInfoGetAPIRequest
func GetXiamiContentMusicInfoGetAPIRequest() *XiamiContentMusicInfoGetAPIRequest {
	return poolXiamiContentMusicInfoGetAPIRequest.Get().(*XiamiContentMusicInfoGetAPIRequest)
}

// ReleaseXiamiContentMusicInfoGetAPIRequest 将 XiamiContentMusicInfoGetAPIRequest 放入 sync.Pool
func ReleaseXiamiContentMusicInfoGetAPIRequest(v *XiamiContentMusicInfoGetAPIRequest) {
	v.Reset()
	poolXiamiContentMusicInfoGetAPIRequest.Put(v)
}

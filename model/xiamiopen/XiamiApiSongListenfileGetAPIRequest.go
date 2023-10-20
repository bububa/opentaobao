package xiamiopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiApiSongListenfileGetAPIRequest 获取歌曲试听文件 API请求
// xiami.api.song.listenfile.get
//
// 获取歌曲试听文件
type XiamiApiSongListenfileGetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// NewXiamiApiSongListenfileGetRequest 初始化XiamiApiSongListenfileGetAPIRequest对象
func NewXiamiApiSongListenfileGetRequest() *XiamiApiSongListenfileGetAPIRequest {
	return &XiamiApiSongListenfileGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiApiSongListenfileGetAPIRequest) Reset() {
	r._songIds = r._songIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiApiSongListenfileGetAPIRequest) GetApiMethodName() string {
	return "xiami.api.song.listenfile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiApiSongListenfileGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiApiSongListenfileGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲id
func (r *XiamiApiSongListenfileGetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiApiSongListenfileGetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

var poolXiamiApiSongListenfileGetAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiApiSongListenfileGetRequest()
	},
}

// GetXiamiApiSongListenfileGetRequest 从 sync.Pool 获取 XiamiApiSongListenfileGetAPIRequest
func GetXiamiApiSongListenfileGetAPIRequest() *XiamiApiSongListenfileGetAPIRequest {
	return poolXiamiApiSongListenfileGetAPIRequest.Get().(*XiamiApiSongListenfileGetAPIRequest)
}

// ReleaseXiamiApiSongListenfileGetAPIRequest 将 XiamiApiSongListenfileGetAPIRequest 放入 sync.Pool
func ReleaseXiamiApiSongListenfileGetAPIRequest(v *XiamiApiSongListenfileGetAPIRequest) {
	v.Reset()
	poolXiamiApiSongListenfileGetAPIRequest.Put(v)
}

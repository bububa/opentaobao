package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongsaudiogetAPIRequest 获取歌曲音频 API请求
// xiami.content.songs.audio.get
//
// 获取歌曲音频
type XiamicontentsongsaudiogetAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamicontentsongsaudiogetRequest 初始化XiamicontentsongsaudiogetAPIRequest对象
func NewXiamicontentsongsaudiogetRequest() *XiamicontentsongsaudiogetAPIRequest {
	return &XiamicontentsongsaudiogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentsongsaudiogetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.audio.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentsongsaudiogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentsongsaudiogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲ID
func (r *XiamicontentsongsaudiogetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamicontentsongsaudiogetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

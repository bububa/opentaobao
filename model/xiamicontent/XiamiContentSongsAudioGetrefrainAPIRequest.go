package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongsaudiogetrefrainAPIRequest 获取副歌信息 API请求
// xiami.content.songs.audio.getrefrain
//
// 获取歌曲音频副歌
type XiamicontentsongsaudiogetrefrainAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamicontentsongsaudiogetrefrainRequest 初始化XiamicontentsongsaudiogetrefrainAPIRequest对象
func NewXiamicontentsongsaudiogetrefrainRequest() *XiamicontentsongsaudiogetrefrainAPIRequest {
	return &XiamicontentsongsaudiogetrefrainAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentsongsaudiogetrefrainAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.audio.getrefrain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentsongsaudiogetrefrainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentsongsaudiogetrefrainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲ID
func (r *XiamicontentsongsaudiogetrefrainAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamicontentsongsaudiogetrefrainAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

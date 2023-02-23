package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentSongsInfoGetAPIRequest 获取歌曲信息 API请求
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
type XiamiContentSongsInfoGetAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamiContentSongsInfoGetRequest 初始化XiamiContentSongsInfoGetAPIRequest对象
func NewXiamiContentSongsInfoGetRequest() *XiamiContentSongsInfoGetAPIRequest {
	return &XiamiContentSongsInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentSongsInfoGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentSongsInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentSongsInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲ID
func (r *XiamiContentSongsInfoGetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiContentSongsInfoGetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

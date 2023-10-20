package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongsinfogetAPIRequest 获取歌曲信息 API请求
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
type XiamicontentsongsinfogetAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// NewXiamicontentsongsinfogetRequest 初始化XiamicontentsongsinfogetAPIRequest对象
func NewXiamicontentsongsinfogetRequest() *XiamicontentsongsinfogetAPIRequest {
	return &XiamicontentsongsinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentsongsinfogetAPIRequest) GetApiMethodName() string {
	return "xiami.content.songs.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentsongsinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentsongsinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲ID
func (r *XiamicontentsongsinfogetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamicontentsongsinfogetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

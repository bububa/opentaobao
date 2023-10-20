package xiamiopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiapisonglistenfilegetAPIRequest 获取歌曲试听文件 API请求
// xiami.api.song.listenfile.get
//
// 获取歌曲试听文件
type XiamiapisonglistenfilegetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// NewXiamiapisonglistenfilegetRequest 初始化XiamiapisonglistenfilegetAPIRequest对象
func NewXiamiapisonglistenfilegetRequest() *XiamiapisonglistenfilegetAPIRequest {
	return &XiamiapisonglistenfilegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiapisonglistenfilegetAPIRequest) GetApiMethodName() string {
	return "xiami.api.song.listenfile.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiapisonglistenfilegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiapisonglistenfilegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲id
func (r *XiamiapisonglistenfilegetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiapisonglistenfilegetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

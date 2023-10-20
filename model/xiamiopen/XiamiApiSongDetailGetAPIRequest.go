package xiamiopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiapisongdetailgetAPIRequest 获取歌曲详情 API请求
// xiami.api.song.detail.get
//
// 获取歌曲详情
type XiamiapisongdetailgetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// NewXiamiapisongdetailgetRequest 初始化XiamiapisongdetailgetAPIRequest对象
func NewXiamiapisongdetailgetRequest() *XiamiapisongdetailgetAPIRequest {
	return &XiamiapisongdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiapisongdetailgetAPIRequest) GetApiMethodName() string {
	return "xiami.api.song.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiapisongdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiapisongdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲id
func (r *XiamiapisongdetailgetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiapisongdetailgetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentalbuminfogetAPIRequest 获取专辑信息 API请求
// xiami.content.album.info.get
//
// 获取专辑信息
type XiamicontentalbuminfogetAPIRequest struct {
	model.Params
	// 专辑id
	_albumIds int64
}

// NewXiamicontentalbuminfogetRequest 初始化XiamicontentalbuminfogetAPIRequest对象
func NewXiamicontentalbuminfogetRequest() *XiamicontentalbuminfogetAPIRequest {
	return &XiamicontentalbuminfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentalbuminfogetAPIRequest) GetApiMethodName() string {
	return "xiami.content.album.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentalbuminfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentalbuminfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlbumIds is AlbumIds Setter
// 专辑id
func (r *XiamicontentalbuminfogetAPIRequest) SetAlbumIds(_albumIds int64) error {
	r._albumIds = _albumIds
	r.Set("album_ids", _albumIds)
	return nil
}

// GetAlbumIds AlbumIds Getter
func (r XiamicontentalbuminfogetAPIRequest) GetAlbumIds() int64 {
	return r._albumIds
}

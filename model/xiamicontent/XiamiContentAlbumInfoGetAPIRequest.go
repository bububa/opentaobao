package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentAlbumInfoGetAPIRequest 获取专辑信息 API请求
// xiami.content.album.info.get
//
// 获取专辑信息
type XiamiContentAlbumInfoGetAPIRequest struct {
	model.Params
	// 专辑id
	_albumIds int64
}

// NewXiamiContentAlbumInfoGetRequest 初始化XiamiContentAlbumInfoGetAPIRequest对象
func NewXiamiContentAlbumInfoGetRequest() *XiamiContentAlbumInfoGetAPIRequest {
	return &XiamiContentAlbumInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentAlbumInfoGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.album.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentAlbumInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAlbumIds is AlbumIds Setter
// 专辑id
func (r *XiamiContentAlbumInfoGetAPIRequest) SetAlbumIds(_albumIds int64) error {
	r._albumIds = _albumIds
	r.Set("album_ids", _albumIds)
	return nil
}

// GetAlbumIds AlbumIds Getter
func (r XiamiContentAlbumInfoGetAPIRequest) GetAlbumIds() int64 {
	return r._albumIds
}

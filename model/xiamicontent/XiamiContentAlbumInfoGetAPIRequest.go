package xiamicontent

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiContentAlbumInfoGetAPIRequest) Reset() {
	r._albumIds = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentAlbumInfoGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.album.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentAlbumInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiContentAlbumInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolXiamiContentAlbumInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiContentAlbumInfoGetRequest()
	},
}

// GetXiamiContentAlbumInfoGetRequest 从 sync.Pool 获取 XiamiContentAlbumInfoGetAPIRequest
func GetXiamiContentAlbumInfoGetAPIRequest() *XiamiContentAlbumInfoGetAPIRequest {
	return poolXiamiContentAlbumInfoGetAPIRequest.Get().(*XiamiContentAlbumInfoGetAPIRequest)
}

// ReleaseXiamiContentAlbumInfoGetAPIRequest 将 XiamiContentAlbumInfoGetAPIRequest 放入 sync.Pool
func ReleaseXiamiContentAlbumInfoGetAPIRequest(v *XiamiContentAlbumInfoGetAPIRequest) {
	v.Reset()
	poolXiamiContentAlbumInfoGetAPIRequest.Put(v)
}

package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildNodeitemOfflineAPIRequest 少儿大厅类目内容下线接口 API请求
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
type YunosTvpubadminContentChildNodeitemOfflineAPIRequest struct {
	model.Params
	// 类目内容ID
	_contentId int64
}

// NewYunosTvpubadminContentChildNodeitemOfflineRequest 初始化YunosTvpubadminContentChildNodeitemOfflineAPIRequest对象
func NewYunosTvpubadminContentChildNodeitemOfflineRequest() *YunosTvpubadminContentChildNodeitemOfflineAPIRequest {
	return &YunosTvpubadminContentChildNodeitemOfflineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChildNodeitemOfflineAPIRequest) Reset() {
	r._contentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildNodeitemOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.nodeitem.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildNodeitemOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChildNodeitemOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContentId is ContentId Setter
// 类目内容ID
func (r *YunosTvpubadminContentChildNodeitemOfflineAPIRequest) SetContentId(_contentId int64) error {
	r._contentId = _contentId
	r.Set("content_id", _contentId)
	return nil
}

// GetContentId ContentId Getter
func (r YunosTvpubadminContentChildNodeitemOfflineAPIRequest) GetContentId() int64 {
	return r._contentId
}

var poolYunosTvpubadminContentChildNodeitemOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChildNodeitemOfflineRequest()
	},
}

// GetYunosTvpubadminContentChildNodeitemOfflineRequest 从 sync.Pool 获取 YunosTvpubadminContentChildNodeitemOfflineAPIRequest
func GetYunosTvpubadminContentChildNodeitemOfflineAPIRequest() *YunosTvpubadminContentChildNodeitemOfflineAPIRequest {
	return poolYunosTvpubadminContentChildNodeitemOfflineAPIRequest.Get().(*YunosTvpubadminContentChildNodeitemOfflineAPIRequest)
}

// ReleaseYunosTvpubadminContentChildNodeitemOfflineAPIRequest 将 YunosTvpubadminContentChildNodeitemOfflineAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChildNodeitemOfflineAPIRequest(v *YunosTvpubadminContentChildNodeitemOfflineAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChildNodeitemOfflineAPIRequest.Put(v)
}

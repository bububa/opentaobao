package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildNodeitemOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.nodeitem.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildNodeitemOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

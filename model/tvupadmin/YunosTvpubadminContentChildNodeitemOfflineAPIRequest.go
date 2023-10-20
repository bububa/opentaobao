package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchildnodeitemofflineAPIRequest 少儿大厅类目内容下线接口 API请求
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
type YunostvpubadmincontentchildnodeitemofflineAPIRequest struct {
	model.Params
	// 类目内容ID
	_contentId int64
}

// NewYunostvpubadmincontentchildnodeitemofflineRequest 初始化YunostvpubadmincontentchildnodeitemofflineAPIRequest对象
func NewYunostvpubadmincontentchildnodeitemofflineRequest() *YunostvpubadmincontentchildnodeitemofflineAPIRequest {
	return &YunostvpubadmincontentchildnodeitemofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentchildnodeitemofflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.nodeitem.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentchildnodeitemofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentchildnodeitemofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContentId is ContentId Setter
// 类目内容ID
func (r *YunostvpubadmincontentchildnodeitemofflineAPIRequest) SetContentId(_contentId int64) error {
	r._contentId = _contentId
	r.Set("content_id", _contentId)
	return nil
}

// GetContentId ContentId Getter
func (r YunostvpubadmincontentchildnodeitemofflineAPIRequest) GetContentId() int64 {
	return r._contentId
}

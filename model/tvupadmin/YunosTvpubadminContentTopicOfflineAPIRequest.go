package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTopicOfflineAPIRequest 迎客松专题下线 API请求
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
type YunosTvpubadminContentTopicOfflineAPIRequest struct {
	model.Params
	// id
	_id int64
}

// NewYunosTvpubadminContentTopicOfflineRequest 初始化YunosTvpubadminContentTopicOfflineAPIRequest对象
func NewYunosTvpubadminContentTopicOfflineRequest() *YunosTvpubadminContentTopicOfflineAPIRequest {
	return &YunosTvpubadminContentTopicOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTopicOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.topic.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTopicOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentTopicOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// id
func (r *YunosTvpubadminContentTopicOfflineAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminContentTopicOfflineAPIRequest) GetId() int64 {
	return r._id
}

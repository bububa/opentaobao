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
func (r YunosTvpubadminContentTopicOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// id
func (r *YunosTvpubadminContentTopicOfflineAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r YunosTvpubadminContentTopicOfflineAPIRequest) GetId() int64 {
	return r._id
}

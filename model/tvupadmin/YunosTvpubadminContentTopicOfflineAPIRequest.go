package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontenttopicofflineAPIRequest 迎客松专题下线 API请求
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
type YunostvpubadmincontenttopicofflineAPIRequest struct {
	model.Params
	// id
	_id int64
}

// NewYunostvpubadmincontenttopicofflineRequest 初始化YunostvpubadmincontenttopicofflineAPIRequest对象
func NewYunostvpubadmincontenttopicofflineRequest() *YunostvpubadmincontenttopicofflineAPIRequest {
	return &YunostvpubadmincontenttopicofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontenttopicofflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.topic.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontenttopicofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontenttopicofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// id
func (r *YunostvpubadmincontenttopicofflineAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmincontenttopicofflineAPIRequest) GetId() int64 {
	return r._id
}

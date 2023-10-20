package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicFindbyidAPIRequest 根据id获取专题信息 API请求
// yunos.tvpubadmin.manage.topic.findbyid
//
// 根据id获取专题信息
type YunosTvpubadminManageTopicFindbyidAPIRequest struct {
	model.Params
	// 专题id
	_id int64
}

// NewYunosTvpubadminManageTopicFindbyidRequest 初始化YunosTvpubadminManageTopicFindbyidAPIRequest对象
func NewYunosTvpubadminManageTopicFindbyidRequest() *YunosTvpubadminManageTopicFindbyidAPIRequest {
	return &YunosTvpubadminManageTopicFindbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicFindbyidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.findbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicFindbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicFindbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 专题id
func (r *YunosTvpubadminManageTopicFindbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminManageTopicFindbyidAPIRequest) GetId() int64 {
	return r._id
}

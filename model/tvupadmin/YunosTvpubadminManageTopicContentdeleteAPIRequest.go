package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContentdeleteAPIRequest 删除专题下内容 API请求
// yunos.tvpubadmin.manage.topic.contentdelete
//
// 删除专题下内容信息
type YunosTvpubadminManageTopicContentdeleteAPIRequest struct {
	model.Params
	// 节目id
	_id int64
}

// NewYunosTvpubadminManageTopicContentdeleteRequest 初始化YunosTvpubadminManageTopicContentdeleteAPIRequest对象
func NewYunosTvpubadminManageTopicContentdeleteRequest() *YunosTvpubadminManageTopicContentdeleteAPIRequest {
	return &YunosTvpubadminManageTopicContentdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentdeleteAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicContentdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 节目id
func (r *YunosTvpubadminManageTopicContentdeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminManageTopicContentdeleteAPIRequest) GetId() int64 {
	return r._id
}

package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiccontentdeleteAPIRequest 删除专题下内容 API请求
// yunos.tvpubadmin.manage.topic.contentdelete
//
// 删除专题下内容信息
type YunostvpubadminmanagetopiccontentdeleteAPIRequest struct {
	model.Params
	// 节目id
	_id int64
}

// NewYunostvpubadminmanagetopiccontentdeleteRequest 初始化YunostvpubadminmanagetopiccontentdeleteAPIRequest对象
func NewYunostvpubadminmanagetopiccontentdeleteRequest() *YunostvpubadminmanagetopiccontentdeleteAPIRequest {
	return &YunostvpubadminmanagetopiccontentdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopiccontentdeleteAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopiccontentdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopiccontentdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 节目id
func (r *YunostvpubadminmanagetopiccontentdeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadminmanagetopiccontentdeleteAPIRequest) GetId() int64 {
	return r._id
}

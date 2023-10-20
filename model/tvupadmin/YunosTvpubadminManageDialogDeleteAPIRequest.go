package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagedialogdeleteAPIRequest 删除全局弹窗 API请求
// yunos.tvpubadmin.manage.dialog.delete
//
// 删除全局弹窗
type YunostvpubadminmanagedialogdeleteAPIRequest struct {
	model.Params
	// 全局弹窗id
	_id int64
}

// NewYunostvpubadminmanagedialogdeleteRequest 初始化YunostvpubadminmanagedialogdeleteAPIRequest对象
func NewYunostvpubadminmanagedialogdeleteRequest() *YunostvpubadminmanagedialogdeleteAPIRequest {
	return &YunostvpubadminmanagedialogdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagedialogdeleteAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagedialogdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagedialogdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 全局弹窗id
func (r *YunostvpubadminmanagedialogdeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadminmanagedialogdeleteAPIRequest) GetId() int64 {
	return r._id
}

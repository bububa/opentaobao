package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogDeleteAPIRequest 删除全局弹窗 API请求
// yunos.tvpubadmin.manage.dialog.delete
//
// 删除全局弹窗
type YunosTvpubadminManageDialogDeleteAPIRequest struct {
	model.Params
	// 全局弹窗id
	_id int64
}

// NewYunosTvpubadminManageDialogDeleteRequest 初始化YunosTvpubadminManageDialogDeleteAPIRequest对象
func NewYunosTvpubadminManageDialogDeleteRequest() *YunosTvpubadminManageDialogDeleteAPIRequest {
	return &YunosTvpubadminManageDialogDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogDeleteAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageDialogDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 全局弹窗id
func (r *YunosTvpubadminManageDialogDeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminManageDialogDeleteAPIRequest) GetId() int64 {
	return r._id
}

package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogEditAPIRequest 编辑全局弹窗 API请求
// yunos.tvpubadmin.manage.dialog.edit
//
// 编辑全局弹窗
type YunosTvpubadminManageDialogEditAPIRequest struct {
	model.Params
	// 待编辑的全局弹窗
	_dialogJson string
}

// NewYunosTvpubadminManageDialogEditRequest 初始化YunosTvpubadminManageDialogEditAPIRequest对象
func NewYunosTvpubadminManageDialogEditRequest() *YunosTvpubadminManageDialogEditAPIRequest {
	return &YunosTvpubadminManageDialogEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminManageDialogEditAPIRequest) Reset() {
	r._dialogJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogEditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageDialogEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDialogJson is DialogJson Setter
// 待编辑的全局弹窗
func (r *YunosTvpubadminManageDialogEditAPIRequest) SetDialogJson(_dialogJson string) error {
	r._dialogJson = _dialogJson
	r.Set("dialog_json", _dialogJson)
	return nil
}

// GetDialogJson DialogJson Getter
func (r YunosTvpubadminManageDialogEditAPIRequest) GetDialogJson() string {
	return r._dialogJson
}

var poolYunosTvpubadminManageDialogEditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminManageDialogEditRequest()
	},
}

// GetYunosTvpubadminManageDialogEditRequest 从 sync.Pool 获取 YunosTvpubadminManageDialogEditAPIRequest
func GetYunosTvpubadminManageDialogEditAPIRequest() *YunosTvpubadminManageDialogEditAPIRequest {
	return poolYunosTvpubadminManageDialogEditAPIRequest.Get().(*YunosTvpubadminManageDialogEditAPIRequest)
}

// ReleaseYunosTvpubadminManageDialogEditAPIRequest 将 YunosTvpubadminManageDialogEditAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminManageDialogEditAPIRequest(v *YunosTvpubadminManageDialogEditAPIRequest) {
	v.Reset()
	poolYunosTvpubadminManageDialogEditAPIRequest.Put(v)
}

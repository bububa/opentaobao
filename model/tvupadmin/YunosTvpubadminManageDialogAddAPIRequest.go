package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagedialogaddAPIRequest 新增全局弹窗 API请求
// yunos.tvpubadmin.manage.dialog.add
//
// 新增全局弹窗
type YunostvpubadminmanagedialogaddAPIRequest struct {
	model.Params
	// 新增的全局弹窗json
	_dialogJson string
}

// NewYunostvpubadminmanagedialogaddRequest 初始化YunostvpubadminmanagedialogaddAPIRequest对象
func NewYunostvpubadminmanagedialogaddRequest() *YunostvpubadminmanagedialogaddAPIRequest {
	return &YunostvpubadminmanagedialogaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagedialogaddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagedialogaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagedialogaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDialogJson is DialogJson Setter
// 新增的全局弹窗json
func (r *YunostvpubadminmanagedialogaddAPIRequest) SetDialogJson(_dialogJson string) error {
	r._dialogJson = _dialogJson
	r.Set("dialog_json", _dialogJson)
	return nil
}

// GetDialogJson DialogJson Getter
func (r YunostvpubadminmanagedialogaddAPIRequest) GetDialogJson() string {
	return r._dialogJson
}

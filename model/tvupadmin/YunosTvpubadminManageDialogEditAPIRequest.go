package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagedialogeditAPIRequest 编辑全局弹窗 API请求
// yunos.tvpubadmin.manage.dialog.edit
//
// 编辑全局弹窗
type YunostvpubadminmanagedialogeditAPIRequest struct {
	model.Params
	// 待编辑的全局弹窗
	_dialogJson string
}

// NewYunostvpubadminmanagedialogeditRequest 初始化YunostvpubadminmanagedialogeditAPIRequest对象
func NewYunostvpubadminmanagedialogeditRequest() *YunostvpubadminmanagedialogeditAPIRequest {
	return &YunostvpubadminmanagedialogeditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagedialogeditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagedialogeditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagedialogeditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDialogJson is DialogJson Setter
// 待编辑的全局弹窗
func (r *YunostvpubadminmanagedialogeditAPIRequest) SetDialogJson(_dialogJson string) error {
	r._dialogJson = _dialogJson
	r.Set("dialog_json", _dialogJson)
	return nil
}

// GetDialogJson DialogJson Getter
func (r YunostvpubadminmanagedialogeditAPIRequest) GetDialogJson() string {
	return r._dialogJson
}

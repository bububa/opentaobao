package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageDialogEditAPIRequest
编辑全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.edit

编辑全局弹窗 */
type YunosTvpubadminManageDialogEditAPIRequest struct {
	model.Params
	// 待编辑的全局弹窗
	_dialogJson string
}

// New

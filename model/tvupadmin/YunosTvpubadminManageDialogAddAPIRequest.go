package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageDialogAddAPIRequest
新增全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.add

新增全局弹窗 */
type YunosTvpubadminManageDialogAddAPIRequest struct {
	model.Params
	// 新增的全局弹窗json
	_dialogJson string
}

// New

package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageDialogDeleteAPIRequest
删除全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.delete

删除全局弹窗 */
type YunosTvpubadminManageDialogDeleteAPIRequest struct {
	model.Params
	// 全局弹窗id
	_id int64
}

// New

package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageDialogFindbyidAPIRequest
根据id查询全局弹窗 API请求
yunos.tvpubadmin.manage.dialog.findbyid

根据id查询全局弹窗 */
type YunosTvpubadminManageDialogFindbyidAPIRequest struct {
	model.Params
	// 全局弹窗id
	_id int64
}

// New

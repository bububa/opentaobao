package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageDialogListAPIRequest
分页获取弹窗列表 API请求
yunos.tvpubadmin.manage.dialog.list

分页获取弹窗配置列表 */
type YunosTvpubadminManageDialogListAPIRequest struct {
	model.Params
	// 查询的query参数
	_query string
}

// New

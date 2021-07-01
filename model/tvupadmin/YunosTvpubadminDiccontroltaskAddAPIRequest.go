package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDiccontroltaskAddAPIRequest
新增停开服任务 API请求
yunos.tvpubadmin.diccontroltask.add

新增停开服任务 */
type YunosTvpubadminDiccontroltaskAddAPIRequest struct {
	model.Params
	// 任务信息
	_task *DicControlTaskDo
}

// New

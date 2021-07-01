package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDiccontroltaskUpdateAPIRequest
停开服任务状态变更 API请求
yunos.tvpubadmin.diccontroltask.update

停开服任务状态变更 */
type YunosTvpubadminDiccontroltaskUpdateAPIRequest struct {
	model.Params
	// 任务ID
	_id int64
	// 任务状态
	_status int64
	// 牌照方
	_license int64
}

// New

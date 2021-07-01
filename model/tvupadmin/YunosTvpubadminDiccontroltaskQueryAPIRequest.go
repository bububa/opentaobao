package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDiccontroltaskQueryAPIRequest
停开服任务列表 API请求
yunos.tvpubadmin.diccontroltask.query

牌照方对终端设备的停开服管理 */
type YunosTvpubadminDiccontroltaskQueryAPIRequest struct {
	model.Params
	// 任务名称
	_name string
	// 任务状态
	_status int64
	// 牌照方
	_license int64
	// 当前页码值
	_pageNo int64
	// 每页条数
	_pageSize int64
}

// New

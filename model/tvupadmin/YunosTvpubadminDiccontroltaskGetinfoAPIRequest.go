package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDiccontroltaskGetinfoAPIRequest
获取停开服任务详情 API请求
yunos.tvpubadmin.diccontroltask.getinfo

获取停开服任务详情 */
type YunosTvpubadminDiccontroltaskGetinfoAPIRequest struct {
	model.Params
	// 任务ID
	_id int64
	// 牌照方
	_license int64
}

// New

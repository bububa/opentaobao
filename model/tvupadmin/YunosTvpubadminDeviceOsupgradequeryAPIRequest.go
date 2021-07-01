package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceOsupgradequeryAPIRequest
系统升级查询 API请求
yunos.tvpubadmin.device.osupgradequery

系统升级查询 */
type YunosTvpubadminDeviceOsupgradequeryAPIRequest struct {
	model.Params
	// 牌照方
	_license int64
	// 审核状态
	_status string
	// 时间范围
	_dayRange int64
	// 第几页
	_pageNo int64
	// 数据大小
	_pageSize int64
}

// New

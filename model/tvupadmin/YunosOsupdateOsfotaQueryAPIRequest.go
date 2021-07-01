package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateOsfotaQueryAPIRequest
系统升级分页查询 API请求
yunos.osupdate.osfota.query

分页查询osoupdate系统升级列表 */
type YunosOsupdateOsfotaQueryAPIRequest struct {
	model.Params
	// 设备型号ID
	_modleId int64
	// 页码
	_page int64
	// 每页数量
	_pageSize int64
}

// New

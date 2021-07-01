package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateVersionstatusUpdateAPIRequest
更新应用升级状态 API请求
yunos.osupdate.versionstatus.update

更新应用升级状态 */
type YunosOsupdateVersionstatusUpdateAPIRequest struct {
	model.Params
	// 升级任务ID
	_id int64
	// 状态值
	_status string
}

// New

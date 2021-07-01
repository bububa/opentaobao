package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateOsfotaAddAPIRequest
添加系统升级任务 API请求
yunos.osupdate.osfota.add

添加osupdate系统升级任务 */
type YunosOsupdateOsfotaAddAPIRequest struct {
	model.Params
	// 系统升级任务json格式
	_osFotaJson string
}

// New

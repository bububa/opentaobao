package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateOsfotaPublishAPIRequest
系统升级发布 API请求
yunos.osupdate.osfota.publish

发布osupdate系统升级任务 */
type YunosOsupdateOsfotaPublishAPIRequest struct {
	model.Params
	// 入参json格式
	_publishJson string
}

// New

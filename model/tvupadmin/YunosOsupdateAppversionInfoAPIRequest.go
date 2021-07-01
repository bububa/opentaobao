package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionInfoAPIRequest
获取应用升级详情 API请求
yunos.osupdate.appversion.info

获取应用升级详情 */
type YunosOsupdateAppversionInfoAPIRequest struct {
	model.Params
	// 升级任务ID
	_id int64
}

// New

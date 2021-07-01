package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionCreateAPIRequest
创建应用升级任务 API请求
yunos.osupdate.appversion.create

创建应用升级任务 */
type YunosOsupdateAppversionCreateAPIRequest struct {
	model.Params
	// 应用版本信息
	_appVersion *TvAppVersion
}

// New

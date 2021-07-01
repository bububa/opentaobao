package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionUpdateAPIRequest
应用升级任务更新 API请求
yunos.osupdate.appversion.update

应用升级任务更新 */
type YunosOsupdateAppversionUpdateAPIRequest struct {
	model.Params
	// 应用版本升级信息
	_appVersion *TvAppVersion
}

// New

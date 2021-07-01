package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionPublishAPIRequest
发布应用升级 API请求
yunos.osupdate.appversion.publish

发布应用升级任务 */
type YunosOsupdateAppversionPublishAPIRequest struct {
	model.Params
	// 发布应用升级入参json
	_publishJson string
}

// New

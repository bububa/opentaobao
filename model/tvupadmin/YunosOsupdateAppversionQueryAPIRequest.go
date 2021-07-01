package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateAppversionQueryAPIRequest
分页获取桌面升级任务 API请求
yunos.osupdate.appversion.query

分页获取桌面升级任务 */
type YunosOsupdateAppversionQueryAPIRequest struct {
	model.Params
	// 应用ID
	_appId int64
	// 页码值
	_page int64
	// 页大小
	_size int64
}

// New

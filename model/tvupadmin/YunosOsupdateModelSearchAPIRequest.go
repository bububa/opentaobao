package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateModelSearchAPIRequest
机型检索 API请求
yunos.osupdate.model.search

机型检索 */
type YunosOsupdateModelSearchAPIRequest struct {
	model.Params
	// 应用ID
	_appId int64
	// 关键词
	_name string
}

// New

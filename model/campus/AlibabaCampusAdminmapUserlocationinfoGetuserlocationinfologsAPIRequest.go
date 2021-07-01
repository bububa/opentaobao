package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest
分时间段获取用户历史位置信息 API请求
alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs

分时间段获取用户历史位置信息 */
type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfoQuery
}

// New

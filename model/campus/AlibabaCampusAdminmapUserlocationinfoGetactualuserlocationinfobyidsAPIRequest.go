package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest
根据userId(支持单个或批量)获取用户实时位置信息 API请求
alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids

根据userId(支持单个或批量)获取用户实时位置信息
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：getActualUserLocationInfoByIds */
type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfoQuery
}

// New

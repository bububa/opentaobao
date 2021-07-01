package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest
上传用户实时位置 API请求
alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo

上传用户实时位置
HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
HSF方法名称：insertActualUserLocationInfo */
type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfo
}

// New

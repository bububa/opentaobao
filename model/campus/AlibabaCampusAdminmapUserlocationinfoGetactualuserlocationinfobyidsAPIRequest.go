package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest 根据userId(支持单个或批量)获取用户实时位置信息 API请求
// alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids
//
// 根据userId(支持单个或批量)获取用户实时位置信息
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：getActualUserLocationInfoByIds
type AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfoQuery
}

// NewAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest 初始化AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest对象
func NewAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest() *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest {
	return &AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) SetParam1(_param1 *UserLocationInfoQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetParam1() *UserLocationInfoQuery {
	return r._param1
}

package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest 上传用户实时位置 API请求
// alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo
//
// 上传用户实时位置
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：insertActualUserLocationInfo
type AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfo
}

// NewAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest 初始化AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest对象
func NewAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest() *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest {
	return &AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) SetParam1(_param1 *UserLocationInfo) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetParam1() *UserLocationInfo {
	return r._param1
}

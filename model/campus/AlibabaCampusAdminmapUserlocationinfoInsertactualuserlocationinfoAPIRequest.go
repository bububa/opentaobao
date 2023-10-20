package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest 上传用户实时位置 API请求
// alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo
//
// 上传用户实时位置
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：insertActualUserLocationInfo
type AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfo
}

// NewAlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoRequest 初始化AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest对象
func NewAlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoRequest() *AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest {
	return &AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 环境参数
func (r *AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数
func (r *AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) SetParam1(_param1 *UserLocationInfo) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIRequest) GetParam1() *UserLocationInfo {
	return r._param1
}

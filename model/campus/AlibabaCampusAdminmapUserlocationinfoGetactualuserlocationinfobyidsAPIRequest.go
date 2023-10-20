package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) SetParam1(_param1 *UserLocationInfoQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) GetParam1() *UserLocationInfoQuery {
	return r._param1
}

var poolAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest()
	},
}

// GetAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsRequest 从 sync.Pool 获取 AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest
func GetAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest() *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest {
	return poolAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest.Get().(*AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest)
}

// ReleaseAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest 将 AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest(v *AlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest) {
	v.Reset()
	poolAlibabaCampusAdminmapUserlocationinfoGetactualuserlocationinfobyidsAPIRequest.Put(v)
}

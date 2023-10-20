package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) SetParam1(_param1 *UserLocationInfo) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) GetParam1() *UserLocationInfo {
	return r._param1
}

var poolAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest()
	},
}

// GetAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoRequest 从 sync.Pool 获取 AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest
func GetAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest() *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest {
	return poolAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest.Get().(*AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest)
}

// ReleaseAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest 将 AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest(v *AlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest) {
	v.Reset()
	poolAlibabaCampusAdminmapUserlocationinfoInsertactualuserlocationinfoAPIRequest.Put(v)
}

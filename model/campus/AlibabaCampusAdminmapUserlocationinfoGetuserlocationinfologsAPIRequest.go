package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest 分时间段获取用户历史位置信息 API请求
// alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs
//
// 分时间段获取用户历史位置信息
type AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfoQuery
}

// NewAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest 初始化AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest对象
func NewAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest() *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest {
	return &AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 环境参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数
func (r *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) SetParam1(_param1 *UserLocationInfoQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) GetParam1() *UserLocationInfoQuery {
	return r._param1
}

var poolAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest()
	},
}

// GetAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsRequest 从 sync.Pool 获取 AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest
func GetAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest() *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest {
	return poolAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest.Get().(*AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest)
}

// ReleaseAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest 将 AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest(v *AlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest) {
	v.Reset()
	poolAlibabaCampusAdminmapUserlocationinfoGetuserlocationinfologsAPIRequest.Put(v)
}

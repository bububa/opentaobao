package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest 分时间段获取用户历史位置信息 API请求
// alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs
//
// 分时间段获取用户历史位置信息
type AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *UserLocationInfoQuery
}

// NewAlibabacampusadminmapuserlocationinfogetuserlocationinfologsRequest 初始化AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest对象
func NewAlibabacampusadminmapuserlocationinfogetuserlocationinfologsRequest() *AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest {
	return &AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 环境参数
func (r *AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数
func (r *AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) SetParam1(_param1 *UserLocationInfoQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest) GetParam1() *UserLocationInfoQuery {
	return r._param1
}

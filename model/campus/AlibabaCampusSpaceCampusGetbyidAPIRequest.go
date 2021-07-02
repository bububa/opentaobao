package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceCampusGetbyidAPIRequest 根据园区id获取园区信息 API请求
// alibaba.campus.space.campus.getbyid
//
// 根据园区id获取园区信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
// HSF方法名称：getCampusById
type AlibabaCampusSpaceCampusGetbyidAPIRequest struct {
	model.Params
	// 园区ID
	_param0 *WorkBenchContext
	// 园区ID
	_param1 int64
}

// NewAlibabaCampusSpaceCampusGetbyidRequest 初始化AlibabaCampusSpaceCampusGetbyidAPIRequest对象
func NewAlibabaCampusSpaceCampusGetbyidRequest() *AlibabaCampusSpaceCampusGetbyidAPIRequest {
	return &AlibabaCampusSpaceCampusGetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.campus.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 园区ID
func (r *AlibabaCampusSpaceCampusGetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is Param1 Setter
// 园区ID
func (r *AlibabaCampusSpaceCampusGetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}

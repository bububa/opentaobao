package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest 根据分组ID查询相应的空间单元 API请求
// alibaba.campus.space.unit.getlistbygroupid
//
// 根据分组ID查询相应的空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByGroupId
type AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest struct {
	model.Params
	// 分组ID
	_param0 *WorkBenchContext
	// 分组ID
	_param1 int64
}

// NewAlibabaCampusSpaceUnitGetlistbygroupidRequest 初始化AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest对象
func NewAlibabaCampusSpaceUnitGetlistbygroupidRequest() *AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest {
	return &AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlistbygroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 分组ID
func (r *AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 分组ID
func (r *AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest) GetParam1() int64 {
	return r._param1
}

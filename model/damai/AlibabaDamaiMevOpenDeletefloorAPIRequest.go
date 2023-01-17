package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletefloorAPIRequest 大麦换验平台-第三方对外开放-楼层接口deleteFloor API请求
// alibaba.damai.mev.open.deletefloor
//
// deleteFloor
type AlibabaDamaiMevOpenDeletefloorAPIRequest struct {
	model.Params
	// 入参deleteFloorParam
	_deleteFloorParam *FloorIdOpenParam
}

// NewAlibabaDamaiMevOpenDeletefloorRequest 初始化AlibabaDamaiMevOpenDeletefloorAPIRequest对象
func NewAlibabaDamaiMevOpenDeletefloorRequest() *AlibabaDamaiMevOpenDeletefloorAPIRequest {
	return &AlibabaDamaiMevOpenDeletefloorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletefloorAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deletefloor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletefloorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeletefloorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteFloorParam is DeleteFloorParam Setter
// 入参deleteFloorParam
func (r *AlibabaDamaiMevOpenDeletefloorAPIRequest) SetDeleteFloorParam(_deleteFloorParam *FloorIdOpenParam) error {
	r._deleteFloorParam = _deleteFloorParam
	r.Set("delete_floor_param", _deleteFloorParam)
	return nil
}

// GetDeleteFloorParam DeleteFloorParam Getter
func (r AlibabaDamaiMevOpenDeletefloorAPIRequest) GetDeleteFloorParam() *FloorIdOpenParam {
	return r._deleteFloorParam
}

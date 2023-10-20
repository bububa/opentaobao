package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushfloorAPIRequest 大麦换验平台-第三方对外开放-楼层接口pushFloor API请求
// alibaba.damai.mev.open.pushfloor
//
// pushFloor
type AlibabadamaimevopenpushfloorAPIRequest struct {
	model.Params
	// 入参pushFloorParam
	_pushFloorParam *ThirdFloorPushOpenParam
}

// NewAlibabadamaimevopenpushfloorRequest 初始化AlibabadamaimevopenpushfloorAPIRequest对象
func NewAlibabadamaimevopenpushfloorRequest() *AlibabadamaimevopenpushfloorAPIRequest {
	return &AlibabadamaimevopenpushfloorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushfloorAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushfloor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushfloorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushfloorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushFloorParam is PushFloorParam Setter
// 入参pushFloorParam
func (r *AlibabadamaimevopenpushfloorAPIRequest) SetPushFloorParam(_pushFloorParam *ThirdFloorPushOpenParam) error {
	r._pushFloorParam = _pushFloorParam
	r.Set("push_floor_param", _pushFloorParam)
	return nil
}

// GetPushFloorParam PushFloorParam Getter
func (r AlibabadamaimevopenpushfloorAPIRequest) GetPushFloorParam() *ThirdFloorPushOpenParam {
	return r._pushFloorParam
}

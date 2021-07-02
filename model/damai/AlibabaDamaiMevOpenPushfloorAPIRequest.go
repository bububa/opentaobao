package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfloorAPIRequest 大麦换验平台-第三方对外开放-楼层接口pushFloor API请求
// alibaba.damai.mev.open.pushfloor
//
// pushFloor
type AlibabaDamaiMevOpenPushfloorAPIRequest struct {
	model.Params
	// 入参pushFloorParam
	_pushFloorParam *ThirdFloorPushOpenParam
}

// NewAlibabaDamaiMevOpenPushfloorRequest 初始化AlibabaDamaiMevOpenPushfloorAPIRequest对象
func NewAlibabaDamaiMevOpenPushfloorRequest() *AlibabaDamaiMevOpenPushfloorAPIRequest {
	return &AlibabaDamaiMevOpenPushfloorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushfloor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushFloorParam Setter
// 入参pushFloorParam
func (r *AlibabaDamaiMevOpenPushfloorAPIRequest) SetPushFloorParam(_pushFloorParam *ThirdFloorPushOpenParam) error {
	r._pushFloorParam = _pushFloorParam
	r.Set("push_floor_param", _pushFloorParam)
	return nil
}

// Get PushFloorParam Getter
func (r AlibabaDamaiMevOpenPushfloorAPIRequest) GetPushFloorParam() *ThirdFloorPushOpenParam {
	return r._pushFloorParam
}

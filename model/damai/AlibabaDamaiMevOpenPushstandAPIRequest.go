package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushstandAPIRequest 大麦换验平台-第三方对外开放-看台接口pushStand API请求
// alibaba.damai.mev.open.pushstand
//
// pushStand
type AlibabadamaimevopenpushstandAPIRequest struct {
	model.Params
	// 入参pushStandParam
	_pushStandParam *ThirdStandPushOpenParam
}

// NewAlibabadamaimevopenpushstandRequest 初始化AlibabadamaimevopenpushstandAPIRequest对象
func NewAlibabadamaimevopenpushstandRequest() *AlibabadamaimevopenpushstandAPIRequest {
	return &AlibabadamaimevopenpushstandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushstandAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushstand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushstandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushstandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushStandParam is PushStandParam Setter
// 入参pushStandParam
func (r *AlibabadamaimevopenpushstandAPIRequest) SetPushStandParam(_pushStandParam *ThirdStandPushOpenParam) error {
	r._pushStandParam = _pushStandParam
	r.Set("push_stand_param", _pushStandParam)
	return nil
}

// GetPushStandParam PushStandParam Getter
func (r AlibabadamaimevopenpushstandAPIRequest) GetPushStandParam() *ThirdStandPushOpenParam {
	return r._pushStandParam
}

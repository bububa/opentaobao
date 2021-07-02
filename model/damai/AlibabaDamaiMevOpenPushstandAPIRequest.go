package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushstandAPIRequest 大麦换验平台-第三方对外开放-看台接口pushStand API请求
// alibaba.damai.mev.open.pushstand
//
// pushStand
type AlibabaDamaiMevOpenPushstandAPIRequest struct {
	model.Params
	// 入参pushStandParam
	_pushStandParam *ThirdStandPushOpenParam
}

// NewAlibabaDamaiMevOpenPushstandRequest 初始化AlibabaDamaiMevOpenPushstandAPIRequest对象
func NewAlibabaDamaiMevOpenPushstandRequest() *AlibabaDamaiMevOpenPushstandAPIRequest {
	return &AlibabaDamaiMevOpenPushstandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushstand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushStandParam Setter
// 入参pushStandParam
func (r *AlibabaDamaiMevOpenPushstandAPIRequest) SetPushStandParam(_pushStandParam *ThirdStandPushOpenParam) error {
	r._pushStandParam = _pushStandParam
	r.Set("push_stand_param", _pushStandParam)
	return nil
}

// Get PushStandParam Getter
func (r AlibabaDamaiMevOpenPushstandAPIRequest) GetPushStandParam() *ThirdStandPushOpenParam {
	return r._pushStandParam
}

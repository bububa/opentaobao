package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixorderconfirmAPIRequest 大麦-出票 API请求
// alibaba.damai.maitix.order.confirm
//
// 出票
type AlibabadamaimaitixorderconfirmAPIRequest struct {
	model.Params
	// 出票入参
	_param *MoaConfirmOrderParam
}

// NewAlibabadamaimaitixorderconfirmRequest 初始化AlibabadamaimaitixorderconfirmAPIRequest对象
func NewAlibabadamaimaitixorderconfirmRequest() *AlibabadamaimaitixorderconfirmAPIRequest {
	return &AlibabadamaimaitixorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixorderconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 出票入参
func (r *AlibabadamaimaitixorderconfirmAPIRequest) SetParam(_param *MoaConfirmOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixorderconfirmAPIRequest) GetParam() *MoaConfirmOrderParam {
	return r._param
}

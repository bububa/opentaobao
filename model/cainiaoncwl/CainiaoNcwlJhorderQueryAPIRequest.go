package cainiaoncwl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoncwljhorderqueryAPIRequest 农村物流集货单查询接口 API请求
// cainiao.ncwl.jhorder.query
//
// 提供给接入商家，查询农村物流集货单
type CainiaoncwljhorderqueryAPIRequest struct {
	model.Params
	// 1
	_param0 *JhRequest
}

// NewCainiaoncwljhorderqueryRequest 初始化CainiaoncwljhorderqueryAPIRequest对象
func NewCainiaoncwljhorderqueryRequest() *CainiaoncwljhorderqueryAPIRequest {
	return &CainiaoncwljhorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoncwljhorderqueryAPIRequest) GetApiMethodName() string {
	return "cainiao.ncwl.jhorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoncwljhorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoncwljhorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 1
func (r *CainiaoncwljhorderqueryAPIRequest) SetParam0(_param0 *JhRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r CainiaoncwljhorderqueryAPIRequest) GetParam0() *JhRequest {
	return r._param0
}

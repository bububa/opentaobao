package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvpvqueryAPIRequest 查询pv属性 API请求
// alibaba.idle.isv.pv.query
//
// 查询pv属性
type AlibabaidleisvpvqueryAPIRequest struct {
	model.Params
	// 入参对象
	_youpinCpvQry *YoupinCpvQry
}

// NewAlibabaidleisvpvqueryRequest 初始化AlibabaidleisvpvqueryAPIRequest对象
func NewAlibabaidleisvpvqueryRequest() *AlibabaidleisvpvqueryAPIRequest {
	return &AlibabaidleisvpvqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvpvqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.pv.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvpvqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvpvqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetYoupinCpvQry is YoupinCpvQry Setter
// 入参对象
func (r *AlibabaidleisvpvqueryAPIRequest) SetYoupinCpvQry(_youpinCpvQry *YoupinCpvQry) error {
	r._youpinCpvQry = _youpinCpvQry
	r.Set("youpin_cpv_qry", _youpinCpvQry)
	return nil
}

// GetYoupinCpvQry YoupinCpvQry Getter
func (r AlibabaidleisvpvqueryAPIRequest) GetYoupinCpvQry() *YoupinCpvQry {
	return r._youpinCpvQry
}

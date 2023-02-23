package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvPvQueryAPIRequest 查询pv属性 API请求
// alibaba.idle.isv.pv.query
//
// 查询pv属性
type AlibabaIdleIsvPvQueryAPIRequest struct {
	model.Params
	// 入参对象
	_youpinCpvQry *YoupinCpvQry
}

// NewAlibabaIdleIsvPvQueryRequest 初始化AlibabaIdleIsvPvQueryAPIRequest对象
func NewAlibabaIdleIsvPvQueryRequest() *AlibabaIdleIsvPvQueryAPIRequest {
	return &AlibabaIdleIsvPvQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvPvQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.pv.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvPvQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvPvQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetYoupinCpvQry is YoupinCpvQry Setter
// 入参对象
func (r *AlibabaIdleIsvPvQueryAPIRequest) SetYoupinCpvQry(_youpinCpvQry *YoupinCpvQry) error {
	r._youpinCpvQry = _youpinCpvQry
	r.Set("youpin_cpv_qry", _youpinCpvQry)
	return nil
}

// GetYoupinCpvQry YoupinCpvQry Getter
func (r AlibabaIdleIsvPvQueryAPIRequest) GetYoupinCpvQry() *YoupinCpvQry {
	return r._youpinCpvQry
}

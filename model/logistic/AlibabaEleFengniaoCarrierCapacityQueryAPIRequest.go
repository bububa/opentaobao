package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoCarrierCapacityQueryAPIRequest 按照门店查询骑手运力状态查询 API请求
// alibaba.ele.fengniao.carrier.capacity.query
//
// 提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量
type AlibabaEleFengniaoCarrierCapacityQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// NewAlibabaEleFengniaoCarrierCapacityQueryRequest 初始化AlibabaEleFengniaoCarrierCapacityQueryAPIRequest对象
func NewAlibabaEleFengniaoCarrierCapacityQueryRequest() *AlibabaEleFengniaoCarrierCapacityQueryAPIRequest {
	return &AlibabaEleFengniaoCarrierCapacityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCarrierCapacityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.carrier.capacity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCarrierCapacityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoCarrierCapacityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoCarrierCapacityQueryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoCarrierCapacityQueryAPIRequest) GetParam() *Param {
	return r._param
}

package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradecreateAPIRequest 创建订单 API请求
// taobao.openmall.trade.create
//
// 创建Openmall订单
type TaobaoopenmalltradecreateAPIRequest struct {
	model.Params
	// 请求入参
	_paramTopTradeCreateDO *TopTradeCreateDo
}

// NewTaobaoopenmalltradecreateRequest 初始化TaobaoopenmalltradecreateAPIRequest对象
func NewTaobaoopenmalltradecreateRequest() *TaobaoopenmalltradecreateAPIRequest {
	return &TaobaoopenmalltradecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltradecreateAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltradecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltradecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopTradeCreateDO is ParamTopTradeCreateDO Setter
// 请求入参
func (r *TaobaoopenmalltradecreateAPIRequest) SetParamTopTradeCreateDO(_paramTopTradeCreateDO *TopTradeCreateDo) error {
	r._paramTopTradeCreateDO = _paramTopTradeCreateDO
	r.Set("param_top_trade_create_d_o", _paramTopTradeCreateDO)
	return nil
}

// GetParamTopTradeCreateDO ParamTopTradeCreateDO Getter
func (r TaobaoopenmalltradecreateAPIRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
	return r._paramTopTradeCreateDO
}

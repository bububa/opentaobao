package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentuntreatedchangequeryvtwoAPIRequest 火车票代理商接口-查询待处理改签单列表-含鉴权校验 API请求
// taobao.train.agent.untreatedchange.query.vtwo
//
// 火车票代理商接口-查询待处理改签单列表-含鉴权校验
type TaobaotrainagentuntreatedchangequeryvtwoAPIRequest struct {
	model.Params
	// 查询待处理改签列表rq
	_queryUntreatedChangeRq *QueryUntreatedChangeRq
}

// NewTaobaotrainagentuntreatedchangequeryvtwoRequest 初始化TaobaotrainagentuntreatedchangequeryvtwoAPIRequest对象
func NewTaobaotrainagentuntreatedchangequeryvtwoRequest() *TaobaotrainagentuntreatedchangequeryvtwoAPIRequest {
	return &TaobaotrainagentuntreatedchangequeryvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentuntreatedchangequeryvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.untreatedchange.query.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentuntreatedchangequeryvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentuntreatedchangequeryvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryUntreatedChangeRq is QueryUntreatedChangeRq Setter
// 查询待处理改签列表rq
func (r *TaobaotrainagentuntreatedchangequeryvtwoAPIRequest) SetQueryUntreatedChangeRq(_queryUntreatedChangeRq *QueryUntreatedChangeRq) error {
	r._queryUntreatedChangeRq = _queryUntreatedChangeRq
	r.Set("query_untreated_change_rq", _queryUntreatedChangeRq)
	return nil
}

// GetQueryUntreatedChangeRq QueryUntreatedChangeRq Getter
func (r TaobaotrainagentuntreatedchangequeryvtwoAPIRequest) GetQueryUntreatedChangeRq() *QueryUntreatedChangeRq {
	return r._queryUntreatedChangeRq
}

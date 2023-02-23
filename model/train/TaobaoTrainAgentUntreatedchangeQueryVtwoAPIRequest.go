package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest 火车票代理商接口-查询待处理改签单列表-含鉴权校验 API请求
// taobao.train.agent.untreatedchange.query.vtwo
//
// 火车票代理商接口-查询待处理改签单列表-含鉴权校验
type TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest struct {
	model.Params
	// 查询待处理改签列表rq
	_queryUntreatedChangeRq *QueryUntreatedChangeRq
}

// NewTaobaoTrainAgentUntreatedchangeQueryVtwoRequest 初始化TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest对象
func NewTaobaoTrainAgentUntreatedchangeQueryVtwoRequest() *TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest {
	return &TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.untreatedchange.query.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryUntreatedChangeRq is QueryUntreatedChangeRq Setter
// 查询待处理改签列表rq
func (r *TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest) SetQueryUntreatedChangeRq(_queryUntreatedChangeRq *QueryUntreatedChangeRq) error {
	r._queryUntreatedChangeRq = _queryUntreatedChangeRq
	r.Set("query_untreated_change_rq", _queryUntreatedChangeRq)
	return nil
}

// GetQueryUntreatedChangeRq QueryUntreatedChangeRq Getter
func (r TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest) GetQueryUntreatedChangeRq() *QueryUntreatedChangeRq {
	return r._queryUntreatedChangeRq
}

package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验 API请求
// taobao.train.agent.changeorderdetail.query.vtwo
//
// 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
type TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest struct {
	model.Params
	// 查询改签订单详情rq
	_queryChangeDetailRq *QueryChangeDetailRq
}

// NewTaobaoTrainAgentChangeorderdetailQueryVtwoRequest 初始化TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest对象
func NewTaobaoTrainAgentChangeorderdetailQueryVtwoRequest() *TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest {
	return &TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.changeorderdetail.query.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryChangeDetailRq is QueryChangeDetailRq Setter
// 查询改签订单详情rq
func (r *TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest) SetQueryChangeDetailRq(_queryChangeDetailRq *QueryChangeDetailRq) error {
	r._queryChangeDetailRq = _queryChangeDetailRq
	r.Set("query_change_detail_rq", _queryChangeDetailRq)
	return nil
}

// GetQueryChangeDetailRq QueryChangeDetailRq Getter
func (r TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest) GetQueryChangeDetailRq() *QueryChangeDetailRq {
	return r._queryChangeDetailRq
}

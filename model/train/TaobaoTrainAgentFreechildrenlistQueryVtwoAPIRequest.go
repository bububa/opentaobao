package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest 免费儿童列表查询 API请求
// taobao.train.agent.freechildrenlist.query.vtwo
//
// 免费儿童列表查询
type TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest struct {
	model.Params
}

// NewTaobaoTrainAgentFreechildrenlistQueryVtwoRequest 初始化TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest对象
func NewTaobaoTrainAgentFreechildrenlistQueryVtwoRequest() *TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest {
	return &TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.freechildrenlist.query.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

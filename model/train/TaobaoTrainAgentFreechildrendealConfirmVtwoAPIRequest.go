package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest 免费儿童处理 API请求
// taobao.train.agent.freechildrendeal.confirm.vtwo
//
// 免费儿童列表查询
type TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest struct {
	model.Params
	// 免费儿童处理信息
	_freeChildrenTicketDealRq *FreeChildrenTicketDealRq
}

// NewTaobaotrainagentfreechildrendealconfirmvtwoRequest 初始化TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest对象
func NewTaobaotrainagentfreechildrendealconfirmvtwoRequest() *TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest {
	return &TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.freechildrendeal.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFreeChildrenTicketDealRq is FreeChildrenTicketDealRq Setter
// 免费儿童处理信息
func (r *TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest) SetFreeChildrenTicketDealRq(_freeChildrenTicketDealRq *FreeChildrenTicketDealRq) error {
	r._freeChildrenTicketDealRq = _freeChildrenTicketDealRq
	r.Set("free_children_ticket_deal_rq", _freeChildrenTicketDealRq)
	return nil
}

// GetFreeChildrenTicketDealRq FreeChildrenTicketDealRq Getter
func (r TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest) GetFreeChildrenTicketDealRq() *FreeChildrenTicketDealRq {
	return r._freeChildrenTicketDealRq
}

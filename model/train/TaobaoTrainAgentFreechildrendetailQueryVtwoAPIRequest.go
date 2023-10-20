package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest 免费儿童详情 API请求
// taobao.train.agent.freechildrendetail.query.vtwo
//
// 免费儿童列表详情
type TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest struct {
	model.Params
	// 请求参数
	_freeChildrenTicketDetailRq *FreeChildrenTicketDetailRq
}

// NewTaobaoTrainAgentFreechildrendetailQueryVtwoRequest 初始化TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest对象
func NewTaobaoTrainAgentFreechildrendetailQueryVtwoRequest() *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest {
	return &TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) Reset() {
	r._freeChildrenTicketDetailRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.freechildrendetail.query.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFreeChildrenTicketDetailRq is FreeChildrenTicketDetailRq Setter
// 请求参数
func (r *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) SetFreeChildrenTicketDetailRq(_freeChildrenTicketDetailRq *FreeChildrenTicketDetailRq) error {
	r._freeChildrenTicketDetailRq = _freeChildrenTicketDetailRq
	r.Set("free_children_ticket_detail_rq", _freeChildrenTicketDetailRq)
	return nil
}

// GetFreeChildrenTicketDetailRq FreeChildrenTicketDetailRq Getter
func (r TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) GetFreeChildrenTicketDetailRq() *FreeChildrenTicketDetailRq {
	return r._freeChildrenTicketDetailRq
}

var poolTaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentFreechildrendetailQueryVtwoRequest()
	},
}

// GetTaobaoTrainAgentFreechildrendetailQueryVtwoRequest 从 sync.Pool 获取 TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest
func GetTaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest() *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest {
	return poolTaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest.Get().(*TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest)
}

// ReleaseTaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest 将 TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest(v *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest.Put(v)
}

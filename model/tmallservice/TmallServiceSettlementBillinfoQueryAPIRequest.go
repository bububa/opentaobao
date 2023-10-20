package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettlementBillinfoQueryAPIRequest 查询工单结算信息 API请求
// tmall.service.settlement.billinfo.query
//
// 提供给服务商查询工单结算信息，包含结算的分成金额以及结算的收款明细，平台抽佣比例。用于服务商进行对账
type TmallServiceSettlementBillinfoQueryAPIRequest struct {
	model.Params
	// 已经结算的工单ID
	_workcardId int64
}

// NewTmallServiceSettlementBillinfoQueryRequest 初始化TmallServiceSettlementBillinfoQueryAPIRequest对象
func NewTmallServiceSettlementBillinfoQueryRequest() *TmallServiceSettlementBillinfoQueryAPIRequest {
	return &TmallServiceSettlementBillinfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettlementBillinfoQueryAPIRequest) Reset() {
	r._workcardId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettlementBillinfoQueryAPIRequest) GetApiMethodName() string {
	return "tmall.service.settlement.billinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettlementBillinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettlementBillinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 已经结算的工单ID
func (r *TmallServiceSettlementBillinfoQueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServiceSettlementBillinfoQueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

var poolTmallServiceSettlementBillinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettlementBillinfoQueryRequest()
	},
}

// GetTmallServiceSettlementBillinfoQueryRequest 从 sync.Pool 获取 TmallServiceSettlementBillinfoQueryAPIRequest
func GetTmallServiceSettlementBillinfoQueryAPIRequest() *TmallServiceSettlementBillinfoQueryAPIRequest {
	return poolTmallServiceSettlementBillinfoQueryAPIRequest.Get().(*TmallServiceSettlementBillinfoQueryAPIRequest)
}

// ReleaseTmallServiceSettlementBillinfoQueryAPIRequest 将 TmallServiceSettlementBillinfoQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettlementBillinfoQueryAPIRequest(v *TmallServiceSettlementBillinfoQueryAPIRequest) {
	v.Reset()
	poolTmallServiceSettlementBillinfoQueryAPIRequest.Put(v)
}

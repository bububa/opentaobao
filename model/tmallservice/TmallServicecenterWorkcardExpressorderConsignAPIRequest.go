package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExpressorderConsignAPIRequest 物流订单呼叫运力 API请求
// tmall.servicecenter.workcard.expressorder.consign
//
// 天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
type TmallServicecenterWorkcardExpressorderConsignAPIRequest struct {
	model.Params
	// 工单List
	_workcardIdList []int64
	// 真实接单服务商
	_realTpNick string
	// 物流寄件单号（废弃）
	_expressOrderId int64
	// 物流订单号
	_logisticsOrderId int64
}

// NewTmallServicecenterWorkcardExpressorderConsignRequest 初始化TmallServicecenterWorkcardExpressorderConsignAPIRequest对象
func NewTmallServicecenterWorkcardExpressorderConsignRequest() *TmallServicecenterWorkcardExpressorderConsignAPIRequest {
	return &TmallServicecenterWorkcardExpressorderConsignAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) Reset() {
	r._workcardIdList = r._workcardIdList[:0]
	r._realTpNick = ""
	r._expressOrderId = 0
	r._logisticsOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.expressorder.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIdList is WorkcardIdList Setter
// 工单List
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetWorkcardIdList(_workcardIdList []int64) error {
	r._workcardIdList = _workcardIdList
	r.Set("workcard_id_list", _workcardIdList)
	return nil
}

// GetWorkcardIdList WorkcardIdList Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetWorkcardIdList() []int64 {
	return r._workcardIdList
}

// SetRealTpNick is RealTpNick Setter
// 真实接单服务商
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetExpressOrderId is ExpressOrderId Setter
// 物流寄件单号（废弃）
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetExpressOrderId(_expressOrderId int64) error {
	r._expressOrderId = _expressOrderId
	r.Set("express_order_id", _expressOrderId)
	return nil
}

// GetExpressOrderId ExpressOrderId Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetExpressOrderId() int64 {
	return r._expressOrderId
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流订单号
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}

var poolTmallServicecenterWorkcardExpressorderConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardExpressorderConsignRequest()
	},
}

// GetTmallServicecenterWorkcardExpressorderConsignRequest 从 sync.Pool 获取 TmallServicecenterWorkcardExpressorderConsignAPIRequest
func GetTmallServicecenterWorkcardExpressorderConsignAPIRequest() *TmallServicecenterWorkcardExpressorderConsignAPIRequest {
	return poolTmallServicecenterWorkcardExpressorderConsignAPIRequest.Get().(*TmallServicecenterWorkcardExpressorderConsignAPIRequest)
}

// ReleaseTmallServicecenterWorkcardExpressorderConsignAPIRequest 将 TmallServicecenterWorkcardExpressorderConsignAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardExpressorderConsignAPIRequest(v *TmallServicecenterWorkcardExpressorderConsignAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardExpressorderConsignAPIRequest.Put(v)
}

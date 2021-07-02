package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExpressorderConsignAPIRequest 物流订单呼叫运力 API请求
// tmall.servicecenter.workcard.expressorder.consign
//
// 天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
type TmallServicecenterWorkcardExpressorderConsignAPIRequest struct {
	model.Params
	// 物流寄件单号（废弃）
	_expressOrderId int64
	// 工单List
	_workcardIdList []int64
	// 真实接单服务商
	_realTpNick string
	// 物流订单号
	_logisticsOrderId int64
}

// NewTmallServicecenterWorkcardExpressorderConsignRequest 初始化TmallServicecenterWorkcardExpressorderConsignAPIRequest对象
func NewTmallServicecenterWorkcardExpressorderConsignRequest() *TmallServicecenterWorkcardExpressorderConsignAPIRequest {
	return &TmallServicecenterWorkcardExpressorderConsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.expressorder.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExpressOrderId Setter
// 物流寄件单号（废弃）
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetExpressOrderId(_expressOrderId int64) error {
	r._expressOrderId = _expressOrderId
	r.Set("express_order_id", _expressOrderId)
	return nil
}

// Get ExpressOrderId Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetExpressOrderId() int64 {
	return r._expressOrderId
}

// Set is WorkcardIdList Setter
// 工单List
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetWorkcardIdList(_workcardIdList []int64) error {
	r._workcardIdList = _workcardIdList
	r.Set("workcard_id_list", _workcardIdList)
	return nil
}

// Get WorkcardIdList Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetWorkcardIdList() []int64 {
	return r._workcardIdList
}

// Set is RealTpNick Setter
// 真实接单服务商
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// Get RealTpNick Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// Set is LogisticsOrderId Setter
// 物流订单号
func (r *TmallServicecenterWorkcardExpressorderConsignAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// Get LogisticsOrderId Getter
func (r TmallServicecenterWorkcardExpressorderConsignAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}

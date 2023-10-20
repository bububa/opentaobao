package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportthreeplofflineconsignAPIRequest 3PL直邮线下发货 API请求
// taobao.wlb.import.threepl.offline.consign
//
// 菜鸟认证直邮线下发货
type TaobaowlbimportthreeplofflineconsignAPIRequest struct {
	model.Params
	// 资源code
	_resCode string
	// 运单号
	_waybillNo string
	// 交易单号
	_tradeId int64
	// 资源id
	_resId int64
	// 发件人地址库id
	_fromId int64
}

// NewTaobaowlbimportthreeplofflineconsignRequest 初始化TaobaowlbimportthreeplofflineconsignAPIRequest对象
func NewTaobaowlbimportthreeplofflineconsignRequest() *TaobaowlbimportthreeplofflineconsignAPIRequest {
	return &TaobaowlbimportthreeplofflineconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.import.threepl.offline.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResCode is ResCode Setter
// 资源code
func (r *TaobaowlbimportthreeplofflineconsignAPIRequest) SetResCode(_resCode string) error {
	r._resCode = _resCode
	r.Set("res_code", _resCode)
	return nil
}

// GetResCode ResCode Getter
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetResCode() string {
	return r._resCode
}

// SetWaybillNo is WaybillNo Setter
// 运单号
func (r *TaobaowlbimportthreeplofflineconsignAPIRequest) SetWaybillNo(_waybillNo string) error {
	r._waybillNo = _waybillNo
	r.Set("waybill_no", _waybillNo)
	return nil
}

// GetWaybillNo WaybillNo Getter
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetWaybillNo() string {
	return r._waybillNo
}

// SetTradeId is TradeId Setter
// 交易单号
func (r *TaobaowlbimportthreeplofflineconsignAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// SetResId is ResId Setter
// 资源id
func (r *TaobaowlbimportthreeplofflineconsignAPIRequest) SetResId(_resId int64) error {
	r._resId = _resId
	r.Set("res_id", _resId)
	return nil
}

// GetResId ResId Getter
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetResId() int64 {
	return r._resId
}

// SetFromId is FromId Setter
// 发件人地址库id
func (r *TaobaowlbimportthreeplofflineconsignAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r TaobaowlbimportthreeplofflineconsignAPIRequest) GetFromId() int64 {
	return r._fromId
}

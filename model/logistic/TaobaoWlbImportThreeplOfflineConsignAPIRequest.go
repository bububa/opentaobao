package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportThreeplOfflineConsignAPIRequest 3PL直邮线下发货 API请求
// taobao.wlb.import.threepl.offline.consign
//
// 菜鸟认证直邮线下发货
type TaobaoWlbImportThreeplOfflineConsignAPIRequest struct {
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

// NewTaobaoWlbImportThreeplOfflineConsignRequest 初始化TaobaoWlbImportThreeplOfflineConsignAPIRequest对象
func NewTaobaoWlbImportThreeplOfflineConsignRequest() *TaobaoWlbImportThreeplOfflineConsignAPIRequest {
	return &TaobaoWlbImportThreeplOfflineConsignAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) Reset() {
	r._resCode = ""
	r._waybillNo = ""
	r._tradeId = 0
	r._resId = 0
	r._fromId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.import.threepl.offline.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResCode is ResCode Setter
// 资源code
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetResCode(_resCode string) error {
	r._resCode = _resCode
	r.Set("res_code", _resCode)
	return nil
}

// GetResCode ResCode Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetResCode() string {
	return r._resCode
}

// SetWaybillNo is WaybillNo Setter
// 运单号
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetWaybillNo(_waybillNo string) error {
	r._waybillNo = _waybillNo
	r.Set("waybill_no", _waybillNo)
	return nil
}

// GetWaybillNo WaybillNo Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetWaybillNo() string {
	return r._waybillNo
}

// SetTradeId is TradeId Setter
// 交易单号
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetTradeId() int64 {
	return r._tradeId
}

// SetResId is ResId Setter
// 资源id
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetResId(_resId int64) error {
	r._resId = _resId
	r.Set("res_id", _resId)
	return nil
}

// GetResId ResId Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetResId() int64 {
	return r._resId
}

// SetFromId is FromId Setter
// 发件人地址库id
func (r *TaobaoWlbImportThreeplOfflineConsignAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r TaobaoWlbImportThreeplOfflineConsignAPIRequest) GetFromId() int64 {
	return r._fromId
}

var poolTaobaoWlbImportThreeplOfflineConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbImportThreeplOfflineConsignRequest()
	},
}

// GetTaobaoWlbImportThreeplOfflineConsignRequest 从 sync.Pool 获取 TaobaoWlbImportThreeplOfflineConsignAPIRequest
func GetTaobaoWlbImportThreeplOfflineConsignAPIRequest() *TaobaoWlbImportThreeplOfflineConsignAPIRequest {
	return poolTaobaoWlbImportThreeplOfflineConsignAPIRequest.Get().(*TaobaoWlbImportThreeplOfflineConsignAPIRequest)
}

// ReleaseTaobaoWlbImportThreeplOfflineConsignAPIRequest 将 TaobaoWlbImportThreeplOfflineConsignAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbImportThreeplOfflineConsignAPIRequest(v *TaobaoWlbImportThreeplOfflineConsignAPIRequest) {
	v.Reset()
	poolTaobaoWlbImportThreeplOfflineConsignAPIRequest.Put(v)
}

package exchange

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeConfirmConsignAPIRequest 换货商家确认收货并发货 API请求
// tmall.exchange.confirm.consign
//
// 卖家确认收货并发货
type TmallExchangeConfirmConsignAPIRequest struct {
	model.Params
	// 返回字段
	_fields []string
	// 卖家发货的物流单号
	_logisticsNo string
	// 卖家发货的快递公司
	_logisticsCompanyName string
	// 换货单号ID
	_disputeId int64
	// 卖家发货的物流类型，100表示平邮，200表示快递
	_logisticsType int64
}

// NewTmallExchangeConfirmConsignRequest 初始化TmallExchangeConfirmConsignAPIRequest对象
func NewTmallExchangeConfirmConsignRequest() *TmallExchangeConfirmConsignAPIRequest {
	return &TmallExchangeConfirmConsignAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallExchangeConfirmConsignAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._logisticsNo = ""
	r._logisticsCompanyName = ""
	r._disputeId = 0
	r._logisticsType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeConfirmConsignAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.confirm.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeConfirmConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallExchangeConfirmConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段
func (r *TmallExchangeConfirmConsignAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeConfirmConsignAPIRequest) GetFields() []string {
	return r._fields
}

// SetLogisticsNo is LogisticsNo Setter
// 卖家发货的物流单号
func (r *TmallExchangeConfirmConsignAPIRequest) SetLogisticsNo(_logisticsNo string) error {
	r._logisticsNo = _logisticsNo
	r.Set("logistics_no", _logisticsNo)
	return nil
}

// GetLogisticsNo LogisticsNo Getter
func (r TmallExchangeConfirmConsignAPIRequest) GetLogisticsNo() string {
	return r._logisticsNo
}

// SetLogisticsCompanyName is LogisticsCompanyName Setter
// 卖家发货的快递公司
func (r *TmallExchangeConfirmConsignAPIRequest) SetLogisticsCompanyName(_logisticsCompanyName string) error {
	r._logisticsCompanyName = _logisticsCompanyName
	r.Set("logistics_company_name", _logisticsCompanyName)
	return nil
}

// GetLogisticsCompanyName LogisticsCompanyName Getter
func (r TmallExchangeConfirmConsignAPIRequest) GetLogisticsCompanyName() string {
	return r._logisticsCompanyName
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeConfirmConsignAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeConfirmConsignAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetLogisticsType is LogisticsType Setter
// 卖家发货的物流类型，100表示平邮，200表示快递
func (r *TmallExchangeConfirmConsignAPIRequest) SetLogisticsType(_logisticsType int64) error {
	r._logisticsType = _logisticsType
	r.Set("logistics_type", _logisticsType)
	return nil
}

// GetLogisticsType LogisticsType Getter
func (r TmallExchangeConfirmConsignAPIRequest) GetLogisticsType() int64 {
	return r._logisticsType
}

var poolTmallExchangeConfirmConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallExchangeConfirmConsignRequest()
	},
}

// GetTmallExchangeConfirmConsignRequest 从 sync.Pool 获取 TmallExchangeConfirmConsignAPIRequest
func GetTmallExchangeConfirmConsignAPIRequest() *TmallExchangeConfirmConsignAPIRequest {
	return poolTmallExchangeConfirmConsignAPIRequest.Get().(*TmallExchangeConfirmConsignAPIRequest)
}

// ReleaseTmallExchangeConfirmConsignAPIRequest 将 TmallExchangeConfirmConsignAPIRequest 放入 sync.Pool
func ReleaseTmallExchangeConfirmConsignAPIRequest(v *TmallExchangeConfirmConsignAPIRequest) {
	v.Reset()
	poolTmallExchangeConfirmConsignAPIRequest.Put(v)
}

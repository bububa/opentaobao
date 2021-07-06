package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeConsigngoodsAPIRequest 卖家发货 API请求
// tmall.exchange.consigngoods
//
// 卖家发货
type TmallExchangeConsigngoodsAPIRequest struct {
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

// NewTmallExchangeConsigngoodsRequest 初始化TmallExchangeConsigngoodsAPIRequest对象
func NewTmallExchangeConsigngoodsRequest() *TmallExchangeConsigngoodsAPIRequest {
	return &TmallExchangeConsigngoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeConsigngoodsAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.consigngoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeConsigngoodsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 返回字段
func (r *TmallExchangeConsigngoodsAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeConsigngoodsAPIRequest) GetFields() []string {
	return r._fields
}

// SetLogisticsNo is LogisticsNo Setter
// 卖家发货的物流单号
func (r *TmallExchangeConsigngoodsAPIRequest) SetLogisticsNo(_logisticsNo string) error {
	r._logisticsNo = _logisticsNo
	r.Set("logistics_no", _logisticsNo)
	return nil
}

// GetLogisticsNo LogisticsNo Getter
func (r TmallExchangeConsigngoodsAPIRequest) GetLogisticsNo() string {
	return r._logisticsNo
}

// SetLogisticsCompanyName is LogisticsCompanyName Setter
// 卖家发货的快递公司
func (r *TmallExchangeConsigngoodsAPIRequest) SetLogisticsCompanyName(_logisticsCompanyName string) error {
	r._logisticsCompanyName = _logisticsCompanyName
	r.Set("logistics_company_name", _logisticsCompanyName)
	return nil
}

// GetLogisticsCompanyName LogisticsCompanyName Getter
func (r TmallExchangeConsigngoodsAPIRequest) GetLogisticsCompanyName() string {
	return r._logisticsCompanyName
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeConsigngoodsAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeConsigngoodsAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetLogisticsType is LogisticsType Setter
// 卖家发货的物流类型，100表示平邮，200表示快递
func (r *TmallExchangeConsigngoodsAPIRequest) SetLogisticsType(_logisticsType int64) error {
	r._logisticsType = _logisticsType
	r.Set("logistics_type", _logisticsType)
	return nil
}

// GetLogisticsType LogisticsType Getter
func (r TmallExchangeConsigngoodsAPIRequest) GetLogisticsType() int64 {
	return r._logisticsType
}

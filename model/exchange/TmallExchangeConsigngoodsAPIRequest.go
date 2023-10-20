package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangeconsigngoodsAPIRequest 卖家发货 API请求
// tmall.exchange.consigngoods
//
// 卖家发货
type TmallexchangeconsigngoodsAPIRequest struct {
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

// NewTmallexchangeconsigngoodsRequest 初始化TmallexchangeconsigngoodsAPIRequest对象
func NewTmallexchangeconsigngoodsRequest() *TmallexchangeconsigngoodsAPIRequest {
	return &TmallexchangeconsigngoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangeconsigngoodsAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.consigngoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangeconsigngoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangeconsigngoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段
func (r *TmallexchangeconsigngoodsAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangeconsigngoodsAPIRequest) GetFields() []string {
	return r._fields
}

// SetLogisticsNo is LogisticsNo Setter
// 卖家发货的物流单号
func (r *TmallexchangeconsigngoodsAPIRequest) SetLogisticsNo(_logisticsNo string) error {
	r._logisticsNo = _logisticsNo
	r.Set("logistics_no", _logisticsNo)
	return nil
}

// GetLogisticsNo LogisticsNo Getter
func (r TmallexchangeconsigngoodsAPIRequest) GetLogisticsNo() string {
	return r._logisticsNo
}

// SetLogisticsCompanyName is LogisticsCompanyName Setter
// 卖家发货的快递公司
func (r *TmallexchangeconsigngoodsAPIRequest) SetLogisticsCompanyName(_logisticsCompanyName string) error {
	r._logisticsCompanyName = _logisticsCompanyName
	r.Set("logistics_company_name", _logisticsCompanyName)
	return nil
}

// GetLogisticsCompanyName LogisticsCompanyName Getter
func (r TmallexchangeconsigngoodsAPIRequest) GetLogisticsCompanyName() string {
	return r._logisticsCompanyName
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangeconsigngoodsAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangeconsigngoodsAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetLogisticsType is LogisticsType Setter
// 卖家发货的物流类型，100表示平邮，200表示快递
func (r *TmallexchangeconsigngoodsAPIRequest) SetLogisticsType(_logisticsType int64) error {
	r._logisticsType = _logisticsType
	r.Set("logistics_type", _logisticsType)
	return nil
}

// GetLogisticsType LogisticsType Getter
func (r TmallexchangeconsigngoodsAPIRequest) GetLogisticsType() int64 {
	return r._logisticsType
}

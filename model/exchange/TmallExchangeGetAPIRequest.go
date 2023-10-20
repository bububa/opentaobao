package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangegetAPIRequest 获取单笔换货详情 API请求
// tmall.exchange.get
//
// 获取单笔换货详情
type TmallexchangegetAPIRequest struct {
	model.Params
	// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
	_fields []string
	// 换货单号ID
	_disputeId int64
}

// NewTmallexchangegetRequest 初始化TmallexchangegetAPIRequest对象
func NewTmallexchangegetRequest() *TmallexchangegetAPIRequest {
	return &TmallexchangegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangegetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
func (r *TmallexchangegetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangegetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangegetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangegetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

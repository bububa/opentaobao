package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeGetAPIRequest 获取单笔换货详情 API请求
// tmall.exchange.get
//
// 获取单笔换货详情
type TmallExchangeGetAPIRequest struct {
	model.Params
	// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
	_fields []string
	// 换货单号ID
	_disputeId int64
}

// NewTmallExchangeGetRequest 初始化TmallExchangeGetAPIRequest对象
func NewTmallExchangeGetRequest() *TmallExchangeGetAPIRequest {
	return &TmallExchangeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeGetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
func (r *TmallExchangeGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeGetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeGetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

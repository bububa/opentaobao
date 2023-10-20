package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinnoticeordercancelAPIRequest 福州疫苗取消预约 API请求
// alibaba.health.vaccin.notice.order.cancel
//
// 福州疫苗用户取消预约接口
type AlibabahealthvaccinnoticeordercancelAPIRequest struct {
	model.Params
	// 支付宝用户id
	_alipayUserId string
	// 预约id
	_orderId string
}

// NewAlibabahealthvaccinnoticeordercancelRequest 初始化AlibabahealthvaccinnoticeordercancelAPIRequest对象
func NewAlibabahealthvaccinnoticeordercancelRequest() *AlibabahealthvaccinnoticeordercancelAPIRequest {
	return &AlibabahealthvaccinnoticeordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinnoticeordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinnoticeordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinnoticeordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝用户id
func (r *AlibabahealthvaccinnoticeordercancelAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabahealthvaccinnoticeordercancelAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetOrderId is OrderId Setter
// 预约id
func (r *AlibabahealthvaccinnoticeordercancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthvaccinnoticeordercancelAPIRequest) GetOrderId() string {
	return r._orderId
}

package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinnoticeordersignAPIRequest 福州疫苗签到成功通知 API请求
// alibaba.health.vaccin.notice.order.sign
//
// 福州疫苗用户签到成功记录
type AlibabahealthvaccinnoticeordersignAPIRequest struct {
	model.Params
	// 支付宝用户id
	_alipayUserId string
	// 预约id
	_orderId string
}

// NewAlibabahealthvaccinnoticeordersignRequest 初始化AlibabahealthvaccinnoticeordersignAPIRequest对象
func NewAlibabahealthvaccinnoticeordersignRequest() *AlibabahealthvaccinnoticeordersignAPIRequest {
	return &AlibabahealthvaccinnoticeordersignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinnoticeordersignAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.order.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinnoticeordersignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinnoticeordersignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝用户id
func (r *AlibabahealthvaccinnoticeordersignAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabahealthvaccinnoticeordersignAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetOrderId is OrderId Setter
// 预约id
func (r *AlibabahealthvaccinnoticeordersignAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthvaccinnoticeordersignAPIRequest) GetOrderId() string {
	return r._orderId
}

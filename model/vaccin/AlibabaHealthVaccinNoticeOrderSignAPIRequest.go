package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeOrderSignAPIRequest
福州疫苗签到成功通知 API请求
alibaba.health.vaccin.notice.order.sign

福州疫苗用户签到成功记录 */
type AlibabaHealthVaccinNoticeOrderSignAPIRequest struct {
	model.Params
	// 支付宝用户id
	_alipayUserId string
	// 预约id
	_orderId string
}

// NewAlibabaHealthVaccinNoticeOrderSignRequest 初始化AlibabaHealthVaccinNoticeOrderSignAPIRequest对象
func NewAlibabaHealthVaccinNoticeOrderSignRequest() *AlibabaHealthVaccinNoticeOrderSignAPIRequest {
	return &AlibabaHealthVaccinNoticeOrderSignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.order.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderSignAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// Get AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// Set is OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderSignAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetOrderId() string {
	return r._orderId
}

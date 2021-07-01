package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeReplantRemindAPIRequest
支付宝疫苗补种提醒信息 API请求
alibaba.health.vaccin.notice.replant.remind

支付宝疫苗补种提醒 */
type AlibabaHealthVaccinNoticeReplantRemindAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayUserId string
	// 针次
	_theTimes string
	// 预约id
	_orderId string
}

// NewAlibabaHealthVaccinNoticeReplantRemindRequest 初始化AlibabaHealthVaccinNoticeReplantRemindAPIRequest对象
func NewAlibabaHealthVaccinNoticeReplantRemindRequest() *AlibabaHealthVaccinNoticeReplantRemindAPIRequest {
	return &AlibabaHealthVaccinNoticeReplantRemindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeReplantRemindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.replant.remind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeReplantRemindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayUserId Setter
// 支付宝ID
func (r *AlibabaHealthVaccinNoticeReplantRemindAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// Get AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeReplantRemindAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// Set is TheTimes Setter
// 针次
func (r *AlibabaHealthVaccinNoticeReplantRemindAPIRequest) SetTheTimes(_theTimes string) error {
	r._theTimes = _theTimes
	r.Set("the_times", _theTimes)
	return nil
}

// Get TheTimes Getter
func (r AlibabaHealthVaccinNoticeReplantRemindAPIRequest) GetTheTimes() string {
	return r._theTimes
}

// Set is OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeReplantRemindAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHealthVaccinNoticeReplantRemindAPIRequest) GetOrderId() string {
	return r._orderId
}

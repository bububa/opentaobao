package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeOrderSignAPIRequest 福州疫苗签到成功通知 API请求
// alibaba.health.vaccin.notice.order.sign
//
// 福州疫苗用户签到成功记录
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthVaccinNoticeOrderSignAPIRequest) Reset() {
	r._alipayUserId = ""
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.order.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderSignAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetOrderId is OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderSignAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderSignAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlibabaHealthVaccinNoticeOrderSignAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthVaccinNoticeOrderSignRequest()
	},
}

// GetAlibabaHealthVaccinNoticeOrderSignRequest 从 sync.Pool 获取 AlibabaHealthVaccinNoticeOrderSignAPIRequest
func GetAlibabaHealthVaccinNoticeOrderSignAPIRequest() *AlibabaHealthVaccinNoticeOrderSignAPIRequest {
	return poolAlibabaHealthVaccinNoticeOrderSignAPIRequest.Get().(*AlibabaHealthVaccinNoticeOrderSignAPIRequest)
}

// ReleaseAlibabaHealthVaccinNoticeOrderSignAPIRequest 将 AlibabaHealthVaccinNoticeOrderSignAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeOrderSignAPIRequest(v *AlibabaHealthVaccinNoticeOrderSignAPIRequest) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeOrderSignAPIRequest.Put(v)
}

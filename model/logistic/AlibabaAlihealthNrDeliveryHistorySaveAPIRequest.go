package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrDeliveryHistorySaveAPIRequest
物流信息回传接口 API请求
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息 */
type AlibabaAlihealthNrDeliveryHistorySaveAPIRequest struct {
	model.Params
	// 入参
	_deliveryEvent *DeliveryEventDto
}

// NewAlibabaAlihealthNrDeliveryHistorySaveRequest 初始化AlibabaAlihealthNrDeliveryHistorySaveAPIRequest对象
func NewAlibabaAlihealthNrDeliveryHistorySaveRequest() *AlibabaAlihealthNrDeliveryHistorySaveAPIRequest {
	return &AlibabaAlihealthNrDeliveryHistorySaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrDeliveryHistorySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.delivery.history.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrDeliveryHistorySaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeliveryEvent Setter
// 入参
func (r *AlibabaAlihealthNrDeliveryHistorySaveAPIRequest) SetDeliveryEvent(_deliveryEvent *DeliveryEventDto) error {
	r._deliveryEvent = _deliveryEvent
	r.Set("delivery_event", _deliveryEvent)
	return nil
}

// Get DeliveryEvent Getter
func (r AlibabaAlihealthNrDeliveryHistorySaveAPIRequest) GetDeliveryEvent() *DeliveryEventDto {
	return r._deliveryEvent
}

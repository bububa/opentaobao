package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressFulfillmentEventAPIRequest AE履约事件处理 API请求
// aliexpress.fulfillment.event
//
// AE用 履约底层声明发货能力
type AliexpressFulfillmentEventAPIRequest struct {
	model.Params
	// 入参对象
	_param *FulfillmentOrderStatusUpdateRequest
}

// NewAliexpressFulfillmentEventRequest 初始化AliexpressFulfillmentEventAPIRequest对象
func NewAliexpressFulfillmentEventRequest() *AliexpressFulfillmentEventAPIRequest {
	return &AliexpressFulfillmentEventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressFulfillmentEventAPIRequest) GetApiMethodName() string {
	return "aliexpress.fulfillment.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressFulfillmentEventAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参对象
func (r *AliexpressFulfillmentEventAPIRequest) SetParam(_param *FulfillmentOrderStatusUpdateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AliexpressFulfillmentEventAPIRequest) GetParam() *FulfillmentOrderStatusUpdateRequest {
	return r._param
}

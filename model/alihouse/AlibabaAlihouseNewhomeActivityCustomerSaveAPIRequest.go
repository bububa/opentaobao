package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitycustomersaveAPIRequest 销售活动批量保存定向用户 API请求
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
type AlibabaalihousenewhomeactivitycustomersaveAPIRequest struct {
	model.Params
	// 保存属性
	_activityCustomerSaveDto *ActivityCustomerSaveDto
}

// NewAlibabaalihousenewhomeactivitycustomersaveRequest 初始化AlibabaalihousenewhomeactivitycustomersaveAPIRequest对象
func NewAlibabaalihousenewhomeactivitycustomersaveRequest() *AlibabaalihousenewhomeactivitycustomersaveAPIRequest {
	return &AlibabaalihousenewhomeactivitycustomersaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeactivitycustomersaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.customer.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeactivitycustomersaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeactivitycustomersaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityCustomerSaveDto is ActivityCustomerSaveDto Setter
// 保存属性
func (r *AlibabaalihousenewhomeactivitycustomersaveAPIRequest) SetActivityCustomerSaveDto(_activityCustomerSaveDto *ActivityCustomerSaveDto) error {
	r._activityCustomerSaveDto = _activityCustomerSaveDto
	r.Set("activity_customer_save_dto", _activityCustomerSaveDto)
	return nil
}

// GetActivityCustomerSaveDto ActivityCustomerSaveDto Getter
func (r AlibabaalihousenewhomeactivitycustomersaveAPIRequest) GetActivityCustomerSaveDto() *ActivityCustomerSaveDto {
	return r._activityCustomerSaveDto
}

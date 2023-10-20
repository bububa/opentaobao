package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest 销售活动批量保存定向用户 API请求
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
type AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest struct {
	model.Params
	// 保存属性
	_activityCustomerSaveDto *ActivityCustomerSaveDto
}

// NewAlibabaAlihouseNewhomeActivityCustomerSaveRequest 初始化AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest对象
func NewAlibabaAlihouseNewhomeActivityCustomerSaveRequest() *AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest {
	return &AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.customer.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityCustomerSaveDto is ActivityCustomerSaveDto Setter
// 保存属性
func (r *AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest) SetActivityCustomerSaveDto(_activityCustomerSaveDto *ActivityCustomerSaveDto) error {
	r._activityCustomerSaveDto = _activityCustomerSaveDto
	r.Set("activity_customer_save_dto", _activityCustomerSaveDto)
	return nil
}

// GetActivityCustomerSaveDto ActivityCustomerSaveDto Getter
func (r AlibabaAlihouseNewhomeActivityCustomerSaveAPIRequest) GetActivityCustomerSaveDto() *ActivityCustomerSaveDto {
	return r._activityCustomerSaveDto
}

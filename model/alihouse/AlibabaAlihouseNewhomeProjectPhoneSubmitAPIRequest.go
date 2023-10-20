package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectphonesubmitAPIRequest 提交楼盘电话 API请求
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
type AlibabaalihousenewhomeprojectphonesubmitAPIRequest struct {
	model.Params
	// 外部门店ID
	_outerStoreId string
	// 楼盘电话
	_projectPhoneDto *ProjectPhoneDto
}

// NewAlibabaalihousenewhomeprojectphonesubmitRequest 初始化AlibabaalihousenewhomeprojectphonesubmitAPIRequest对象
func NewAlibabaalihousenewhomeprojectphonesubmitRequest() *AlibabaalihousenewhomeprojectphonesubmitAPIRequest {
	return &AlibabaalihousenewhomeprojectphonesubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectphonesubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.phone.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectphonesubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectphonesubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaalihousenewhomeprojectphonesubmitAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeprojectphonesubmitAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetProjectPhoneDto is ProjectPhoneDto Setter
// 楼盘电话
func (r *AlibabaalihousenewhomeprojectphonesubmitAPIRequest) SetProjectPhoneDto(_projectPhoneDto *ProjectPhoneDto) error {
	r._projectPhoneDto = _projectPhoneDto
	r.Set("project_phone_dto", _projectPhoneDto)
	return nil
}

// GetProjectPhoneDto ProjectPhoneDto Getter
func (r AlibabaalihousenewhomeprojectphonesubmitAPIRequest) GetProjectPhoneDto() *ProjectPhoneDto {
	return r._projectPhoneDto
}

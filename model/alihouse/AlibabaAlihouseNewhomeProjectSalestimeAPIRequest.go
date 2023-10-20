package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectsalestimeAPIRequest 楼盘销售时刻同步 API请求
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
type AlibabaalihousenewhomeprojectsalestimeAPIRequest struct {
	model.Params
	// 楼盘销售时刻对象
	_projectSalesTimeDto *ProjectSalesTimeDto
}

// NewAlibabaalihousenewhomeprojectsalestimeRequest 初始化AlibabaalihousenewhomeprojectsalestimeAPIRequest对象
func NewAlibabaalihousenewhomeprojectsalestimeRequest() *AlibabaalihousenewhomeprojectsalestimeAPIRequest {
	return &AlibabaalihousenewhomeprojectsalestimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectsalestimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.salestime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectsalestimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectsalestimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectSalesTimeDto is ProjectSalesTimeDto Setter
// 楼盘销售时刻对象
func (r *AlibabaalihousenewhomeprojectsalestimeAPIRequest) SetProjectSalesTimeDto(_projectSalesTimeDto *ProjectSalesTimeDto) error {
	r._projectSalesTimeDto = _projectSalesTimeDto
	r.Set("project_sales_time_dto", _projectSalesTimeDto)
	return nil
}

// GetProjectSalesTimeDto ProjectSalesTimeDto Getter
func (r AlibabaalihousenewhomeprojectsalestimeAPIRequest) GetProjectSalesTimeDto() *ProjectSalesTimeDto {
	return r._projectSalesTimeDto
}

package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectSalestimeAPIRequest 楼盘销售时刻同步 API请求
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
type AlibabaAlihouseNewhomeProjectSalestimeAPIRequest struct {
	model.Params
	// 楼盘销售时刻对象
	_projectSalesTimeDto *ProjectSalesTimeDto
}

// NewAlibabaAlihouseNewhomeProjectSalestimeRequest 初始化AlibabaAlihouseNewhomeProjectSalestimeAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectSalestimeRequest() *AlibabaAlihouseNewhomeProjectSalestimeAPIRequest {
	return &AlibabaAlihouseNewhomeProjectSalestimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectSalestimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.salestime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectSalestimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectSalestimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectSalesTimeDto is ProjectSalesTimeDto Setter
// 楼盘销售时刻对象
func (r *AlibabaAlihouseNewhomeProjectSalestimeAPIRequest) SetProjectSalesTimeDto(_projectSalesTimeDto *ProjectSalesTimeDto) error {
	r._projectSalesTimeDto = _projectSalesTimeDto
	r.Set("project_sales_time_dto", _projectSalesTimeDto)
	return nil
}

// GetProjectSalesTimeDto ProjectSalesTimeDto Getter
func (r AlibabaAlihouseNewhomeProjectSalestimeAPIRequest) GetProjectSalesTimeDto() *ProjectSalesTimeDto {
	return r._projectSalesTimeDto
}

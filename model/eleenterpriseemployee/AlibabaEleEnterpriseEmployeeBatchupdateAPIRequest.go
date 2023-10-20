package eleenterpriseemployee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseemployeebatchupdateAPIRequest 批量新增更新员工 API请求
// alibaba.ele.enterprise.employee.batchupdate
//
// 批量新增更新员工
type AlibabaeleenterpriseemployeebatchupdateAPIRequest struct {
	model.Params
	// 批量员工信息
	_enterpriseDatas []EmployeeInfoDto
}

// NewAlibabaeleenterpriseemployeebatchupdateRequest 初始化AlibabaeleenterpriseemployeebatchupdateAPIRequest对象
func NewAlibabaeleenterpriseemployeebatchupdateRequest() *AlibabaeleenterpriseemployeebatchupdateAPIRequest {
	return &AlibabaeleenterpriseemployeebatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriseemployeebatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.employee.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriseemployeebatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriseemployeebatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEnterpriseDatas is EnterpriseDatas Setter
// 批量员工信息
func (r *AlibabaeleenterpriseemployeebatchupdateAPIRequest) SetEnterpriseDatas(_enterpriseDatas []EmployeeInfoDto) error {
	r._enterpriseDatas = _enterpriseDatas
	r.Set("enterprise_datas", _enterpriseDatas)
	return nil
}

// GetEnterpriseDatas EnterpriseDatas Getter
func (r AlibabaeleenterpriseemployeebatchupdateAPIRequest) GetEnterpriseDatas() []EmployeeInfoDto {
	return r._enterpriseDatas
}

package eleenterpriseemployee

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest 批量新增更新员工 API请求
// alibaba.ele.enterprise.employee.batchupdate
//
// 批量新增更新员工
type AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest struct {
	model.Params
	// 批量员工信息
	_enterpriseDatas []EmployeeInfoDto
}

// NewAlibabaEleEnterpriseEmployeeBatchupdateRequest 初始化AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest对象
func NewAlibabaEleEnterpriseEmployeeBatchupdateRequest() *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest {
	return &AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) Reset() {
	r._enterpriseDatas = r._enterpriseDatas[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.employee.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEnterpriseDatas is EnterpriseDatas Setter
// 批量员工信息
func (r *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) SetEnterpriseDatas(_enterpriseDatas []EmployeeInfoDto) error {
	r._enterpriseDatas = _enterpriseDatas
	r.Set("enterprise_datas", _enterpriseDatas)
	return nil
}

// GetEnterpriseDatas EnterpriseDatas Getter
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetEnterpriseDatas() []EmployeeInfoDto {
	return r._enterpriseDatas
}

var poolAlibabaEleEnterpriseEmployeeBatchupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseEmployeeBatchupdateRequest()
	},
}

// GetAlibabaEleEnterpriseEmployeeBatchupdateRequest 从 sync.Pool 获取 AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest
func GetAlibabaEleEnterpriseEmployeeBatchupdateAPIRequest() *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest {
	return poolAlibabaEleEnterpriseEmployeeBatchupdateAPIRequest.Get().(*AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest)
}

// ReleaseAlibabaEleEnterpriseEmployeeBatchupdateAPIRequest 将 AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseEmployeeBatchupdateAPIRequest(v *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseEmployeeBatchupdateAPIRequest.Put(v)
}

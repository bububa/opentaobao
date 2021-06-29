package eleenterpriseemployee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量新增更新员工 API请求
alibaba.ele.enterprise.employee.batchupdate

批量新增更新员工
*/
type AlibabaEleEnterpriseEmployeeBatchupdateRequest struct {
    model.Params
    // 批量员工信息
    _enterpriseDatas   []EmployeeInfoDto
}

// 初始化AlibabaEleEnterpriseEmployeeBatchupdateRequest对象
func NewAlibabaEleEnterpriseEmployeeBatchupdateRequest() *AlibabaEleEnterpriseEmployeeBatchupdateRequest{
    return &AlibabaEleEnterpriseEmployeeBatchupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseEmployeeBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.employee.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseEmployeeBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EnterpriseDatas Setter
// 批量员工信息
func (r *AlibabaEleEnterpriseEmployeeBatchupdateRequest) SetEnterpriseDatas(_enterpriseDatas []EmployeeInfoDto) error {
    r._enterpriseDatas = _enterpriseDatas
    r.Set("enterprise_datas", _enterpriseDatas)
    return nil
}

// EnterpriseDatas Getter
func (r AlibabaEleEnterpriseEmployeeBatchupdateRequest) GetEnterpriseDatas() []EmployeeInfoDto {
    return r._enterpriseDatas
}

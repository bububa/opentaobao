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
type AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest struct {
    model.Params
    // 批量员工信息
    _enterpriseDatas   []EmployeeInfoDTO
}

// 初始化AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest对象
func NewAlibabaEleEnterpriseEmployeeBatchupdateRequest() *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest{
    return &AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.employee.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EnterpriseDatas Setter
// 批量员工信息
func (r *AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) SetEnterpriseDatas(_enterpriseDatas []EmployeeInfoDTO) error {
    r._enterpriseDatas = _enterpriseDatas
    r.Set("enterprise_datas", _enterpriseDatas)
    return nil
}

// EnterpriseDatas Getter
func (r AlibabaEleEnterpriseEmployeeBatchupdateAPIRequest) GetEnterpriseDatas() []EmployeeInfoDTO {
    return r._enterpriseDatas
}

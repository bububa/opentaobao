package eleenterpriseemployee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量新增更新员工 APIRequest
alibaba.ele.enterprise.employee.batchupdate

批量新增更新员工
*/
type AlibabaEleEnterpriseEmployeeBatchupdateRequest struct {
    model.Params

    // 批量员工信息
    enterpriseDatas   []EmployeeInfoDto 

}

func NewAlibabaEleEnterpriseEmployeeBatchupdateRequest() *AlibabaEleEnterpriseEmployeeBatchupdateRequest{
    return &AlibabaEleEnterpriseEmployeeBatchupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseEmployeeBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.employee.batchupdate"
}

func (r AlibabaEleEnterpriseEmployeeBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseEmployeeBatchupdateRequest) SetEnterpriseDatas(enterpriseDatas []EmployeeInfoDto) error {
    r.enterpriseDatas = enterpriseDatas
    r.Set("enterprise_datas", enterpriseDatas)
    return nil
}

func (r AlibabaEleEnterpriseEmployeeBatchupdateRequest) GetEnterpriseDatas() []EmployeeInfoDto {
    return r.enterpriseDatas
}


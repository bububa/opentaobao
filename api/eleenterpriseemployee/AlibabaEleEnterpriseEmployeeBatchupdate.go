package eleenterpriseemployee

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriseemployee"
)

/* 
批量新增更新员工 
alibaba.ele.enterprise.employee.batchupdate

批量新增更新员工
*/
func AlibabaEleEnterpriseEmployeeBatchupdate(clt *core.SDKClient, req *eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateRequest, session string) (*eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse, error) {
    var resp eleenterpriseemployee.AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

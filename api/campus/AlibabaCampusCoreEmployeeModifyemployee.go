package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
修改员工基本信息 
alibaba.campus.core.employee.modifyemployee

根据用户ID和公司ID更新员工基本信息（头像、性别、昵称）
*/
func AlibabaCampusCoreEmployeeModifyemployee(clt *core.SDKClient, req *campus.AlibabaCampusCoreEmployeeModifyemployeeRequest, session string) (*campus.AlibabaCampusCoreEmployeeModifyemployeeAPIResponse, error) {
    var resp campus.AlibabaCampusCoreEmployeeModifyemployeeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

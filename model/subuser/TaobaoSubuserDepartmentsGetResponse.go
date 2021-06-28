package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定账户的所有部门列表 APIResponse
taobao.subuser.departments.get

获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
*/
type TaobaoSubuserDepartmentsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSubuserDepartmentsGetResponse `json:"subuser_departments_get_response,omitempty"` 
    TaobaoSubuserDepartmentsGetResponse
}

/* model for simplify = false
type TaobaoSubuserDepartmentsGetResponse struct {

    // 部门信息
    
    Departments  struct {
        Department  []Department `json:"department,omitempty"`
    } `json:"departments,omitempty"`
    

}
*/

type TaobaoSubuserDepartmentsGetResponse struct {

    // 部门信息
    Departments   []Department `json:"departments,omitempty"`

}

package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubuserDepartmentsGetAPIResponse 获取指定账户的所有部门列表 API返回值
// taobao.subuser.departments.get
//
// 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
type TaobaoSubuserDepartmentsGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubuserDepartmentsGetAPIResponseModel
}

// TaobaoSubuserDepartmentsGetAPIResponseModel is 获取指定账户的所有部门列表 成功返回结果
type TaobaoSubuserDepartmentsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subuser_departments_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 部门信息
	Departments []Department `json:"departments,omitempty" xml:"departments>department,omitempty"`
}

package subuser

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoSubuserDepartmentsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubuserDepartmentsGetAPIResponseModel).Reset()
}

// TaobaoSubuserDepartmentsGetAPIResponseModel is 获取指定账户的所有部门列表 成功返回结果
type TaobaoSubuserDepartmentsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subuser_departments_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 部门信息
	Departments []Department `json:"departments,omitempty" xml:"departments>department,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubuserDepartmentsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Departments = m.Departments[:0]
}

var poolTaobaoSubuserDepartmentsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubuserDepartmentsGetAPIResponse)
	},
}

// GetTaobaoSubuserDepartmentsGetAPIResponse 从 sync.Pool 获取 TaobaoSubuserDepartmentsGetAPIResponse
func GetTaobaoSubuserDepartmentsGetAPIResponse() *TaobaoSubuserDepartmentsGetAPIResponse {
	return poolTaobaoSubuserDepartmentsGetAPIResponse.Get().(*TaobaoSubuserDepartmentsGetAPIResponse)
}

// ReleaseTaobaoSubuserDepartmentsGetAPIResponse 将 TaobaoSubuserDepartmentsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubuserDepartmentsGetAPIResponse(v *TaobaoSubuserDepartmentsGetAPIResponse) {
	v.Reset()
	poolTaobaoSubuserDepartmentsGetAPIResponse.Put(v)
}

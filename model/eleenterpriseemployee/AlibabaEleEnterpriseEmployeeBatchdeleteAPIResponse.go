package eleenterpriseemployee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse 批量删除员工 API返回值
// alibaba.ele.enterprise.employee.batchdelete
//
// 批量删除员工
type AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponseModel is 批量删除员工 成功返回结果
type AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_employee_batchdelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回值信息
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse)
	},
}

// GetAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse
func GetAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse() *AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse {
	return poolAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse.Get().(*AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse)
}

// ReleaseAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse 将 AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse(v *AlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseEmployeeBatchdeleteAPIResponse.Put(v)
}

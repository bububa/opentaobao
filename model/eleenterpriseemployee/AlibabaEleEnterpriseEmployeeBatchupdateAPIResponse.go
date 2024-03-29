package eleenterpriseemployee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse 批量新增更新员工 API返回值
// alibaba.ele.enterprise.employee.batchupdate
//
// 批量新增更新员工
type AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseEmployeeBatchupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseEmployeeBatchupdateAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseEmployeeBatchupdateAPIResponseModel is 批量新增更新员工 成功返回结果
type AlibabaEleEnterpriseEmployeeBatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_employee_batchupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回结果
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseEmployeeBatchupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse)
	},
}

// GetAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse
func GetAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse() *AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse {
	return poolAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse.Get().(*AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse)
}

// ReleaseAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse 将 AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse(v *AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseEmployeeBatchupdateAPIResponse.Put(v)
}

package eleenterpriseemployee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量新增更新员工 APIResponse
alibaba.ele.enterprise.employee.batchupdate

批量新增更新员工
*/
type AlibabaEleEnterpriseEmployeeBatchupdateAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseEmployeeBatchupdateResponse
}

type AlibabaEleEnterpriseEmployeeBatchupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_employee_batchupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应code
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 响应信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
    // 返回结果
    
    EnterpriseData   *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`

    
}

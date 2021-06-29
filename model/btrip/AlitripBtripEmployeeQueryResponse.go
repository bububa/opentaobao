package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
企业员工查询 APIResponse
alitrip.btrip.employee.query

企业员工查询
*/
type AlitripBtripEmployeeQueryAPIResponse struct {
    model.CommonResponse
    AlitripBtripEmployeeQueryResponse
}

type AlitripBtripEmployeeQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_employee_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象。
    
    Result   *BtmsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

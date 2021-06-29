package eleenterprisecartnew

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新版购物车查询 APIResponse
alibaba.ele.enterprise.cartnew.query

新版购物车查询
*/
type AlibabaEleEnterpriseCartnewQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseCartnewQueryResponse
}

type AlibabaEleEnterpriseCartnewQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_cartnew_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回码
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 返回信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 系统自动生成
    
    EnterpriseData   *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
}

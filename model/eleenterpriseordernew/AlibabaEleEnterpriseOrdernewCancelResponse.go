package eleenterpriseordernew

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单取消 APIResponse
alibaba.ele.enterprise.ordernew.cancel

订单取消
*/
type AlibabaEleEnterpriseOrdernewCancelAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseOrdernewCancelResponse
}

type AlibabaEleEnterpriseOrdernewCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应code
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 响应信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
}

package eleenterpriseordernew

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单详情 APIResponse
alibaba.ele.enterprise.ordernew.get

查询订单详情
*/
type AlibabaEleEnterpriseOrdernewGetAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseOrdernewGetResponse
}

type AlibabaEleEnterpriseOrdernewGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值信息
    
    EnterpriseData   *StandardOrderTrackingInfoDto `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`

    
    // 响应code
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 响应信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
}

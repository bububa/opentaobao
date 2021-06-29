package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购添加活动留资入口 APIResponse
alibaba.lsy.crm.customer.add

私域导购添加活动留资入口
*/
type AlibabaLsyCrmCustomerAddAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmCustomerAddResponse
}

type AlibabaLsyCrmCustomerAddResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_customer_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *AlibabaLsyCrmCustomerAddResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}

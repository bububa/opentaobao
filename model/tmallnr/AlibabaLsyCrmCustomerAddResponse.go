package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购添加活动留资入口 API返回值 
alibaba.lsy.crm.customer.add

私域导购添加活动留资入口
*/
type AlibabaLsyCrmCustomerAddAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmCustomerAddResponse
}

// 私域导购添加活动留资入口 成功返回结果
type AlibabaLsyCrmCustomerAddResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_customer_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *AlibabaLsyCrmCustomerAddResultDO `json:"result,omitempty" xml:"result,omitempty"`
}

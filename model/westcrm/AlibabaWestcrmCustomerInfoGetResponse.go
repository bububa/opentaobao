package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息查询接口 APIResponse
alibaba.westcrm.customer.info.get

会员信息查询接口
*/
type AlibabaWestcrmCustomerInfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmCustomerInfoGetResponse
}

type AlibabaWestcrmCustomerInfoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_customer_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

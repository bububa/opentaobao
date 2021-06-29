package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取RFQ详情 APIResponse
alibaba.icbu.rfqdetail.get

查看RFQ的详情信息
*/
type AlibabaIcbuRfqdetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuRfqdetailGetResponse
}

type AlibabaIcbuRfqdetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_rfqdetail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果集
    
    Result   *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

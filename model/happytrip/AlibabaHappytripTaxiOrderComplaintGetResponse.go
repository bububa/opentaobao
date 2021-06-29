package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
投诉详情 APIResponse
alibaba.happytrip.taxi.order.complaint.get

获取投诉处理进度详情
*/
type AlibabaHappytripTaxiOrderComplaintGetAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiOrderComplaintGetResponse
}

type AlibabaHappytripTaxiOrderComplaintGetResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_complaint_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
    // 投诉详情获取结果
    
    Data   *AlibabaHappytripTaxiOrderComplaintGetStruct `json:"data,omitempty" xml:"data,omitempty"`

    
}

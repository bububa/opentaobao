package mydata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取客户产品相关表现数据的可用时间范围 APIResponse
alibaba.mydata.self.product.date.get

获取客户产品相关表现数据的可用时间范围
*/
type AlibabaMydataSelfProductDateGetAPIResponse struct {
    model.CommonResponse
    AlibabaMydataSelfProductDateGetResponse
}

type AlibabaMydataSelfProductDateGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mydata_self_product_date_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // endDate
    
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`

    
    // startDate
    
    StartDate   string `json:"start_date,omitempty" xml:"start_date,omitempty"`

    
}

package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获得当前系统时间 APIResponse
alibaba.wdk.time.get

获得当前系统时间
*/
type AlibabaWdkTimeGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTimeGetResponse
}

type AlibabaWdkTimeGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_time_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // dateTime
    
    DateTime   int64 `json:"date_time,omitempty" xml:"date_time,omitempty"`

    
    // date
    
    Date   string `json:"date,omitempty" xml:"date,omitempty"`

    
}

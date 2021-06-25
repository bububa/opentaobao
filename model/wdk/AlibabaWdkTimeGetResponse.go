package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获得当前系统时间 APIResponse
alibaba.wdk.time.get

获得当前系统时间
*/
type AlibabaWdkTimeGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkTimeGetResponse `json:"alibaba_wdk_time_get_response,omitempty"`
}

type AlibabaWdkTimeGetResponse struct {

    // dateTime
    DateTime   int64 `json:"date_time,omitempty"`

    // date
    Date   string `json:"date,omitempty"`

}

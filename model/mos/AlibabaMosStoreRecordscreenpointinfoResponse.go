package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云屏埋点数据记录接口 APIResponse
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据
*/
type AlibabaMosStoreRecordscreenpointinfoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_store_recordscreenpointinfo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaMosStoreRecordscreenpointinfoResultDo `json:"result,omitempty" xml:"
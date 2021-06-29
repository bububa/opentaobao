package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云屏埋点数据记录接口 API返回值 
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据
*/
type AlibabaMosStoreRecordscreenpointinfoAPIResponse struct {
    model.CommonResponse
    AlibabaMosStoreRecordscreenpointinfoResponse
}

// 云屏埋点数据记录接口 成功返回结果
type AlibabaMosStoreRecordscreenpointinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_store_recordscreenpointinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaMosStoreRecordscreenpointinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
云屏埋点数据记录接口 APIResponse
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据
*/
type AlibabaMosStoreRecordscreenpointinfoAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosStoreRecordscreenpointinfoResponse `json:"alibaba_mos_store_recordscreenpointinfo_response,omitempty"`
}

type AlibabaMosStoreRecordscreenpointinfoResponse struct {

    // result
    Result   *AlibabaMosStoreRecordscreenpointinfoResultDo `json:"result,omitempty"`

}

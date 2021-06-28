package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
支付参考号回传 APIResponse
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc
*/
type AlibabaMjOcSyncpayinfoAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjOcSyncpayinfoResponse `json:"alibaba_mj_oc_syncpayinfo_response,omitempty"` 
    AlibabaMjOcSyncpayinfoResponse
}

/* model for simplify = false
type AlibabaMjOcSyncpayinfoResponse struct {

    // result
    
    Result  *struct {
        AlibabaMjOcSyncpayinfoResultDo  *AlibabaMjOcSyncpayinfoResultDo `json:"alibaba_mj_oc_syncpayinfo_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMjOcSyncpayinfoResponse struct {

    // result
    Result   *AlibabaMjOcSyncpayinfoResultDo `json:"result,omitempty"`

}

package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
支付参考号回传 APIResponse
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc
*/
type AlibabaMjOcSyncpayinfoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_syncpayinfo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaMjOcSyncpayinfoResultDo `json:"result,omitempty" xml:"
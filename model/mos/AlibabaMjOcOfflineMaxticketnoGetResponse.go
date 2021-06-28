package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
pos机获取线下最大小票号 APIResponse
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询
*/
type AlibabaMjOcOfflineMaxticketnoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_offline_maxticketno_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 时间
    
    PayDate   string `json:"pay_date,omitempty" xml:"
package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资源位数据推送接口 APIResponse
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据
*/
type AlibabaInteractAopdataRegisterAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_aopdata_register_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaInteractAopdataRegisterResult `json:"result,omitempty" xml:"
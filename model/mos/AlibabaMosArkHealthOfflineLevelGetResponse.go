package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mall的离线等级 APIResponse
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级
*/
type AlibabaMosArkHealthOfflineLevelGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_ark_health_offline_level_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // model
    
    Data   string `json:"data,omitempty" xml:"
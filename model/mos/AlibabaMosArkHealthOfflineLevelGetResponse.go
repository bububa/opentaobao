package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取mall的离线等级 APIResponse
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级
*/
type AlibabaMosArkHealthOfflineLevelGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosArkHealthOfflineLevelGetResponse `json:"alibaba_mos_ark_health_offline_level_get_response,omitempty"`
}

type AlibabaMosArkHealthOfflineLevelGetResponse struct {

    // model
    Data   string `json:"data,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

}

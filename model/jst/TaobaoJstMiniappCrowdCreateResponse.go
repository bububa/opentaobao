package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动创建 APIResponse
taobao.jst.miniapp.crowd.create

小程序活动创建
*/
type TaobaoJstMiniappCrowdCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_miniapp_crowd_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动Code，活动的唯一标识
    
    Result   string `json:"result,omitempty" xml:"
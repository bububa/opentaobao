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
    TaobaoJstMiniappCrowdCreateResponse
}

type TaobaoJstMiniappCrowdCreateResponse struct {
    XMLName xml.Name `xml:"jst_miniapp_crowd_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动Code，活动的唯一标识
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
    // 状态码
    
    RCode   int64 `json:"r_code,omitempty" xml:"r_code,omitempty"`

    
}

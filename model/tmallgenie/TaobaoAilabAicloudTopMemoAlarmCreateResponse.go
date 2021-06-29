package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵闹钟创建 APIResponse
taobao.ailab.aicloud.top.memo.alarm.create

天猫精灵闹钟创建
*/
type TaobaoAilabAicloudTopMemoAlarmCreateAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMemoAlarmCreateResponse
}

type TaobaoAilabAicloudTopMemoAlarmCreateResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态描述
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 闹钟ID
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
    // 状态码
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
}

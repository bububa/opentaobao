package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵会议查询 APIResponse
taobao.ailab.aicloud.top.memo.meeting.list

查询天猫精灵用户设置的所有会议
*/
type TaobaoAilabAicloudTopMemoMeetingListAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMemoMeetingListResponse
}

type TaobaoAilabAicloudTopMemoMeetingListResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_memo_meeting_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务的结果封装
    
    Result   *TaobaoAilabAicloudTopMemoMeetingListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

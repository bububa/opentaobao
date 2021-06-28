package wangwang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聊天关系获取接口 APIResponse
taobao.wangwang.eservice.chatrelation.get

获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
*/
type TaobaoWangwangEserviceChatrelationGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wangwang_eservice_chatrelation_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ChatRelationResult `json:"result,omitempty" xml:"
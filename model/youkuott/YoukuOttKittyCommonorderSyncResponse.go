package youkuott

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商一般订单同步 APIResponse
youku.ott.kitty.commonorder.sync

运营商一般订单同步
*/
type YoukuOttKittyCommonorderSyncAPIResponse struct {
    model.CommonResponse
    YoukuOttKittyCommonorderSyncResponse
}

type YoukuOttKittyCommonorderSyncResponse struct {
    XMLName xml.Name `xml:"youku_ott_kitty_commonorder_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功 true:成功 false:失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回码（一般情况请求方只需要关心success，除非特殊情况需要关心错误码）
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
}

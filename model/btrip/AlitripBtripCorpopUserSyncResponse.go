package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部人员同步 APIResponse
alitrip.btrip.corpop.user.sync

同步外部平台用户信息至商旅内部
*/
type AlitripBtripCorpopUserSyncAPIResponse struct {
    model.CommonResponse
    AlitripBtripCorpopUserSyncResponse
}

type AlitripBtripCorpopUserSyncResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_corpop_user_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 错误码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 返回错误用户信息
    
    Module   string `json:"module,omitempty" xml:"module,omitempty"`

    
    // 成功标示
    
    SuccessFlag   bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`

    
}

package cloudgame

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏混淆用户ID校验 APIResponse
alibaba.cloudgame.user.mixuserid.check

验证混淆用户ID是否合法
*/
type AlibabaCloudgameUserMixuseridCheckAPIResponse struct {
    model.CommonResponse
    AlibabaCloudgameUserMixuseridCheckResponse
}

type AlibabaCloudgameUserMixuseridCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_cloudgame_user_mixuserid_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 是否有效
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
    // 返回码描述
    
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`

    
}

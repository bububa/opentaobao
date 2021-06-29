package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户技能 Oauth 授权（淘宝 openId） API返回值 
alibaba.ailab.tb.user.skill.oauth

定制机厂商，在用户配网完成后，厂商调用此接口，写入特定技能的 Oauth 信息
*/
type AlibabaAilabTbUserSkillOauthAPIResponse struct {
    model.CommonResponse
    AlibabaAilabTbUserSkillOauthResponse
}

// 用户技能 Oauth 授权（淘宝 openId） 成功返回结果
type AlibabaAilabTbUserSkillOauthResponse struct {
    XMLName xml.Name `xml:"alibaba_ailab_tb_user_skill_oauth_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 返回码，200 代表成功
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}

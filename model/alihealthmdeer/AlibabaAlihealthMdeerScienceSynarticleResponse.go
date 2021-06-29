package alihealthmdeer

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
医知鹿文章同步【保存/更新】 API返回值 
alibaba.alihealth.mdeer.science.synarticle

文章同步【保存/更新】
*/
type AlibabaAlihealthMdeerScienceSynarticleAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMdeerScienceSynarticleResponse
}

// 医知鹿文章同步【保存/更新】 成功返回结果
type AlibabaAlihealthMdeerScienceSynarticleResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_mdeer_science_synarticle_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 信息code
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 返回值
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}

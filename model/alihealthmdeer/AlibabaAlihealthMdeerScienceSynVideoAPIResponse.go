package alihealthmdeer

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频同步【保存/更新】 API返回值 
alibaba.alihealth.mdeer.science.synVideo

视频同步【保存/更新】
*/
type AlibabaAlihealthMdeerScienceSynVideoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel
}

// 视频同步【保存/更新】 成功返回结果
type AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_mdeer_science_synVideo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果描述
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 调用结果code
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 调用是否成功
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}

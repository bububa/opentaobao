package alihealthpw

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
申请节点变更回调 API返回值 
alibaba.alihealth.pw.applynode.update

基金会回调阿里健康更新药品援助申请单的状态
*/
type AlibabaAlihealthPwApplynodeUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPwApplynodeUpdateResponse
}

// 申请节点变更回调 成功返回结果
type AlibabaAlihealthPwApplynodeUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pw_applynode_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // pap项目状态码
    PapCode   string `json:"pap_code,omitempty" xml:"pap_code,omitempty"`
    // pap项目状态描述
    PapMessage   string `json:"pap_message,omitempty" xml:"pap_message,omitempty"`
}

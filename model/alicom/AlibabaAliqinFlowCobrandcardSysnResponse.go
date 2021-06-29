package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联名卡信息同步 API返回值 
alibaba.aliqin.flow.cobrandcard.sysn

提供给浙江移动同步联名卡信息接口。
*/
type AlibabaAliqinFlowCobrandcardSysnAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowCobrandcardSysnResponse
}

// 联名卡信息同步 成功返回结果
type AlibabaAliqinFlowCobrandcardSysnResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_cobrandcard_sysn_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // errorCode，错误码
    BizErrorCode   string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
    // errorMsg，错误描述
    BizErrorMsg   string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
    // 是否请求成功true 或者false
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    // error,如果有错，这个为true
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
}

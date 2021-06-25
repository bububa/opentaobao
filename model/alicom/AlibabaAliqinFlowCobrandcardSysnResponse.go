package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
联名卡信息同步 APIResponse
alibaba.aliqin.flow.cobrandcard.sysn

提供给浙江移动同步联名卡信息接口。
*/
type AlibabaAliqinFlowCobrandcardSysnAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowCobrandcardSysnResponse `json:"alibaba_aliqin_flow_cobrandcard_sysn_response,omitempty"`
}

type AlibabaAliqinFlowCobrandcardSysnResponse struct {

    // errorCode，错误码
    BizErrorCode   string `json:"biz_error_code,omitempty"`

    // errorMsg，错误描述
    BizErrorMsg   string `json:"biz_error_msg,omitempty"`

    // 是否请求成功true 或者false
    Value   string `json:"value,omitempty"`

    // error,如果有错，这个为true
    Error   bool `json:"error,omitempty"`

}

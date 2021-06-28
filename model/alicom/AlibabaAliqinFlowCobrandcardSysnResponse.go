package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联名卡信息同步 APIResponse
alibaba.aliqin.flow.cobrandcard.sysn

提供给浙江移动同步联名卡信息接口。
*/
type AlibabaAliqinFlowCobrandcardSysnAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_cobrandcard_sysn_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode，错误码
    
    BizErrorCode   string `json:"biz_error_code,omitempty" xml:"
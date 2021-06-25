package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
应用加固接口 APIResponse
alibaba.security.jaq.app.shield

提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
*/
type AlibabaSecurityJaqAppShieldAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSecurityJaqAppShieldResponse `json:"alibaba_security_jaq_app_shield_response,omitempty"`
}

type AlibabaSecurityJaqAppShieldResponse struct {

    // 加固任务信息
    Result   *ScanTaskInfo `json:"result,omitempty"`

}

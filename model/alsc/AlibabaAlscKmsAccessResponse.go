package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
本地生活风控数据接入 APIResponse
alibaba.alsc.kms.access

第三方使用本地生活数据对外提供服务，上报访问日志信息接口
*/
type AlibabaAlscKmsAccessAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscKmsAccessResponse `json:"alibaba_alsc_kms_access_response,omitempty"`
}

type AlibabaAlscKmsAccessResponse struct {

    // code
    Resultcode   string `json:"resultcode,omitempty"`

    // 是否成功
    Resultsuccess   bool `json:"resultsuccess,omitempty"`

    // message
    Resultmessage   string `json:"resultmessage,omitempty"`

}

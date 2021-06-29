package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端平台单据更新 API返回值 
alibaba.alihealth.drug.kyt.updatebillinfo

零售端平台单据更新
*/
type AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytUpdatebillinfoResponse
}

// 零售端平台单据更新 成功返回结果
type AlibabaAlihealthDrugKytUpdatebillinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_updatebillinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
    // 返回编码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 返回描述
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回是否成功
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

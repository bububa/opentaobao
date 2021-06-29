package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售修改出入库单追溯码 APIResponse
alibaba.alihealth.drug.kyt.updatebillcode

零售修改出入库单追溯码
*/
type AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytUpdatebillcodeResponse
}

type AlibabaAlihealthDrugKytUpdatebillcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_updatebillcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 返回编码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 返回描述
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 返回接口
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
}

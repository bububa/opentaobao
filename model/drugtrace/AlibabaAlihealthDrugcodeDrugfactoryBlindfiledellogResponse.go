package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
接收盲底文件删除日志 APIResponse
alibaba.alihealth.drugcode.drugfactory.blindfiledellog

临床用药试验-接收盲底文件删除日志
*/
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogResponse
}

type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_blindfiledellog_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 删除日志同步结果
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 调用结果
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 调用结果信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}

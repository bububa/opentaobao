package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
传输盲底文件 API返回值 
alibaba.alihealth.drugcode.drugfactory.transferblind

临床用药试验-传输盲底文件
*/
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel
}

// 传输盲底文件 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_transferblind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // 操作码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 操作说明
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}

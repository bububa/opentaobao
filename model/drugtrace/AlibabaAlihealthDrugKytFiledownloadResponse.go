package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处理失败单据下载 APIResponse
alibaba.alihealth.drug.kyt.filedownload

处理失败单据下载
*/
type AlibabaAlihealthDrugKytFiledownloadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytFiledownloadResponse
}

type AlibabaAlihealthDrugKytFiledownloadResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_filedownload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 文件内容
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 返回结果编码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 返回结果
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 返回是否成功
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
}

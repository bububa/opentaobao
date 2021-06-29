package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店销售扫码回传接口 APIResponse
alibaba.alihealth.drug.kyt.uploadretail

门店在销售给顾客时，扫描追溯码的数据按照单据回传；
*/
type AlibabaAlihealthDrugKytUploadretailAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytUploadretailResponse
}

type AlibabaAlihealthDrugKytUploadretailResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadretail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 上传单据文件队列表标识
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 错误码(BILL_DECODE_ERROR 单据转码失败 2.BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 是否成功(true 成功 ,false失败)
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
}

package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零头出入库单据上传 API返回值 
alibaba.alihealth.drug.kyt.remnantbill.upload

零头出入库单据上传
*/
type AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytRemnantbillUploadResponse
}

// 零头出入库单据上传 成功返回结果
type AlibabaAlihealthDrugKytRemnantbillUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_remnantbill_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    ResponseStatus   bool `json:"response_status,omitempty" xml:"response_status,omitempty"`
    // model
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

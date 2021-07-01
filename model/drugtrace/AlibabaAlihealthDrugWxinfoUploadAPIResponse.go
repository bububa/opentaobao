package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序数据回传 API返回值 
alibaba.alihealth.drug.wxinfo.upload

小程序数据回传
*/
type AlibabaAlihealthDrugWxinfoUploadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugWxinfoUploadAPIResponseModel
}

// 小程序数据回传 成功返回结果
type AlibabaAlihealthDrugWxinfoUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_wxinfo_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // model
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

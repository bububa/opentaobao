package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互联网医院批量导入接口 API返回值 
alibaba.alihealth.medicalbase.hospital.sync

互联网医院isv批量通过接口批量导入
*/
type AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel
}

// 互联网医院批量导入接口 成功返回结果
type AlibabaAlihealthMedicalbaseHospitalSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_hospital_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

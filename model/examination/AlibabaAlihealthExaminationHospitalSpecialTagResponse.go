package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构获取特色服务标签 APIResponse
alibaba.alihealth.examination.hospital.special.tag

体检机构获取特色服务标签列表
*/
type AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationHospitalSpecialTagResponse
}

type AlibabaAlihealthExaminationHospitalSpecialTagResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_hospital_special_tag_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

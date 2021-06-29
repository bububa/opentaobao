package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-普通签证-申请人进度推进接口 API返回值 
alitrip.travel.visa.applicant.update

普通签证订单的申请人进度推进接口，用于商家代用户填写申请人基本信息 或 推进单个申请人的签证进度。
*/
type AlitripTravelVisaApplicantUpdateAPIResponse struct {
    model.CommonResponse
    AlitripTravelVisaApplicantUpdateResponse
}

// 飞猪度假-普通签证-申请人进度推进接口 成功返回结果
type AlitripTravelVisaApplicantUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_visa_applicant_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 申请人状态更新是否成功
    UpdateResult   bool `json:"update_result,omitempty" xml:"update_result,omitempty"`
    // 该订单上已经上传的申请人列表信息
    Applicants   []NormalVisaApplicantInfo `json:"applicants,omitempty" xml:"applicants>normal_visa_applicant_info,omitempty"`
}

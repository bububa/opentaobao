package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
签证申请人查询接口 API返回值 
alitrip.travel.visa.applicant.query

签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
*/
type AlitripTravelVisaApplicantQueryAPIResponse struct {
    model.CommonResponse
    AlitripTravelVisaApplicantQueryAPIResponseModel
}

// 签证申请人查询接口 成功返回结果
type AlitripTravelVisaApplicantQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_travel_visa_applicant_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *QueryVisaApplicantResult `json:"result,omitempty" xml:"result,omitempty"`
}

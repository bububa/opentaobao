package normalvisa

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelVisaApplicantImportAPIResponse 签证申请人导入 API返回值
// alitrip.travel.visa.applicant.import
//
// 签证线下申请人导入接口。供商家将线下的签证申请人信息导入，进行签证线上化办理
type AlitripTravelVisaApplicantImportAPIResponse struct {
	model.CommonResponse
	AlitripTravelVisaApplicantImportAPIResponseModel
}

// AlitripTravelVisaApplicantImportAPIResponseModel is 签证申请人导入 成功返回结果
type AlitripTravelVisaApplicantImportAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_visa_applicant_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 外部商家申请人id
	OuterApplyId string `json:"outer_apply_id,omitempty" xml:"outer_apply_id,omitempty"`
	// 申请人id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 拼音姓
	LastNamePinyin string `json:"last_name_pinyin,omitempty" xml:"last_name_pinyin,omitempty"`
	// 拼音名
	FirstNamePinyin string `json:"first_name_pinyin,omitempty" xml:"first_name_pinyin,omitempty"`
	// 护照号
	PassportNumber string `json:"passport_number,omitempty" xml:"passport_number,omitempty"`
	// 错误类型:1-数据重复，2-数据错误
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

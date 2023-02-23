package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelVisaApplicantImportAPIRequest 签证申请人导入 API请求
// alitrip.travel.visa.applicant.import
//
// 签证线下申请人导入接口。供商家将线下的签证申请人信息导入，进行签证线上化办理
type AlitripTravelVisaApplicantImportAPIRequest struct {
	model.Params
	// 外部商家申请人id
	_outerApplyId string
	// 护照文件类型
	_passportFileType string
	// 证件照文件类型
	_photoFileType string
	// 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
	_formDataJson string
	// 证件照文件字节流
	_photoFile *model.File
	// 国家id。目前支持越南(27027)
	_nationId int64
	// 护照文件字节流
	_passportFile *model.File
}

// NewAlitripTravelVisaApplicantImportRequest 初始化AlitripTravelVisaApplicantImportAPIRequest对象
func NewAlitripTravelVisaApplicantImportRequest() *AlitripTravelVisaApplicantImportAPIRequest {
	return &AlitripTravelVisaApplicantImportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaApplicantImportAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.applicant.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaApplicantImportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelVisaApplicantImportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterApplyId is OuterApplyId Setter
// 外部商家申请人id
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetOuterApplyId(_outerApplyId string) error {
	r._outerApplyId = _outerApplyId
	r.Set("outer_apply_id", _outerApplyId)
	return nil
}

// GetOuterApplyId OuterApplyId Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetOuterApplyId() string {
	return r._outerApplyId
}

// SetPassportFileType is PassportFileType Setter
// 护照文件类型
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetPassportFileType(_passportFileType string) error {
	r._passportFileType = _passportFileType
	r.Set("passport_file_type", _passportFileType)
	return nil
}

// GetPassportFileType PassportFileType Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetPassportFileType() string {
	return r._passportFileType
}

// SetPhotoFileType is PhotoFileType Setter
// 证件照文件类型
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetPhotoFileType(_photoFileType string) error {
	r._photoFileType = _photoFileType
	r.Set("photo_file_type", _photoFileType)
	return nil
}

// GetPhotoFileType PhotoFileType Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetPhotoFileType() string {
	return r._photoFileType
}

// SetFormDataJson is FormDataJson Setter
// 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetFormDataJson(_formDataJson string) error {
	r._formDataJson = _formDataJson
	r.Set("form_data_json", _formDataJson)
	return nil
}

// GetFormDataJson FormDataJson Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetFormDataJson() string {
	return r._formDataJson
}

// SetPhotoFile is PhotoFile Setter
// 证件照文件字节流
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetPhotoFile(_photoFile *model.File) error {
	r._photoFile = _photoFile
	r.Set("photo_file", _photoFile)
	return nil
}

// GetPhotoFile PhotoFile Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetPhotoFile() *model.File {
	return r._photoFile
}

// SetNationId is NationId Setter
// 国家id。目前支持越南(27027)
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetNationId(_nationId int64) error {
	r._nationId = _nationId
	r.Set("nation_id", _nationId)
	return nil
}

// GetNationId NationId Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetNationId() int64 {
	return r._nationId
}

// SetPassportFile is PassportFile Setter
// 护照文件字节流
func (r *AlitripTravelVisaApplicantImportAPIRequest) SetPassportFile(_passportFile *model.File) error {
	r._passportFile = _passportFile
	r.Set("passport_file", _passportFile)
	return nil
}

// GetPassportFile PassportFile Getter
func (r AlitripTravelVisaApplicantImportAPIRequest) GetPassportFile() *model.File {
	return r._passportFile
}

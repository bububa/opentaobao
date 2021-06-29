package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签证申请人导入 API请求
alitrip.travel.visa.applicant.import

签证线下申请人导入接口。供商家将线下的签证申请人信息导入，进行签证线上化办理
*/
type AlitripTravelVisaApplicantImportRequest struct {
    model.Params
    // 国家id。目前支持越南(27027)
    _nationId   int64
    // 证件照文件字节流
    _photoFile   *model.File
    // 外部商家申请人id
    _outerApplyId   string
    // 护照文件类型
    _passportFileType   string
    // 护照文件字节流
    _passportFile   *model.File
    // 证件照文件类型
    _photoFileType   string
    // 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
    _formDataJson   string
}

// 初始化AlitripTravelVisaApplicantImportRequest对象
func NewAlitripTravelVisaApplicantImportRequest() *AlitripTravelVisaApplicantImportRequest{
    return &AlitripTravelVisaApplicantImportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaApplicantImportRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.applicant.import"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaApplicantImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NationId Setter
// 国家id。目前支持越南(27027)
func (r *AlitripTravelVisaApplicantImportRequest) SetNationId(_nationId int64) error {
    r._nationId = _nationId
    r.Set("nation_id", _nationId)
    return nil
}

// NationId Getter
func (r AlitripTravelVisaApplicantImportRequest) GetNationId() int64 {
    return r._nationId
}
// PhotoFile Setter
// 证件照文件字节流
func (r *AlitripTravelVisaApplicantImportRequest) SetPhotoFile(_photoFile *model.File) error {
    r._photoFile = _photoFile
    r.Set("photo_file", _photoFile)
    return nil
}

// PhotoFile Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPhotoFile() *model.File {
    return r._photoFile
}
// OuterApplyId Setter
// 外部商家申请人id
func (r *AlitripTravelVisaApplicantImportRequest) SetOuterApplyId(_outerApplyId string) error {
    r._outerApplyId = _outerApplyId
    r.Set("outer_apply_id", _outerApplyId)
    return nil
}

// OuterApplyId Getter
func (r AlitripTravelVisaApplicantImportRequest) GetOuterApplyId() string {
    return r._outerApplyId
}
// PassportFileType Setter
// 护照文件类型
func (r *AlitripTravelVisaApplicantImportRequest) SetPassportFileType(_passportFileType string) error {
    r._passportFileType = _passportFileType
    r.Set("passport_file_type", _passportFileType)
    return nil
}

// PassportFileType Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPassportFileType() string {
    return r._passportFileType
}
// PassportFile Setter
// 护照文件字节流
func (r *AlitripTravelVisaApplicantImportRequest) SetPassportFile(_passportFile *model.File) error {
    r._passportFile = _passportFile
    r.Set("passport_file", _passportFile)
    return nil
}

// PassportFile Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPassportFile() *model.File {
    return r._passportFile
}
// PhotoFileType Setter
// 证件照文件类型
func (r *AlitripTravelVisaApplicantImportRequest) SetPhotoFileType(_photoFileType string) error {
    r._photoFileType = _photoFileType
    r.Set("photo_file_type", _photoFileType)
    return nil
}

// PhotoFileType Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPhotoFileType() string {
    return r._photoFileType
}
// FormDataJson Setter
// 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
func (r *AlitripTravelVisaApplicantImportRequest) SetFormDataJson(_formDataJson string) error {
    r._formDataJson = _formDataJson
    r.Set("form_data_json", _formDataJson)
    return nil
}

// FormDataJson Getter
func (r AlitripTravelVisaApplicantImportRequest) GetFormDataJson() string {
    return r._formDataJson
}

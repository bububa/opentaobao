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
    nationId   int64
    // 证件照文件字节流
    photoFile   []*model.File
    // 外部商家申请人id
    outerApplyId   string
    // 护照文件类型
    passportFileType   string
    // 护照文件字节流
    passportFile   []*model.File
    // 证件照文件类型
    photoFileType   string
    // 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
    formDataJson   string
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
func (r *AlitripTravelVisaApplicantImportRequest) SetNationId(nationId int64) error {
    r.nationId = nationId
    r.Set("nation_id", nationId)
    return nil
}

// NationId Getter
func (r AlitripTravelVisaApplicantImportRequest) GetNationId() int64 {
    return r.nationId
}
// PhotoFile Setter
// 证件照文件字节流
func (r *AlitripTravelVisaApplicantImportRequest) SetPhotoFile(photoFile []*model.File) error {
    r.photoFile = photoFile
    r.Set("photo_file", photoFile)
    return nil
}

// PhotoFile Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPhotoFile() []*model.File {
    return r.photoFile
}
// OuterApplyId Setter
// 外部商家申请人id
func (r *AlitripTravelVisaApplicantImportRequest) SetOuterApplyId(outerApplyId string) error {
    r.outerApplyId = outerApplyId
    r.Set("outer_apply_id", outerApplyId)
    return nil
}

// OuterApplyId Getter
func (r AlitripTravelVisaApplicantImportRequest) GetOuterApplyId() string {
    return r.outerApplyId
}
// PassportFileType Setter
// 护照文件类型
func (r *AlitripTravelVisaApplicantImportRequest) SetPassportFileType(passportFileType string) error {
    r.passportFileType = passportFileType
    r.Set("passport_file_type", passportFileType)
    return nil
}

// PassportFileType Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPassportFileType() string {
    return r.passportFileType
}
// PassportFile Setter
// 护照文件字节流
func (r *AlitripTravelVisaApplicantImportRequest) SetPassportFile(passportFile []*model.File) error {
    r.passportFile = passportFile
    r.Set("passport_file", passportFile)
    return nil
}

// PassportFile Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPassportFile() []*model.File {
    return r.passportFile
}
// PhotoFileType Setter
// 证件照文件类型
func (r *AlitripTravelVisaApplicantImportRequest) SetPhotoFileType(photoFileType string) error {
    r.photoFileType = photoFileType
    r.Set("photo_file_type", photoFileType)
    return nil
}

// PhotoFileType Getter
func (r AlitripTravelVisaApplicantImportRequest) GetPhotoFileType() string {
    return r.photoFileType
}
// FormDataJson Setter
// 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
func (r *AlitripTravelVisaApplicantImportRequest) SetFormDataJson(formDataJson string) error {
    r.formDataJson = formDataJson
    r.Set("form_data_json", formDataJson)
    return nil
}

// FormDataJson Getter
func (r AlitripTravelVisaApplicantImportRequest) GetFormDataJson() string {
    return r.formDataJson
}

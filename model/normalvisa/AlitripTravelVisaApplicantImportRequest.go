package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签证申请人导入 APIRequest
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

func NewAlitripTravelVisaApplicantImportRequest() *AlitripTravelVisaApplicantImportRequest{
    return &AlitripTravelVisaApplicantImportRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelVisaApplicantImportRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.applicant.import"
}

func (r AlitripTravelVisaApplicantImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelVisaApplicantImportRequest) SetNationId(nationId int64) error {
    r.nationId = nationId
    r.Set("nation_id", nationId)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetNationId() int64 {
    return r.nationId
}

func (r *AlitripTravelVisaApplicantImportRequest) SetPhotoFile(photoFile []*model.File) error {
    r.photoFile = photoFile
    r.Set("photo_file", photoFile)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetPhotoFile() []*model.File {
    return r.photoFile
}

func (r *AlitripTravelVisaApplicantImportRequest) SetOuterApplyId(outerApplyId string) error {
    r.outerApplyId = outerApplyId
    r.Set("outer_apply_id", outerApplyId)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetOuterApplyId() string {
    return r.outerApplyId
}

func (r *AlitripTravelVisaApplicantImportRequest) SetPassportFileType(passportFileType string) error {
    r.passportFileType = passportFileType
    r.Set("passport_file_type", passportFileType)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetPassportFileType() string {
    return r.passportFileType
}

func (r *AlitripTravelVisaApplicantImportRequest) SetPassportFile(passportFile []*model.File) error {
    r.passportFile = passportFile
    r.Set("passport_file", passportFile)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetPassportFile() []*model.File {
    return r.passportFile
}

func (r *AlitripTravelVisaApplicantImportRequest) SetPhotoFileType(photoFileType string) error {
    r.photoFileType = photoFileType
    r.Set("photo_file_type", photoFileType)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetPhotoFileType() string {
    return r.photoFileType
}

func (r *AlitripTravelVisaApplicantImportRequest) SetFormDataJson(formDataJson string) error {
    r.formDataJson = formDataJson
    r.Set("form_data_json", formDataJson)
    return nil
}

func (r AlitripTravelVisaApplicantImportRequest) GetFormDataJson() string {
    return r.formDataJson
}


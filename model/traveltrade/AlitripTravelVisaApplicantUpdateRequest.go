package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-普通签证-申请人进度推进接口 APIRequest
alitrip.travel.visa.applicant.update

普通签证订单的申请人进度推进接口，用于商家代用户填写申请人基本信息 或 推进单个申请人的签证进度。
*/
type AlitripTravelVisaApplicantUpdateRequest struct {
    model.Params

    // 必填，子订单id
    subOrderId   string 

    // 必填，操作类型。1-上传新申请人基本信息（商家代填申请人），2-更新已有申请人基本信息，3-更新已有申请人的签证进度，4-商家代传申请人信息和材料(使馆直连订单)
    operType   int64 

    // 特殊必填，申请人基本信息（证件号，姓名，手机号）列表。当operType为1或2或4时必填
    applicantInfos   []NormalVisaApplicantInfo 

    // 特殊必填，签证申请人进度推进操作（目前只支持对单个申请人状态进行推进）。当operType为3时必填
    applicantOp   *NormalVisaApplicantOperation 

    // 特殊必填，pdf文件字节流。用于上传电子签pdf结果 或者 预约面试信pdf文件。
    fileBytes   []*model.File 

    // 特殊必填，文件字节流，用于上传证件照，必须和photoType同时传
    photoBytes   []*model.File 

    // 证件照文件类型
    photoType   string 

    // 特殊必填，文件字节流，用于上传护照，必须和passportType同时传
    passportBytes   []*model.File 

    // 护照文件类型
    passportType   string 

    // 酒店预订文件类型
    hotelBookingFormType   string 

    // 特殊必填，文件字节流，用于上传酒店预订，必须和hotelBookingFormType同时传
    hotelBookingFormBytes   []*model.File 

    // 机票预订文件类型
    flightBookingFormType   string 

    // 特殊必填，文件字节流，用于上传机票预订，必须和flightBookingFormType同时传
    flightBookingFormBytes   []*model.File 

    // 特殊必填，更多材料
    documentInfos   []NormalVisaDocumentInfo 

}

func NewAlitripTravelVisaApplicantUpdateRequest() *AlitripTravelVisaApplicantUpdateRequest{
    return &AlitripTravelVisaApplicantUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.applicant.update"
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelVisaApplicantUpdateRequest) SetSubOrderId(subOrderId string) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetSubOrderId() string {
    return r.subOrderId
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetOperType(operType int64) error {
    r.operType = operType
    r.Set("oper_type", operType)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetOperType() int64 {
    return r.operType
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetApplicantInfos(applicantInfos []NormalVisaApplicantInfo) error {
    r.applicantInfos = applicantInfos
    r.Set("applicant_infos", applicantInfos)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetApplicantInfos() []NormalVisaApplicantInfo {
    return r.applicantInfos
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetApplicantOp(applicantOp *NormalVisaApplicantOperation) error {
    r.applicantOp = applicantOp
    r.Set("applicant_op", applicantOp)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetApplicantOp() *NormalVisaApplicantOperation {
    return r.applicantOp
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetFileBytes(fileBytes []*model.File) error {
    r.fileBytes = fileBytes
    r.Set("file_bytes", fileBytes)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetFileBytes() []*model.File {
    return r.fileBytes
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetPhotoBytes(photoBytes []*model.File) error {
    r.photoBytes = photoBytes
    r.Set("photo_bytes", photoBytes)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetPhotoBytes() []*model.File {
    return r.photoBytes
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetPhotoType(photoType string) error {
    r.photoType = photoType
    r.Set("photo_type", photoType)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetPhotoType() string {
    return r.photoType
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetPassportBytes(passportBytes []*model.File) error {
    r.passportBytes = passportBytes
    r.Set("passport_bytes", passportBytes)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetPassportBytes() []*model.File {
    return r.passportBytes
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetPassportType(passportType string) error {
    r.passportType = passportType
    r.Set("passport_type", passportType)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetPassportType() string {
    return r.passportType
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetHotelBookingFormType(hotelBookingFormType string) error {
    r.hotelBookingFormType = hotelBookingFormType
    r.Set("hotel_booking_form_type", hotelBookingFormType)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetHotelBookingFormType() string {
    return r.hotelBookingFormType
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetHotelBookingFormBytes(hotelBookingFormBytes []*model.File) error {
    r.hotelBookingFormBytes = hotelBookingFormBytes
    r.Set("hotel_booking_form_bytes", hotelBookingFormBytes)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetHotelBookingFormBytes() []*model.File {
    return r.hotelBookingFormBytes
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetFlightBookingFormType(flightBookingFormType string) error {
    r.flightBookingFormType = flightBookingFormType
    r.Set("flight_booking_form_type", flightBookingFormType)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetFlightBookingFormType() string {
    return r.flightBookingFormType
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetFlightBookingFormBytes(flightBookingFormBytes []*model.File) error {
    r.flightBookingFormBytes = flightBookingFormBytes
    r.Set("flight_booking_form_bytes", flightBookingFormBytes)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetFlightBookingFormBytes() []*model.File {
    return r.flightBookingFormBytes
}

func (r *AlitripTravelVisaApplicantUpdateRequest) SetDocumentInfos(documentInfos []NormalVisaDocumentInfo) error {
    r.documentInfos = documentInfos
    r.Set("document_infos", documentInfos)
    return nil
}

func (r AlitripTravelVisaApplicantUpdateRequest) GetDocumentInfos() []NormalVisaDocumentInfo {
    return r.documentInfos
}


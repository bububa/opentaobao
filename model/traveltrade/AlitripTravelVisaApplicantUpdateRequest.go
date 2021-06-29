package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-普通签证-申请人进度推进接口 API请求
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

// 初始化AlitripTravelVisaApplicantUpdateRequest对象
func NewAlitripTravelVisaApplicantUpdateRequest() *AlitripTravelVisaApplicantUpdateRequest{
    return &AlitripTravelVisaApplicantUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaApplicantUpdateRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.applicant.update"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaApplicantUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubOrderId Setter
// 必填，子订单id
func (r *AlitripTravelVisaApplicantUpdateRequest) SetSubOrderId(subOrderId string) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetSubOrderId() string {
    return r.subOrderId
}
// OperType Setter
// 必填，操作类型。1-上传新申请人基本信息（商家代填申请人），2-更新已有申请人基本信息，3-更新已有申请人的签证进度，4-商家代传申请人信息和材料(使馆直连订单)
func (r *AlitripTravelVisaApplicantUpdateRequest) SetOperType(operType int64) error {
    r.operType = operType
    r.Set("oper_type", operType)
    return nil
}

// OperType Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetOperType() int64 {
    return r.operType
}
// ApplicantInfos Setter
// 特殊必填，申请人基本信息（证件号，姓名，手机号）列表。当operType为1或2或4时必填
func (r *AlitripTravelVisaApplicantUpdateRequest) SetApplicantInfos(applicantInfos []NormalVisaApplicantInfo) error {
    r.applicantInfos = applicantInfos
    r.Set("applicant_infos", applicantInfos)
    return nil
}

// ApplicantInfos Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetApplicantInfos() []NormalVisaApplicantInfo {
    return r.applicantInfos
}
// ApplicantOp Setter
// 特殊必填，签证申请人进度推进操作（目前只支持对单个申请人状态进行推进）。当operType为3时必填
func (r *AlitripTravelVisaApplicantUpdateRequest) SetApplicantOp(applicantOp *NormalVisaApplicantOperation) error {
    r.applicantOp = applicantOp
    r.Set("applicant_op", applicantOp)
    return nil
}

// ApplicantOp Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetApplicantOp() *NormalVisaApplicantOperation {
    return r.applicantOp
}
// FileBytes Setter
// 特殊必填，pdf文件字节流。用于上传电子签pdf结果 或者 预约面试信pdf文件。
func (r *AlitripTravelVisaApplicantUpdateRequest) SetFileBytes(fileBytes []*model.File) error {
    r.fileBytes = fileBytes
    r.Set("file_bytes", fileBytes)
    return nil
}

// FileBytes Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetFileBytes() []*model.File {
    return r.fileBytes
}
// PhotoBytes Setter
// 特殊必填，文件字节流，用于上传证件照，必须和photoType同时传
func (r *AlitripTravelVisaApplicantUpdateRequest) SetPhotoBytes(photoBytes []*model.File) error {
    r.photoBytes = photoBytes
    r.Set("photo_bytes", photoBytes)
    return nil
}

// PhotoBytes Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetPhotoBytes() []*model.File {
    return r.photoBytes
}
// PhotoType Setter
// 证件照文件类型
func (r *AlitripTravelVisaApplicantUpdateRequest) SetPhotoType(photoType string) error {
    r.photoType = photoType
    r.Set("photo_type", photoType)
    return nil
}

// PhotoType Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetPhotoType() string {
    return r.photoType
}
// PassportBytes Setter
// 特殊必填，文件字节流，用于上传护照，必须和passportType同时传
func (r *AlitripTravelVisaApplicantUpdateRequest) SetPassportBytes(passportBytes []*model.File) error {
    r.passportBytes = passportBytes
    r.Set("passport_bytes", passportBytes)
    return nil
}

// PassportBytes Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetPassportBytes() []*model.File {
    return r.passportBytes
}
// PassportType Setter
// 护照文件类型
func (r *AlitripTravelVisaApplicantUpdateRequest) SetPassportType(passportType string) error {
    r.passportType = passportType
    r.Set("passport_type", passportType)
    return nil
}

// PassportType Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetPassportType() string {
    return r.passportType
}
// HotelBookingFormType Setter
// 酒店预订文件类型
func (r *AlitripTravelVisaApplicantUpdateRequest) SetHotelBookingFormType(hotelBookingFormType string) error {
    r.hotelBookingFormType = hotelBookingFormType
    r.Set("hotel_booking_form_type", hotelBookingFormType)
    return nil
}

// HotelBookingFormType Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetHotelBookingFormType() string {
    return r.hotelBookingFormType
}
// HotelBookingFormBytes Setter
// 特殊必填，文件字节流，用于上传酒店预订，必须和hotelBookingFormType同时传
func (r *AlitripTravelVisaApplicantUpdateRequest) SetHotelBookingFormBytes(hotelBookingFormBytes []*model.File) error {
    r.hotelBookingFormBytes = hotelBookingFormBytes
    r.Set("hotel_booking_form_bytes", hotelBookingFormBytes)
    return nil
}

// HotelBookingFormBytes Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetHotelBookingFormBytes() []*model.File {
    return r.hotelBookingFormBytes
}
// FlightBookingFormType Setter
// 机票预订文件类型
func (r *AlitripTravelVisaApplicantUpdateRequest) SetFlightBookingFormType(flightBookingFormType string) error {
    r.flightBookingFormType = flightBookingFormType
    r.Set("flight_booking_form_type", flightBookingFormType)
    return nil
}

// FlightBookingFormType Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetFlightBookingFormType() string {
    return r.flightBookingFormType
}
// FlightBookingFormBytes Setter
// 特殊必填，文件字节流，用于上传机票预订，必须和flightBookingFormType同时传
func (r *AlitripTravelVisaApplicantUpdateRequest) SetFlightBookingFormBytes(flightBookingFormBytes []*model.File) error {
    r.flightBookingFormBytes = flightBookingFormBytes
    r.Set("flight_booking_form_bytes", flightBookingFormBytes)
    return nil
}

// FlightBookingFormBytes Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetFlightBookingFormBytes() []*model.File {
    return r.flightBookingFormBytes
}
// DocumentInfos Setter
// 特殊必填，更多材料
func (r *AlitripTravelVisaApplicantUpdateRequest) SetDocumentInfos(documentInfos []NormalVisaDocumentInfo) error {
    r.documentInfos = documentInfos
    r.Set("document_infos", documentInfos)
    return nil
}

// DocumentInfos Getter
func (r AlitripTravelVisaApplicantUpdateRequest) GetDocumentInfos() []NormalVisaDocumentInfo {
    return r.documentInfos
}

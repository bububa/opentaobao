package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_门店发布／更新 APIRequest
alibaba.alihealth.examination.hospital.publish

第三方B端有新的门店发布，或者老的门店更新的时候，使用这个接口
*/
type AlibabaAlihealthExaminationHospitalPublishRequest struct {
    model.Params

    // 门店简介
    detail   string 

    // 门店联系电话
    tel   string 

    // 门店所属城市
    cityName   string 

    // 门店城市code（国标）
    cityCode   string 

    // 操作类型: publish=发布，update=更新
    type   string 

    // 医院等级，三甲、
    keyWord   string 

    // “须知”使用下面note_category字段
    examNotice   string 

    // 门店位置经度高德 坐标系
    pointX   string 

    // 门店位置纬度高德 坐标系
    pointY   string 

    // 门店地址
    address   string 

    // 工作时间
    workTime   string 

    // 门店名称
    hospitalName   string 

    // 门店code，机构保证唯一
    hospitalCode   string 

    // 交通线路，通过\r\n 进行换行
    routes   string 

    // http://images.aliyun.com/image?id=123
    logo   string 

    // 是否支持在线报告。0:不支持;1:支持
    onlineReport   int64 

    // 社会统一信用代码
    socialCreditCode   string 

    // 线下报告获取说明（必填）
    reportWay   string 

    // 线上体检报告几天出具（如果有电子报告必填）
    reportWayOnline   string 

    // 环境图片(json字符串数组)，第一张是头图；（传图前先找运营同学要图片规范，别瞎传）
    envImgsUrl   string 

    // 免费停车场,绿色VIP通道,免费早餐,3天出报告,1V1导检,接待引导,独家签约,专家会诊,当天出报告；多个逗号分隔
    specialTagsCode   string 

    // 通知信息
    notify   string 

    // 不同种类的预约须知；
    noteCategory   string 

    // 经营模式 0自营模式、1平台模式
    mode   string 

    // 门店与医院协议
    agreement   string 

    // 营业执照
    businessLicense   string 

    // 医疗经营许可
    medicalLicense   string 

    // 类目:1=体检；2=核酸；3=上门；4=健康证；多个类目以逗号分割
    category   string 

}

func NewAlibabaAlihealthExaminationHospitalPublishRequest() *AlibabaAlihealthExaminationHospitalPublishRequest{
    return &AlibabaAlihealthExaminationHospitalPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.hospital.publish"
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetDetail(detail string) error {
    r.detail = detail
    r.Set("detail", detail)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetDetail() string {
    return r.detail
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetTel(tel string) error {
    r.tel = tel
    r.Set("tel", tel)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetTel() string {
    return r.tel
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetCityName(cityName string) error {
    r.cityName = cityName
    r.Set("city_name", cityName)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetCityName() string {
    return r.cityName
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetCityCode() string {
    return r.cityCode
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetType() string {
    return r.type
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetKeyWord(keyWord string) error {
    r.keyWord = keyWord
    r.Set("key_word", keyWord)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetKeyWord() string {
    return r.keyWord
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetExamNotice(examNotice string) error {
    r.examNotice = examNotice
    r.Set("exam_notice", examNotice)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetExamNotice() string {
    return r.examNotice
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetPointX(pointX string) error {
    r.pointX = pointX
    r.Set("point_x", pointX)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetPointX() string {
    return r.pointX
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetPointY(pointY string) error {
    r.pointY = pointY
    r.Set("point_y", pointY)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetPointY() string {
    return r.pointY
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetWorkTime(workTime string) error {
    r.workTime = workTime
    r.Set("work_time", workTime)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetWorkTime() string {
    return r.workTime
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetHospitalName(hospitalName string) error {
    r.hospitalName = hospitalName
    r.Set("hospital_name", hospitalName)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetHospitalName() string {
    return r.hospitalName
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetHospitalCode(hospitalCode string) error {
    r.hospitalCode = hospitalCode
    r.Set("hospital_code", hospitalCode)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetHospitalCode() string {
    return r.hospitalCode
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetRoutes(routes string) error {
    r.routes = routes
    r.Set("routes", routes)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetRoutes() string {
    return r.routes
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetLogo(logo string) error {
    r.logo = logo
    r.Set("logo", logo)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetLogo() string {
    return r.logo
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetOnlineReport(onlineReport int64) error {
    r.onlineReport = onlineReport
    r.Set("online_report", onlineReport)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetOnlineReport() int64 {
    return r.onlineReport
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetSocialCreditCode(socialCreditCode string) error {
    r.socialCreditCode = socialCreditCode
    r.Set("social_credit_code", socialCreditCode)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetSocialCreditCode() string {
    return r.socialCreditCode
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetReportWay(reportWay string) error {
    r.reportWay = reportWay
    r.Set("report_way", reportWay)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetReportWay() string {
    return r.reportWay
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetReportWayOnline(reportWayOnline string) error {
    r.reportWayOnline = reportWayOnline
    r.Set("report_way_online", reportWayOnline)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetReportWayOnline() string {
    return r.reportWayOnline
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetEnvImgsUrl(envImgsUrl string) error {
    r.envImgsUrl = envImgsUrl
    r.Set("env_imgs_url", envImgsUrl)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetEnvImgsUrl() string {
    return r.envImgsUrl
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetSpecialTagsCode(specialTagsCode string) error {
    r.specialTagsCode = specialTagsCode
    r.Set("special_tags_code", specialTagsCode)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetSpecialTagsCode() string {
    return r.specialTagsCode
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetNotify(notify string) error {
    r.notify = notify
    r.Set("notify", notify)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetNotify() string {
    return r.notify
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetNoteCategory(noteCategory string) error {
    r.noteCategory = noteCategory
    r.Set("note_category", noteCategory)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetNoteCategory() string {
    return r.noteCategory
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetMode(mode string) error {
    r.mode = mode
    r.Set("mode", mode)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetMode() string {
    return r.mode
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetAgreement(agreement string) error {
    r.agreement = agreement
    r.Set("agreement", agreement)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetAgreement() string {
    return r.agreement
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetBusinessLicense(businessLicense string) error {
    r.businessLicense = businessLicense
    r.Set("business_license", businessLicense)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetBusinessLicense() string {
    return r.businessLicense
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetMedicalLicense(medicalLicense string) error {
    r.medicalLicense = medicalLicense
    r.Set("medical_license", medicalLicense)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetMedicalLicense() string {
    return r.medicalLicense
}

func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetCategory(category string) error {
    r.category = category
    r.Set("category", category)
    return nil
}

func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetCategory() string {
    return r.category
}


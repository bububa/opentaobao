package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_门店发布／更新 API请求
alibaba.alihealth.examination.hospital.publish

第三方B端有新的门店发布，或者老的门店更新的时候，使用这个接口
*/
type AlibabaAlihealthExaminationHospitalPublishRequest struct {
    model.Params
    // 门店简介
    _detail   string
    // 门店联系电话
    _tel   string
    // 门店所属城市
    _cityName   string
    // 门店城市code（国标）
    _cityCode   string
    // 操作类型: publish=发布，update=更新
    _type   string
    // 医院等级，三甲、
    _keyWord   string
    // “须知”使用下面note_category字段
    _examNotice   string
    // 门店位置经度高德 坐标系
    _pointX   string
    // 门店位置纬度高德 坐标系
    _pointY   string
    // 门店地址
    _address   string
    // 工作时间
    _workTime   string
    // 门店名称
    _hospitalName   string
    // 门店code，机构保证唯一
    _hospitalCode   string
    // 交通线路，通过\r\n 进行换行
    _routes   string
    // http://images.aliyun.com/image?id=123
    _logo   string
    // 是否支持在线报告。0:不支持;1:支持
    _onlineReport   int64
    // 社会统一信用代码
    _socialCreditCode   string
    // 线下报告获取说明（必填）
    _reportWay   string
    // 线上体检报告几天出具（如果有电子报告必填）
    _reportWayOnline   string
    // 环境图片(json字符串数组)，第一张是头图；（传图前先找运营同学要图片规范，别瞎传）
    _envImgsUrl   string
    // 免费停车场,绿色VIP通道,免费早餐,3天出报告,1V1导检,接待引导,独家签约,专家会诊,当天出报告；多个逗号分隔
    _specialTagsCode   string
    // 通知信息
    _notify   string
    // 不同种类的预约须知；
    _noteCategory   string
    // 经营模式 0自营模式、1平台模式
    _mode   string
    // 门店与医院协议
    _agreement   string
    // 营业执照
    _businessLicense   string
    // 医疗经营许可
    _medicalLicense   string
    // 类目:1=体检；2=核酸；3=上门；4=健康证；多个类目以逗号分割
    _category   string
}

// 初始化AlibabaAlihealthExaminationHospitalPublishRequest对象
func NewAlibabaAlihealthExaminationHospitalPublishRequest() *AlibabaAlihealthExaminationHospitalPublishRequest{
    return &AlibabaAlihealthExaminationHospitalPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.hospital.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Detail Setter
// 门店简介
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetDetail(_detail string) error {
    r._detail = _detail
    r.Set("detail", _detail)
    return nil
}

// Detail Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetDetail() string {
    return r._detail
}
// Tel Setter
// 门店联系电话
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetTel(_tel string) error {
    r._tel = _tel
    r.Set("tel", _tel)
    return nil
}

// Tel Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetTel() string {
    return r._tel
}
// CityName Setter
// 门店所属城市
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetCityName(_cityName string) error {
    r._cityName = _cityName
    r.Set("city_name", _cityName)
    return nil
}

// CityName Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetCityName() string {
    return r._cityName
}
// CityCode Setter
// 门店城市code（国标）
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetCityCode() string {
    return r._cityCode
}
// Type Setter
// 操作类型: publish=发布，update=更新
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetType() string {
    return r._type
}
// KeyWord Setter
// 医院等级，三甲、
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetKeyWord(_keyWord string) error {
    r._keyWord = _keyWord
    r.Set("key_word", _keyWord)
    return nil
}

// KeyWord Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetKeyWord() string {
    return r._keyWord
}
// ExamNotice Setter
// “须知”使用下面note_category字段
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetExamNotice(_examNotice string) error {
    r._examNotice = _examNotice
    r.Set("exam_notice", _examNotice)
    return nil
}

// ExamNotice Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetExamNotice() string {
    return r._examNotice
}
// PointX Setter
// 门店位置经度高德 坐标系
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetPointX(_pointX string) error {
    r._pointX = _pointX
    r.Set("point_x", _pointX)
    return nil
}

// PointX Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetPointX() string {
    return r._pointX
}
// PointY Setter
// 门店位置纬度高德 坐标系
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetPointY(_pointY string) error {
    r._pointY = _pointY
    r.Set("point_y", _pointY)
    return nil
}

// PointY Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetPointY() string {
    return r._pointY
}
// Address Setter
// 门店地址
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetAddress() string {
    return r._address
}
// WorkTime Setter
// 工作时间
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetWorkTime(_workTime string) error {
    r._workTime = _workTime
    r.Set("work_time", _workTime)
    return nil
}

// WorkTime Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetWorkTime() string {
    return r._workTime
}
// HospitalName Setter
// 门店名称
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetHospitalName(_hospitalName string) error {
    r._hospitalName = _hospitalName
    r.Set("hospital_name", _hospitalName)
    return nil
}

// HospitalName Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetHospitalName() string {
    return r._hospitalName
}
// HospitalCode Setter
// 门店code，机构保证唯一
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetHospitalCode(_hospitalCode string) error {
    r._hospitalCode = _hospitalCode
    r.Set("hospital_code", _hospitalCode)
    return nil
}

// HospitalCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetHospitalCode() string {
    return r._hospitalCode
}
// Routes Setter
// 交通线路，通过\r\n 进行换行
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetRoutes(_routes string) error {
    r._routes = _routes
    r.Set("routes", _routes)
    return nil
}

// Routes Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetRoutes() string {
    return r._routes
}
// Logo Setter
// http://images.aliyun.com/image?id=123
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetLogo(_logo string) error {
    r._logo = _logo
    r.Set("logo", _logo)
    return nil
}

// Logo Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetLogo() string {
    return r._logo
}
// OnlineReport Setter
// 是否支持在线报告。0:不支持;1:支持
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetOnlineReport(_onlineReport int64) error {
    r._onlineReport = _onlineReport
    r.Set("online_report", _onlineReport)
    return nil
}

// OnlineReport Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetOnlineReport() int64 {
    return r._onlineReport
}
// SocialCreditCode Setter
// 社会统一信用代码
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetSocialCreditCode(_socialCreditCode string) error {
    r._socialCreditCode = _socialCreditCode
    r.Set("social_credit_code", _socialCreditCode)
    return nil
}

// SocialCreditCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetSocialCreditCode() string {
    return r._socialCreditCode
}
// ReportWay Setter
// 线下报告获取说明（必填）
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetReportWay(_reportWay string) error {
    r._reportWay = _reportWay
    r.Set("report_way", _reportWay)
    return nil
}

// ReportWay Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetReportWay() string {
    return r._reportWay
}
// ReportWayOnline Setter
// 线上体检报告几天出具（如果有电子报告必填）
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetReportWayOnline(_reportWayOnline string) error {
    r._reportWayOnline = _reportWayOnline
    r.Set("report_way_online", _reportWayOnline)
    return nil
}

// ReportWayOnline Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetReportWayOnline() string {
    return r._reportWayOnline
}
// EnvImgsUrl Setter
// 环境图片(json字符串数组)，第一张是头图；（传图前先找运营同学要图片规范，别瞎传）
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetEnvImgsUrl(_envImgsUrl string) error {
    r._envImgsUrl = _envImgsUrl
    r.Set("env_imgs_url", _envImgsUrl)
    return nil
}

// EnvImgsUrl Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetEnvImgsUrl() string {
    return r._envImgsUrl
}
// SpecialTagsCode Setter
// 免费停车场,绿色VIP通道,免费早餐,3天出报告,1V1导检,接待引导,独家签约,专家会诊,当天出报告；多个逗号分隔
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetSpecialTagsCode(_specialTagsCode string) error {
    r._specialTagsCode = _specialTagsCode
    r.Set("special_tags_code", _specialTagsCode)
    return nil
}

// SpecialTagsCode Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetSpecialTagsCode() string {
    return r._specialTagsCode
}
// Notify Setter
// 通知信息
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetNotify(_notify string) error {
    r._notify = _notify
    r.Set("notify", _notify)
    return nil
}

// Notify Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetNotify() string {
    return r._notify
}
// NoteCategory Setter
// 不同种类的预约须知；
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetNoteCategory(_noteCategory string) error {
    r._noteCategory = _noteCategory
    r.Set("note_category", _noteCategory)
    return nil
}

// NoteCategory Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetNoteCategory() string {
    return r._noteCategory
}
// Mode Setter
// 经营模式 0自营模式、1平台模式
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetMode(_mode string) error {
    r._mode = _mode
    r.Set("mode", _mode)
    return nil
}

// Mode Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetMode() string {
    return r._mode
}
// Agreement Setter
// 门店与医院协议
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetAgreement(_agreement string) error {
    r._agreement = _agreement
    r.Set("agreement", _agreement)
    return nil
}

// Agreement Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetAgreement() string {
    return r._agreement
}
// BusinessLicense Setter
// 营业执照
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetBusinessLicense(_businessLicense string) error {
    r._businessLicense = _businessLicense
    r.Set("business_license", _businessLicense)
    return nil
}

// BusinessLicense Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetBusinessLicense() string {
    return r._businessLicense
}
// MedicalLicense Setter
// 医疗经营许可
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetMedicalLicense(_medicalLicense string) error {
    r._medicalLicense = _medicalLicense
    r.Set("medical_license", _medicalLicense)
    return nil
}

// MedicalLicense Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetMedicalLicense() string {
    return r._medicalLicense
}
// Category Setter
// 类目:1=体检；2=核酸；3=上门；4=健康证；多个类目以逗号分割
func (r *AlibabaAlihealthExaminationHospitalPublishRequest) SetCategory(_category string) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r AlibabaAlihealthExaminationHospitalPublishRequest) GetCategory() string {
    return r._category
}

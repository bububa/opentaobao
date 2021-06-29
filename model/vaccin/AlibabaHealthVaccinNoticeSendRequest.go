package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送消息提醒 API请求
alibaba.health.vaccin.notice.send

ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
*/
type AlibabaHealthVaccinNoticeSendRequest struct {
    model.Params
    // 支付宝用户 ID
    alipayUserId   string
    // ISV 侧用户 ID
    isvUserId   string
    // 订单 ID
    orderId   string
    // 接种人姓名
    name   string
    // 联系电话
    mobile   string
    // 接种的疫苗信息
    vaccineList   []VaccineInfo
    // 接种人性别:1=男,2=女
    sex   int64
    // 接种人出生日期
    birthday   string
    // 接种点编码
    povNo   string
    // 接种点名称
    povName   string
    // 接种点地址
    address   string
    // 省份名称
    province   string
    // 城市名称
    city   string
    // 区名称
    area   string
    // 预约日期
    reserveDate   string
    // 预约时间段
    reserveTime   string
    // 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
    messageType   int64
    // 用户入口，支付宝或医鹿，alipay或yl
    appChannel   string
}

// 初始化AlibabaHealthVaccinNoticeSendRequest对象
func NewAlibabaHealthVaccinNoticeSendRequest() *AlibabaHealthVaccinNoticeSendRequest{
    return &AlibabaHealthVaccinNoticeSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeSendRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝用户 ID
func (r *AlibabaHealthVaccinNoticeSendRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeSendRequest) SetIsvUserId(isvUserId string) error {
    r.isvUserId = isvUserId
    r.Set("isv_user_id", isvUserId)
    return nil
}

// IsvUserId Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetIsvUserId() string {
    return r.isvUserId
}
// OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinNoticeSendRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetOrderId() string {
    return r.orderId
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeSendRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetName() string {
    return r.name
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeSendRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetMobile() string {
    return r.mobile
}
// VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinNoticeSendRequest) SetVaccineList(vaccineList []VaccineInfo) error {
    r.vaccineList = vaccineList
    r.Set("vaccine_list", vaccineList)
    return nil
}

// VaccineList Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetVaccineList() []VaccineInfo {
    return r.vaccineList
}
// Sex Setter
// 接种人性别:1=男,2=女
func (r *AlibabaHealthVaccinNoticeSendRequest) SetSex(sex int64) error {
    r.sex = sex
    r.Set("sex", sex)
    return nil
}

// Sex Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetSex() int64 {
    return r.sex
}
// Birthday Setter
// 接种人出生日期
func (r *AlibabaHealthVaccinNoticeSendRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

// Birthday Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetBirthday() string {
    return r.birthday
}
// PovNo Setter
// 接种点编码
func (r *AlibabaHealthVaccinNoticeSendRequest) SetPovNo(povNo string) error {
    r.povNo = povNo
    r.Set("pov_no", povNo)
    return nil
}

// PovNo Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetPovNo() string {
    return r.povNo
}
// PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetPovName(povName string) error {
    r.povName = povName
    r.Set("pov_name", povName)
    return nil
}

// PovName Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetPovName() string {
    return r.povName
}
// Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeSendRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetAddress() string {
    return r.address
}
// Province Setter
// 省份名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetProvince() string {
    return r.province
}
// City Setter
// 城市名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetCity() string {
    return r.city
}
// Area Setter
// 区名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetArea() string {
    return r.area
}
// ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeSendRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetReserveDate() string {
    return r.reserveDate
}
// ReserveTime Setter
// 预约时间段
func (r *AlibabaHealthVaccinNoticeSendRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetReserveTime() string {
    return r.reserveTime
}
// MessageType Setter
// 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
func (r *AlibabaHealthVaccinNoticeSendRequest) SetMessageType(messageType int64) error {
    r.messageType = messageType
    r.Set("message_type", messageType)
    return nil
}

// MessageType Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetMessageType() int64 {
    return r.messageType
}
// AppChannel Setter
// 用户入口，支付宝或医鹿，alipay或yl
func (r *AlibabaHealthVaccinNoticeSendRequest) SetAppChannel(appChannel string) error {
    r.appChannel = appChannel
    r.Set("app_channel", appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetAppChannel() string {
    return r.appChannel
}

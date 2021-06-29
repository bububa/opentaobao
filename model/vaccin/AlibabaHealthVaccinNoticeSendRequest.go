package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送消息提醒 APIRequest
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

func NewAlibabaHealthVaccinNoticeSendRequest() *AlibabaHealthVaccinNoticeSendRequest{
    return &AlibabaHealthVaccinNoticeSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.send"
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeSendRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetIsvUserId(isvUserId string) error {
    r.isvUserId = isvUserId
    r.Set("isv_user_id", isvUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetIsvUserId() string {
    return r.isvUserId
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetName() string {
    return r.name
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetMobile() string {
    return r.mobile
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetVaccineList(vaccineList []VaccineInfo) error {
    r.vaccineList = vaccineList
    r.Set("vaccine_list", vaccineList)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetVaccineList() []VaccineInfo {
    return r.vaccineList
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetSex(sex int64) error {
    r.sex = sex
    r.Set("sex", sex)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetSex() int64 {
    return r.sex
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetBirthday() string {
    return r.birthday
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetPovNo(povNo string) error {
    r.povNo = povNo
    r.Set("pov_no", povNo)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetPovNo() string {
    return r.povNo
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetPovName(povName string) error {
    r.povName = povName
    r.Set("pov_name", povName)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetPovName() string {
    return r.povName
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetProvince() string {
    return r.province
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetCity() string {
    return r.city
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetArea() string {
    return r.area
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetReserveDate() string {
    return r.reserveDate
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetReserveTime() string {
    return r.reserveTime
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetMessageType(messageType int64) error {
    r.messageType = messageType
    r.Set("message_type", messageType)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetMessageType() int64 {
    return r.messageType
}

func (r *AlibabaHealthVaccinNoticeSendRequest) SetAppChannel(appChannel string) error {
    r.appChannel = appChannel
    r.Set("app_channel", appChannel)
    return nil
}

func (r AlibabaHealthVaccinNoticeSendRequest) GetAppChannel() string {
    return r.appChannel
}


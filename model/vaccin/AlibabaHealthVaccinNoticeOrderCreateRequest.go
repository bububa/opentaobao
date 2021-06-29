package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝医疗健康疫苗预约创建 APIRequest
alibaba.health.vaccin.notice.order.create

支付宝医疗健康疫苗预约创建
*/
type AlibabaHealthVaccinNoticeOrderCreateRequest struct {
    model.Params

    // 预约人性别(1男2女)
    sex   int64 

    // 年龄
    age   int64 

    // 预约日期
    reserveDate   string 

    // 支付宝用户id
    alipayUserId   string 

    // 外部渠道用户id
    outerUserId   string 

    // 预约id
    orderId   string 

    // 手机号码
    mobile   string 

    // 接种人姓名
    name   string 

    // 接种点地址
    address   string 

    // 接种点名称
    povStoreName   string 

    // 预约时间
    reserveTime   string 

    // 疫苗信息
    vaccineInfo   string 

    // 年龄类型(1-宝宝2-成人)
    ageType   int64 

    // 支付宝消息通知跳转订单详情链接
    orderDetailUrl   string 

    // 地区名字
    area   string 

    // 城市名字
    city   string 

    // 省份名字
    province   string 

}

func NewAlibabaHealthVaccinNoticeOrderCreateRequest() *AlibabaHealthVaccinNoticeOrderCreateRequest{
    return &AlibabaHealthVaccinNoticeOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.create"
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetSex(sex int64) error {
    r.sex = sex
    r.Set("sex", sex)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetSex() int64 {
    return r.sex
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAge() int64 {
    return r.age
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetReserveDate() string {
    return r.reserveDate
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOuterUserId(outerUserId string) error {
    r.outerUserId = outerUserId
    r.Set("outer_user_id", outerUserId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOuterUserId() string {
    return r.outerUserId
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetMobile() string {
    return r.mobile
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetName() string {
    return r.name
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetPovStoreName(povStoreName string) error {
    r.povStoreName = povStoreName
    r.Set("pov_store_name", povStoreName)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetPovStoreName() string {
    return r.povStoreName
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetReserveTime() string {
    return r.reserveTime
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetVaccineInfo(vaccineInfo string) error {
    r.vaccineInfo = vaccineInfo
    r.Set("vaccine_info", vaccineInfo)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetVaccineInfo() string {
    return r.vaccineInfo
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAgeType(ageType int64) error {
    r.ageType = ageType
    r.Set("age_type", ageType)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAgeType() int64 {
    return r.ageType
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOrderDetailUrl(orderDetailUrl string) error {
    r.orderDetailUrl = orderDetailUrl
    r.Set("order_detail_url", orderDetailUrl)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOrderDetailUrl() string {
    return r.orderDetailUrl
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetArea() string {
    return r.area
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetCity() string {
    return r.city
}

func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetProvince() string {
    return r.province
}


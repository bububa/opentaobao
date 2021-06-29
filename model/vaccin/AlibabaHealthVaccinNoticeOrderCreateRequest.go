package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝医疗健康疫苗预约创建 API请求
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

// 初始化AlibabaHealthVaccinNoticeOrderCreateRequest对象
func NewAlibabaHealthVaccinNoticeOrderCreateRequest() *AlibabaHealthVaccinNoticeOrderCreateRequest{
    return &AlibabaHealthVaccinNoticeOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Sex Setter
// 预约人性别(1男2女)
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetSex(sex int64) error {
    r.sex = sex
    r.Set("sex", sex)
    return nil
}

// Sex Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetSex() int64 {
    return r.sex
}
// Age Setter
// 年龄
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

// Age Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAge() int64 {
    return r.age
}
// ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetReserveDate() string {
    return r.reserveDate
}
// AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// OuterUserId Setter
// 外部渠道用户id
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOuterUserId(outerUserId string) error {
    r.outerUserId = outerUserId
    r.Set("outer_user_id", outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOuterUserId() string {
    return r.outerUserId
}
// OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOrderId() string {
    return r.orderId
}
// Mobile Setter
// 手机号码
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetMobile() string {
    return r.mobile
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetName() string {
    return r.name
}
// Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAddress() string {
    return r.address
}
// PovStoreName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetPovStoreName(povStoreName string) error {
    r.povStoreName = povStoreName
    r.Set("pov_store_name", povStoreName)
    return nil
}

// PovStoreName Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetPovStoreName() string {
    return r.povStoreName
}
// ReserveTime Setter
// 预约时间
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetReserveTime(reserveTime string) error {
    r.reserveTime = reserveTime
    r.Set("reserve_time", reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetReserveTime() string {
    return r.reserveTime
}
// VaccineInfo Setter
// 疫苗信息
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetVaccineInfo(vaccineInfo string) error {
    r.vaccineInfo = vaccineInfo
    r.Set("vaccine_info", vaccineInfo)
    return nil
}

// VaccineInfo Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetVaccineInfo() string {
    return r.vaccineInfo
}
// AgeType Setter
// 年龄类型(1-宝宝2-成人)
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAgeType(ageType int64) error {
    r.ageType = ageType
    r.Set("age_type", ageType)
    return nil
}

// AgeType Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAgeType() int64 {
    return r.ageType
}
// OrderDetailUrl Setter
// 支付宝消息通知跳转订单详情链接
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOrderDetailUrl(orderDetailUrl string) error {
    r.orderDetailUrl = orderDetailUrl
    r.Set("order_detail_url", orderDetailUrl)
    return nil
}

// OrderDetailUrl Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOrderDetailUrl() string {
    return r.orderDetailUrl
}
// Area Setter
// 地区名字
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetArea() string {
    return r.area
}
// City Setter
// 城市名字
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetCity() string {
    return r.city
}
// Province Setter
// 省份名字
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetProvince() string {
    return r.province
}

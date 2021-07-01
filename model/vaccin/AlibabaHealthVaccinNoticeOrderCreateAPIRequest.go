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
type AlibabaHealthVaccinNoticeOrderCreateAPIRequest struct {
    model.Params
    // 预约人性别(1男2女)
    _sex   int64
    // 年龄
    _age   int64
    // 预约日期
    _reserveDate   string
    // 支付宝用户id
    _alipayUserId   string
    // 外部渠道用户id
    _outerUserId   string
    // 预约id
    _orderId   string
    // 手机号码
    _mobile   string
    // 接种人姓名
    _name   string
    // 接种点地址
    _address   string
    // 接种点名称
    _povStoreName   string
    // 预约时间
    _reserveTime   string
    // 疫苗信息
    _vaccineInfo   string
    // 年龄类型(1-宝宝2-成人)
    _ageType   int64
    // 支付宝消息通知跳转订单详情链接
    _orderDetailUrl   string
    // 地区名字
    _area   string
    // 城市名字
    _city   string
    // 省份名字
    _province   string
}

// 初始化AlibabaHealthVaccinNoticeOrderCreateAPIRequest对象
func NewAlibabaHealthVaccinNoticeOrderCreateRequest() *AlibabaHealthVaccinNoticeOrderCreateAPIRequest{
    return &AlibabaHealthVaccinNoticeOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Sex Setter
// 预约人性别(1男2女)
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetSex(_sex int64) error {
    r._sex = _sex
    r.Set("sex", _sex)
    return nil
}

// Sex Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetSex() int64 {
    return r._sex
}
// Age Setter
// 年龄
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetAge(_age int64) error {
    r._age = _age
    r.Set("age", _age)
    return nil
}

// Age Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetAge() int64 {
    return r._age
}
// ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetReserveDate() string {
    return r._reserveDate
}
// AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// OuterUserId Setter
// 外部渠道用户id
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetOuterUserId(_outerUserId string) error {
    r._outerUserId = _outerUserId
    r.Set("outer_user_id", _outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetOuterUserId() string {
    return r._outerUserId
}
// OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetOrderId() string {
    return r._orderId
}
// Mobile Setter
// 手机号码
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetMobile() string {
    return r._mobile
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetName() string {
    return r._name
}
// Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetAddress() string {
    return r._address
}
// PovStoreName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetPovStoreName(_povStoreName string) error {
    r._povStoreName = _povStoreName
    r.Set("pov_store_name", _povStoreName)
    return nil
}

// PovStoreName Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetPovStoreName() string {
    return r._povStoreName
}
// ReserveTime Setter
// 预约时间
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetReserveTime(_reserveTime string) error {
    r._reserveTime = _reserveTime
    r.Set("reserve_time", _reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetReserveTime() string {
    return r._reserveTime
}
// VaccineInfo Setter
// 疫苗信息
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetVaccineInfo(_vaccineInfo string) error {
    r._vaccineInfo = _vaccineInfo
    r.Set("vaccine_info", _vaccineInfo)
    return nil
}

// VaccineInfo Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetVaccineInfo() string {
    return r._vaccineInfo
}
// AgeType Setter
// 年龄类型(1-宝宝2-成人)
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetAgeType(_ageType int64) error {
    r._ageType = _ageType
    r.Set("age_type", _ageType)
    return nil
}

// AgeType Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetAgeType() int64 {
    return r._ageType
}
// OrderDetailUrl Setter
// 支付宝消息通知跳转订单详情链接
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetOrderDetailUrl(_orderDetailUrl string) error {
    r._orderDetailUrl = _orderDetailUrl
    r.Set("order_detail_url", _orderDetailUrl)
    return nil
}

// OrderDetailUrl Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetOrderDetailUrl() string {
    return r._orderDetailUrl
}
// Area Setter
// 地区名字
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetArea(_area string) error {
    r._area = _area
    r.Set("area", _area)
    return nil
}

// Area Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetArea() string {
    return r._area
}
// City Setter
// 城市名字
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetCity() string {
    return r._city
}
// Province Setter
// 省份名字
func (r *AlibabaHealthVaccinNoticeOrderCreateAPIRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaHealthVaccinNoticeOrderCreateAPIRequest) GetProvince() string {
    return r._province
}

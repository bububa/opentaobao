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
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetSex(_sex int64) error {
    r._sex = _sex
    r.Set("sex", _sex)
    return nil
}

// Sex Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetSex() int64 {
    return r._sex
}
// Age Setter
// 年龄
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAge(_age int64) error {
    r._age = _age
    r.Set("age", _age)
    return nil
}

// Age Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAge() int64 {
    return r._age
}
// ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetReserveDate() string {
    return r._reserveDate
}
// AlipayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// OuterUserId Setter
// 外部渠道用户id
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOuterUserId(_outerUserId string) error {
    r._outerUserId = _outerUserId
    r.Set("outer_user_id", _outerUserId)
    return nil
}

// OuterUserId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOuterUserId() string {
    return r._outerUserId
}
// OrderId Setter
// 预约id
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOrderId() string {
    return r._orderId
}
// Mobile Setter
// 手机号码
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetMobile() string {
    return r._mobile
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetName() string {
    return r._name
}
// Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAddress() string {
    return r._address
}
// PovStoreName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetPovStoreName(_povStoreName string) error {
    r._povStoreName = _povStoreName
    r.Set("pov_store_name", _povStoreName)
    return nil
}

// PovStoreName Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetPovStoreName() string {
    return r._povStoreName
}
// ReserveTime Setter
// 预约时间
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetReserveTime(_reserveTime string) error {
    r._reserveTime = _reserveTime
    r.Set("reserve_time", _reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetReserveTime() string {
    return r._reserveTime
}
// VaccineInfo Setter
// 疫苗信息
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetVaccineInfo(_vaccineInfo string) error {
    r._vaccineInfo = _vaccineInfo
    r.Set("vaccine_info", _vaccineInfo)
    return nil
}

// VaccineInfo Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetVaccineInfo() string {
    return r._vaccineInfo
}
// AgeType Setter
// 年龄类型(1-宝宝2-成人)
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetAgeType(_ageType int64) error {
    r._ageType = _ageType
    r.Set("age_type", _ageType)
    return nil
}

// AgeType Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetAgeType() int64 {
    return r._ageType
}
// OrderDetailUrl Setter
// 支付宝消息通知跳转订单详情链接
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetOrderDetailUrl(_orderDetailUrl string) error {
    r._orderDetailUrl = _orderDetailUrl
    r.Set("order_detail_url", _orderDetailUrl)
    return nil
}

// OrderDetailUrl Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetOrderDetailUrl() string {
    return r._orderDetailUrl
}
// Area Setter
// 地区名字
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetArea(_area string) error {
    r._area = _area
    r.Set("area", _area)
    return nil
}

// Area Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetArea() string {
    return r._area
}
// City Setter
// 城市名字
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetCity() string {
    return r._city
}
// Province Setter
// 省份名字
func (r *AlibabaHealthVaccinNoticeOrderCreateRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaHealthVaccinNoticeOrderCreateRequest) GetProvince() string {
    return r._province
}

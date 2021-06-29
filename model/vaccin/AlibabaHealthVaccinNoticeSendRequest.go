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
    _alipayUserId   string
    // ISV 侧用户 ID
    _isvUserId   string
    // 订单 ID
    _orderId   string
    // 接种人姓名
    _name   string
    // 联系电话
    _mobile   string
    // 接种的疫苗信息
    _vaccineList   []VaccineInfo
    // 接种人性别:1=男,2=女
    _sex   int64
    // 接种人出生日期
    _birthday   string
    // 接种点编码
    _povNo   string
    // 接种点名称
    _povName   string
    // 接种点地址
    _address   string
    // 省份名称
    _province   string
    // 城市名称
    _city   string
    // 区名称
    _area   string
    // 预约日期
    _reserveDate   string
    // 预约时间段
    _reserveTime   string
    // 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
    _messageType   int64
    // 用户入口，支付宝或医鹿，alipay或yl
    _appChannel   string
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
func (r *AlibabaHealthVaccinNoticeSendRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeSendRequest) SetIsvUserId(_isvUserId string) error {
    r._isvUserId = _isvUserId
    r.Set("isv_user_id", _isvUserId)
    return nil
}

// IsvUserId Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetIsvUserId() string {
    return r._isvUserId
}
// OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinNoticeSendRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetOrderId() string {
    return r._orderId
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeSendRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetName() string {
    return r._name
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeSendRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetMobile() string {
    return r._mobile
}
// VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinNoticeSendRequest) SetVaccineList(_vaccineList []VaccineInfo) error {
    r._vaccineList = _vaccineList
    r.Set("vaccine_list", _vaccineList)
    return nil
}

// VaccineList Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetVaccineList() []VaccineInfo {
    return r._vaccineList
}
// Sex Setter
// 接种人性别:1=男,2=女
func (r *AlibabaHealthVaccinNoticeSendRequest) SetSex(_sex int64) error {
    r._sex = _sex
    r.Set("sex", _sex)
    return nil
}

// Sex Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetSex() int64 {
    return r._sex
}
// Birthday Setter
// 接种人出生日期
func (r *AlibabaHealthVaccinNoticeSendRequest) SetBirthday(_birthday string) error {
    r._birthday = _birthday
    r.Set("birthday", _birthday)
    return nil
}

// Birthday Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetBirthday() string {
    return r._birthday
}
// PovNo Setter
// 接种点编码
func (r *AlibabaHealthVaccinNoticeSendRequest) SetPovNo(_povNo string) error {
    r._povNo = _povNo
    r.Set("pov_no", _povNo)
    return nil
}

// PovNo Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetPovNo() string {
    return r._povNo
}
// PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetPovName(_povName string) error {
    r._povName = _povName
    r.Set("pov_name", _povName)
    return nil
}

// PovName Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetPovName() string {
    return r._povName
}
// Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeSendRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetAddress() string {
    return r._address
}
// Province Setter
// 省份名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetProvince() string {
    return r._province
}
// City Setter
// 城市名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetCity() string {
    return r._city
}
// Area Setter
// 区名称
func (r *AlibabaHealthVaccinNoticeSendRequest) SetArea(_area string) error {
    r._area = _area
    r.Set("area", _area)
    return nil
}

// Area Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetArea() string {
    return r._area
}
// ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeSendRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetReserveDate() string {
    return r._reserveDate
}
// ReserveTime Setter
// 预约时间段
func (r *AlibabaHealthVaccinNoticeSendRequest) SetReserveTime(_reserveTime string) error {
    r._reserveTime = _reserveTime
    r.Set("reserve_time", _reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetReserveTime() string {
    return r._reserveTime
}
// MessageType Setter
// 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
func (r *AlibabaHealthVaccinNoticeSendRequest) SetMessageType(_messageType int64) error {
    r._messageType = _messageType
    r.Set("message_type", _messageType)
    return nil
}

// MessageType Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetMessageType() int64 {
    return r._messageType
}
// AppChannel Setter
// 用户入口，支付宝或医鹿，alipay或yl
func (r *AlibabaHealthVaccinNoticeSendRequest) SetAppChannel(_appChannel string) error {
    r._appChannel = _appChannel
    r.Set("app_channel", _appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaHealthVaccinNoticeSendRequest) GetAppChannel() string {
    return r._appChannel
}

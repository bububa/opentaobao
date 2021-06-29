package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接种完成反馈接口 API请求
alibaba.health.vaccin.vaccinate.complete

ISV 将用户完成接种的疫苗同步给免疫规划中心
*/
type AlibabaHealthVaccinVaccinateCompleteRequest struct {
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
    // 接种日期
    _vaccinateDate   string
    // 接种时间
    _vaccinateTime   string
    // 接种的疫苗信息
    _vaccineList   []VaccineInfo
}

// 初始化AlibabaHealthVaccinVaccinateCompleteRequest对象
func NewAlibabaHealthVaccinVaccinateCompleteRequest() *AlibabaHealthVaccinVaccinateCompleteRequest{
    return &AlibabaHealthVaccinVaccinateCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.vaccinate.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝用户 ID
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetIsvUserId(_isvUserId string) error {
    r._isvUserId = _isvUserId
    r.Set("isv_user_id", _isvUserId)
    return nil
}

// IsvUserId Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetIsvUserId() string {
    return r._isvUserId
}
// OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetOrderId() string {
    return r._orderId
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetName() string {
    return r._name
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetMobile() string {
    return r._mobile
}
// VaccinateDate Setter
// 接种日期
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetVaccinateDate(_vaccinateDate string) error {
    r._vaccinateDate = _vaccinateDate
    r.Set("vaccinate_date", _vaccinateDate)
    return nil
}

// VaccinateDate Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetVaccinateDate() string {
    return r._vaccinateDate
}
// VaccinateTime Setter
// 接种时间
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetVaccinateTime(_vaccinateTime string) error {
    r._vaccinateTime = _vaccinateTime
    r.Set("vaccinate_time", _vaccinateTime)
    return nil
}

// VaccinateTime Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetVaccinateTime() string {
    return r._vaccinateTime
}
// VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetVaccineList(_vaccineList []VaccineInfo) error {
    r._vaccineList = _vaccineList
    r.Set("vaccine_list", _vaccineList)
    return nil
}

// VaccineList Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetVaccineList() []VaccineInfo {
    return r._vaccineList
}

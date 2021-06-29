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
    alipayUserId   string
    // ISV 侧用户 ID
    isvUserId   string
    // 订单 ID
    orderId   string
    // 接种人姓名
    name   string
    // 联系电话
    mobile   string
    // 接种日期
    vaccinateDate   string
    // 接种时间
    vaccinateTime   string
    // 接种的疫苗信息
    vaccineList   []VaccineInfo
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
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetAlipayUserId() string {
    return r.alipayUserId
}
// IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetIsvUserId(isvUserId string) error {
    r.isvUserId = isvUserId
    r.Set("isv_user_id", isvUserId)
    return nil
}

// IsvUserId Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetIsvUserId() string {
    return r.isvUserId
}
// OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetOrderId() string {
    return r.orderId
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetName() string {
    return r.name
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetMobile() string {
    return r.mobile
}
// VaccinateDate Setter
// 接种日期
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetVaccinateDate(vaccinateDate string) error {
    r.vaccinateDate = vaccinateDate
    r.Set("vaccinate_date", vaccinateDate)
    return nil
}

// VaccinateDate Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetVaccinateDate() string {
    return r.vaccinateDate
}
// VaccinateTime Setter
// 接种时间
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetVaccinateTime(vaccinateTime string) error {
    r.vaccinateTime = vaccinateTime
    r.Set("vaccinate_time", vaccinateTime)
    return nil
}

// VaccinateTime Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetVaccinateTime() string {
    return r.vaccinateTime
}
// VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinVaccinateCompleteRequest) SetVaccineList(vaccineList []VaccineInfo) error {
    r.vaccineList = vaccineList
    r.Set("vaccine_list", vaccineList)
    return nil
}

// VaccineList Getter
func (r AlibabaHealthVaccinVaccinateCompleteRequest) GetVaccineList() []VaccineInfo {
    return r.vaccineList
}

package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单是否由裹裹发送消息 API请求
cainiao.endpoint.locker.top.order.noticesend.query

合作公司查询消息发送的接口，判断是否裹裹发送消息
*/
type CainiaoEndpointLockerTopOrderNoticesendQueryRequest struct {
    model.Params
    // 站点id
    stationId   string
    // 收件人手机号
    getterPhone   string
    // 运单号
    mailNo   string
}

// 初始化CainiaoEndpointLockerTopOrderNoticesendQueryRequest对象
func NewCainiaoEndpointLockerTopOrderNoticesendQueryRequest() *CainiaoEndpointLockerTopOrderNoticesendQueryRequest{
    return &CainiaoEndpointLockerTopOrderNoticesendQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.noticesend.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StationId Setter
// 站点id
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) SetStationId(stationId string) error {
    r.stationId = stationId
    r.Set("station_id", stationId)
    return nil
}

// StationId Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetStationId() string {
    return r.stationId
}
// GetterPhone Setter
// 收件人手机号
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) SetGetterPhone(getterPhone string) error {
    r.getterPhone = getterPhone
    r.Set("getter_phone", getterPhone)
    return nil
}

// GetterPhone Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetGetterPhone() string {
    return r.getterPhone
}
// MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryRequest) GetMailNo() string {
    return r.mailNo
}

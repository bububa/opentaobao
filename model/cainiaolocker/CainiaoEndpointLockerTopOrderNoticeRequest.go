package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手动触发发短信 API请求
cainiao.endpoint.locker.top.order.notice

合作公司对订单手动触发短信，有次数限制
*/
type CainiaoEndpointLockerTopOrderNoticeRequest struct {
    model.Params
    // 合作公司唯一订单编号
    orderCode   string
    // 站点ID
    stationId   string
    // 运单号
    mailNo   string
    // 场景编号：0：重发短信，1：催取短信
    sceneCode   int64
}

// 初始化CainiaoEndpointLockerTopOrderNoticeRequest对象
func NewCainiaoEndpointLockerTopOrderNoticeRequest() *CainiaoEndpointLockerTopOrderNoticeRequest{
    return &CainiaoEndpointLockerTopOrderNoticeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.notice"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 合作公司唯一订单编号
func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetOrderCode() string {
    return r.orderCode
}
// StationId Setter
// 站点ID
func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetStationId(stationId string) error {
    r.stationId = stationId
    r.Set("station_id", stationId)
    return nil
}

// StationId Getter
func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetStationId() string {
    return r.stationId
}
// MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetMailNo() string {
    return r.mailNo
}
// SceneCode Setter
// 场景编号：0：重发短信，1：催取短信
func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetSceneCode(sceneCode int64) error {
    r.sceneCode = sceneCode
    r.Set("scene_code", sceneCode)
    return nil
}

// SceneCode Getter
func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetSceneCode() int64 {
    return r.sceneCode
}

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
type CainiaoEndpointLockerTopOrderNoticeAPIRequest struct {
    model.Params
    // 合作公司唯一订单编号
    _orderCode   string
    // 站点ID
    _stationId   string
    // 运单号
    _mailNo   string
    // 场景编号：0：重发短信，1：催取短信
    _sceneCode   int64
}

// 初始化CainiaoEndpointLockerTopOrderNoticeAPIRequest对象
func NewCainiaoEndpointLockerTopOrderNoticeRequest() *CainiaoEndpointLockerTopOrderNoticeAPIRequest{
    return &CainiaoEndpointLockerTopOrderNoticeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.notice"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 合作公司唯一订单编号
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// StationId Setter
// 站点ID
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetStationId(_stationId string) error {
    r._stationId = _stationId
    r.Set("station_id", _stationId)
    return nil
}

// StationId Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetStationId() string {
    return r._stationId
}
// MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetMailNo(_mailNo string) error {
    r._mailNo = _mailNo
    r.Set("mail_no", _mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetMailNo() string {
    return r._mailNo
}
// SceneCode Setter
// 场景编号：0：重发短信，1：催取短信
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetSceneCode(_sceneCode int64) error {
    r._sceneCode = _sceneCode
    r.Set("scene_code", _sceneCode)
    return nil
}

// SceneCode Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetSceneCode() int64 {
    return r._sceneCode
}

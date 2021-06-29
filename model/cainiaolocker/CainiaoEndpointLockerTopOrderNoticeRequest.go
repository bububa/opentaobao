package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手动触发发短信 APIRequest
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

func NewCainiaoEndpointLockerTopOrderNoticeRequest() *CainiaoEndpointLockerTopOrderNoticeRequest{
    return &CainiaoEndpointLockerTopOrderNoticeRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.notice"
}

func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetStationId(stationId string) error {
    r.stationId = stationId
    r.Set("station_id", stationId)
    return nil
}

func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetStationId() string {
    return r.stationId
}

func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetMailNo() string {
    return r.mailNo
}

func (r *CainiaoEndpointLockerTopOrderNoticeRequest) SetSceneCode(sceneCode int64) error {
    r.sceneCode = sceneCode
    r.Set("scene_code", sceneCode)
    return nil
}

func (r CainiaoEndpointLockerTopOrderNoticeRequest) GetSceneCode() int64 {
    return r.sceneCode
}


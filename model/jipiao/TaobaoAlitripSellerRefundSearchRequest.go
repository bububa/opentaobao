package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】退票申请单列表 APIRequest
taobao.alitrip.seller.refund.search

查询退票申请单列表
*/
type TaobaoAlitripSellerRefundSearchRequest struct {
    model.Params

    // 结束时间
    endTime   string 

    // 开始时间
    startTime   string 

    // 申请单状态（如果为空查询所有状态，1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
    status   int64 

}

func NewTaobaoAlitripSellerRefundSearchRequest() *TaobaoAlitripSellerRefundSearchRequest{
    return &TaobaoAlitripSellerRefundSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerRefundSearchRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.search"
}

func (r TaobaoAlitripSellerRefundSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerRefundSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoAlitripSellerRefundSearchRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoAlitripSellerRefundSearchRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoAlitripSellerRefundSearchRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoAlitripSellerRefundSearchRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoAlitripSellerRefundSearchRequest) GetStatus() int64 {
    return r.status
}


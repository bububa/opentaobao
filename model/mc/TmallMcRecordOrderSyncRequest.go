package mc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息同步 APIRequest
tmall.mc.record.order.sync

订单信息同步(零售云接口)
*/
type TmallMcRecordOrderSyncRequest struct {
    model.Params

    // 设备编码
    deviceCode   string 

    // 原价
    originPrice   int64 

    // 实付价
    payPrice   int64 

    // 用户openId
    openId   string 

    // 核销结果
    result   string 

    // 云码版本号
    version   string 

}

func NewTmallMcRecordOrderSyncRequest() *TmallMcRecordOrderSyncRequest{
    return &TmallMcRecordOrderSyncRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMcRecordOrderSyncRequest) GetApiMethodName() string {
    return "tmall.mc.record.order.sync"
}

func (r TmallMcRecordOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMcRecordOrderSyncRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TmallMcRecordOrderSyncRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TmallMcRecordOrderSyncRequest) SetOriginPrice(originPrice int64) error {
    r.originPrice = originPrice
    r.Set("origin_price", originPrice)
    return nil
}

func (r TmallMcRecordOrderSyncRequest) GetOriginPrice() int64 {
    return r.originPrice
}

func (r *TmallMcRecordOrderSyncRequest) SetPayPrice(payPrice int64) error {
    r.payPrice = payPrice
    r.Set("pay_price", payPrice)
    return nil
}

func (r TmallMcRecordOrderSyncRequest) GetPayPrice() int64 {
    return r.payPrice
}

func (r *TmallMcRecordOrderSyncRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TmallMcRecordOrderSyncRequest) GetOpenId() string {
    return r.openId
}

func (r *TmallMcRecordOrderSyncRequest) SetResult(result string) error {
    r.result = result
    r.Set("result", result)
    return nil
}

func (r TmallMcRecordOrderSyncRequest) GetResult() string {
    return r.result
}

func (r *TmallMcRecordOrderSyncRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r TmallMcRecordOrderSyncRequest) GetVersion() string {
    return r.version
}


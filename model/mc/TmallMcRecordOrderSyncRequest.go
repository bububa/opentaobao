package mc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息同步 API请求
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

// 初始化TmallMcRecordOrderSyncRequest对象
func NewTmallMcRecordOrderSyncRequest() *TmallMcRecordOrderSyncRequest{
    return &TmallMcRecordOrderSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMcRecordOrderSyncRequest) GetApiMethodName() string {
    return "tmall.mc.record.order.sync"
}

// IRequest interface 方法, 获取API参数
func (r TmallMcRecordOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备编码
func (r *TmallMcRecordOrderSyncRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallMcRecordOrderSyncRequest) GetDeviceCode() string {
    return r.deviceCode
}
// OriginPrice Setter
// 原价
func (r *TmallMcRecordOrderSyncRequest) SetOriginPrice(originPrice int64) error {
    r.originPrice = originPrice
    r.Set("origin_price", originPrice)
    return nil
}

// OriginPrice Getter
func (r TmallMcRecordOrderSyncRequest) GetOriginPrice() int64 {
    return r.originPrice
}
// PayPrice Setter
// 实付价
func (r *TmallMcRecordOrderSyncRequest) SetPayPrice(payPrice int64) error {
    r.payPrice = payPrice
    r.Set("pay_price", payPrice)
    return nil
}

// PayPrice Getter
func (r TmallMcRecordOrderSyncRequest) GetPayPrice() int64 {
    return r.payPrice
}
// OpenId Setter
// 用户openId
func (r *TmallMcRecordOrderSyncRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

// OpenId Getter
func (r TmallMcRecordOrderSyncRequest) GetOpenId() string {
    return r.openId
}
// Result Setter
// 核销结果
func (r *TmallMcRecordOrderSyncRequest) SetResult(result string) error {
    r.result = result
    r.Set("result", result)
    return nil
}

// Result Getter
func (r TmallMcRecordOrderSyncRequest) GetResult() string {
    return r.result
}
// Version Setter
// 云码版本号
func (r *TmallMcRecordOrderSyncRequest) SetVersion(version string) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r TmallMcRecordOrderSyncRequest) GetVersion() string {
    return r.version
}

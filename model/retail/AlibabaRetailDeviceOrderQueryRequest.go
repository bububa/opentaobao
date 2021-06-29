package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机订单查询 API请求
alibaba.retail.device.order.query

贩卖机订单查询
*/
type AlibabaRetailDeviceOrderQueryRequest struct {
    model.Params
    // 阿里设备物理ID
    deviceSnList   []string
    // 外部设备编码
    deviceUuid   string
    // 阿里设备编码
    deviceCode   string
    // -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
    status   int64
    // CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
    payType   string
    // 分页大小
    pageSize   int64
    // 页码
    pageNum   int64
    // 查询订单开始时间
    starts   string
    // 查询订单结束时间
    ends   string
}

// 初始化AlibabaRetailDeviceOrderQueryRequest对象
func NewAlibabaRetailDeviceOrderQueryRequest() *AlibabaRetailDeviceOrderQueryRequest{
    return &AlibabaRetailDeviceOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.retail.device.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceSnList Setter
// 阿里设备物理ID
func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceSnList(deviceSnList []string) error {
    r.deviceSnList = deviceSnList
    r.Set("device_sn_list", deviceSnList)
    return nil
}

// DeviceSnList Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceSnList() []string {
    return r.deviceSnList
}
// DeviceUuid Setter
// 外部设备编码
func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceUuid(deviceUuid string) error {
    r.deviceUuid = deviceUuid
    r.Set("device_uuid", deviceUuid)
    return nil
}

// DeviceUuid Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceUuid() string {
    return r.deviceUuid
}
// DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceCode() string {
    return r.deviceCode
}
// Status Setter
// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
func (r *AlibabaRetailDeviceOrderQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetStatus() int64 {
    return r.status
}
// PayType Setter
// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
func (r *AlibabaRetailDeviceOrderQueryRequest) SetPayType(payType string) error {
    r.payType = payType
    r.Set("pay_type", payType)
    return nil
}

// PayType Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetPayType() string {
    return r.payType
}
// PageSize Setter
// 分页大小
func (r *AlibabaRetailDeviceOrderQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNum Setter
// 页码
func (r *AlibabaRetailDeviceOrderQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetPageNum() int64 {
    return r.pageNum
}
// Starts Setter
// 查询订单开始时间
func (r *AlibabaRetailDeviceOrderQueryRequest) SetStarts(starts string) error {
    r.starts = starts
    r.Set("starts", starts)
    return nil
}

// Starts Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetStarts() string {
    return r.starts
}
// Ends Setter
// 查询订单结束时间
func (r *AlibabaRetailDeviceOrderQueryRequest) SetEnds(ends string) error {
    r.ends = ends
    r.Set("ends", ends)
    return nil
}

// Ends Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetEnds() string {
    return r.ends
}

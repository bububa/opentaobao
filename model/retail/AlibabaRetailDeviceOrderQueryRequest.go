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
    _deviceSnList   []string
    // 外部设备编码
    _deviceUuid   string
    // 阿里设备编码
    _deviceCode   string
    // -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
    _status   int64
    // CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
    _payType   string
    // 分页大小
    _pageSize   int64
    // 页码
    _pageNum   int64
    // 查询订单开始时间
    _starts   string
    // 查询订单结束时间
    _ends   string
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
func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceSnList(_deviceSnList []string) error {
    r._deviceSnList = _deviceSnList
    r.Set("device_sn_list", _deviceSnList)
    return nil
}

// DeviceSnList Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceSnList() []string {
    return r._deviceSnList
}
// DeviceUuid Setter
// 外部设备编码
func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceUuid(_deviceUuid string) error {
    r._deviceUuid = _deviceUuid
    r.Set("device_uuid", _deviceUuid)
    return nil
}

// DeviceUuid Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceUuid() string {
    return r._deviceUuid
}
// DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceCode() string {
    return r._deviceCode
}
// Status Setter
// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
func (r *AlibabaRetailDeviceOrderQueryRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetStatus() int64 {
    return r._status
}
// PayType Setter
// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
func (r *AlibabaRetailDeviceOrderQueryRequest) SetPayType(_payType string) error {
    r._payType = _payType
    r.Set("pay_type", _payType)
    return nil
}

// PayType Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetPayType() string {
    return r._payType
}
// PageSize Setter
// 分页大小
func (r *AlibabaRetailDeviceOrderQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNum Setter
// 页码
func (r *AlibabaRetailDeviceOrderQueryRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetPageNum() int64 {
    return r._pageNum
}
// Starts Setter
// 查询订单开始时间
func (r *AlibabaRetailDeviceOrderQueryRequest) SetStarts(_starts string) error {
    r._starts = _starts
    r.Set("starts", _starts)
    return nil
}

// Starts Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetStarts() string {
    return r._starts
}
// Ends Setter
// 查询订单结束时间
func (r *AlibabaRetailDeviceOrderQueryRequest) SetEnds(_ends string) error {
    r._ends = _ends
    r.Set("ends", _ends)
    return nil
}

// Ends Getter
func (r AlibabaRetailDeviceOrderQueryRequest) GetEnds() string {
    return r._ends
}

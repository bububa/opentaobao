package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机订单查询 APIRequest
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

func NewAlibabaRetailDeviceOrderQueryRequest() *AlibabaRetailDeviceOrderQueryRequest{
    return &AlibabaRetailDeviceOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.retail.device.order.query"
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceSnList(deviceSnList []string) error {
    r.deviceSnList = deviceSnList
    r.Set("device_sn_list", deviceSnList)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceSnList() []string {
    return r.deviceSnList
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceUuid(deviceUuid string) error {
    r.deviceUuid = deviceUuid
    r.Set("device_uuid", deviceUuid)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceUuid() string {
    return r.deviceUuid
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetPayType(payType string) error {
    r.payType = payType
    r.Set("pay_type", payType)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetPayType() string {
    return r.payType
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetStarts(starts string) error {
    r.starts = starts
    r.Set("starts", starts)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetStarts() string {
    return r.starts
}

func (r *AlibabaRetailDeviceOrderQueryRequest) SetEnds(ends string) error {
    r.ends = ends
    r.Set("ends", ends)
    return nil
}

func (r AlibabaRetailDeviceOrderQueryRequest) GetEnds() string {
    return r.ends
}


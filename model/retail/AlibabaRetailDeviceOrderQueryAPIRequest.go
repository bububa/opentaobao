package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceOrderQueryAPIRequest 贩卖机订单查询 API请求
// alibaba.retail.device.order.query
//
// 贩卖机订单查询
type AlibabaRetailDeviceOrderQueryAPIRequest struct {
	model.Params
	// 阿里设备物理ID
	_deviceSnList []string
	// 外部设备编码
	_deviceUuid string
	// 阿里设备编码
	_deviceCode string
	// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
	_status int64
	// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
	_payType string
	// 分页大小
	_pageSize int64
	// 页码
	_pageNum int64
	// 查询订单开始时间
	_starts string
	// 查询订单结束时间
	_ends string
}

// NewAlibabaRetailDeviceOrderQueryRequest 初始化AlibabaRetailDeviceOrderQueryAPIRequest对象
func NewAlibabaRetailDeviceOrderQueryRequest() *AlibabaRetailDeviceOrderQueryAPIRequest {
	return &AlibabaRetailDeviceOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceSnList Setter
// 阿里设备物理ID
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetDeviceSnList(_deviceSnList []string) error {
	r._deviceSnList = _deviceSnList
	r.Set("device_sn_list", _deviceSnList)
	return nil
}

// Get DeviceSnList Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetDeviceSnList() []string {
	return r._deviceSnList
}

// Set is DeviceUuid Setter
// 外部设备编码
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetDeviceUuid(_deviceUuid string) error {
	r._deviceUuid = _deviceUuid
	r.Set("device_uuid", _deviceUuid)
	return nil
}

// Get DeviceUuid Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetDeviceUuid() string {
	return r._deviceUuid
}

// Set is DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// Get DeviceCode Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// Set is Status Setter
// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is PayType Setter
// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetPayType(_payType string) error {
	r._payType = _payType
	r.Set("pay_type", _payType)
	return nil
}

// Get PayType Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetPayType() string {
	return r._payType
}

// Set is PageSize Setter
// 分页大小
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNum Setter
// 页码
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// Get PageNum Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// Set is Starts Setter
// 查询订单开始时间
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetStarts(_starts string) error {
	r._starts = _starts
	r.Set("starts", _starts)
	return nil
}

// Get Starts Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetStarts() string {
	return r._starts
}

// Set is Ends Setter
// 查询订单结束时间
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetEnds(_ends string) error {
	r._ends = _ends
	r.Set("ends", _ends)
	return nil
}

// Get Ends Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetEnds() string {
	return r._ends
}

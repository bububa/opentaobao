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
	// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
	_payType string
	// 查询订单开始时间
	_starts string
	// 查询订单结束时间
	_ends string
	// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
	_status int64
	// 分页大小
	_pageSize int64
	// 页码
	_pageNum int64
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

// SetDeviceSnList is DeviceSnList Setter
// 阿里设备物理ID
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetDeviceSnList(_deviceSnList []string) error {
	r._deviceSnList = _deviceSnList
	r.Set("device_sn_list", _deviceSnList)
	return nil
}

// GetDeviceSnList DeviceSnList Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetDeviceSnList() []string {
	return r._deviceSnList
}

// SetDeviceUuid is DeviceUuid Setter
// 外部设备编码
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetDeviceUuid(_deviceUuid string) error {
	r._deviceUuid = _deviceUuid
	r.Set("device_uuid", _deviceUuid)
	return nil
}

// GetDeviceUuid DeviceUuid Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetDeviceUuid() string {
	return r._deviceUuid
}

// SetDeviceCode is DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetPayType is PayType Setter
// CASH 现金，ALIPAY_FACE_PAY_QR 支付宝，VENDING_PRIZE 抽奖，FACE_PAY 人脸
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetPayType(_payType string) error {
	r._payType = _payType
	r.Set("pay_type", _payType)
	return nil
}

// GetPayType PayType Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetPayType() string {
	return r._payType
}

// SetStarts is Starts Setter
// 查询订单开始时间
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetStarts(_starts string) error {
	r._starts = _starts
	r.Set("starts", _starts)
	return nil
}

// GetStarts Starts Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetStarts() string {
	return r._starts
}

// SetEnds is Ends Setter
// 查询订单结束时间
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetEnds(_ends string) error {
	r._ends = _ends
	r.Set("ends", _ends)
	return nil
}

// GetEnds Ends Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetEnds() string {
	return r._ends
}

// SetStatus is Status Setter
// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码
func (r *AlibabaRetailDeviceOrderQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaRetailDeviceOrderQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticssellerordersgetAPIRequest 商家配送核销订单列表查询 API请求
// alibaba.ascp.logistics.seller.orders.get
//
// 商家配送核销订单列表查询
type AlibabaascplogisticssellerordersgetAPIRequest struct {
	model.Params
	// 核销日期
	_writeOffDate string
	// 收货码
	_receiveCode string
	// 淘系交易id
	_tid string
	// 1代表未核销，2代表已核销
	_writeOffStatus string
	// 分页索引
	_pageIndex int64
	// 分页大小
	_pageSize int64
}

// NewAlibabaascplogisticssellerordersgetRequest 初始化AlibabaascplogisticssellerordersgetAPIRequest对象
func NewAlibabaascplogisticssellerordersgetRequest() *AlibabaascplogisticssellerordersgetAPIRequest {
	return &AlibabaascplogisticssellerordersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.seller.orders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWriteOffDate is WriteOffDate Setter
// 核销日期
func (r *AlibabaascplogisticssellerordersgetAPIRequest) SetWriteOffDate(_writeOffDate string) error {
	r._writeOffDate = _writeOffDate
	r.Set("write_off_date", _writeOffDate)
	return nil
}

// GetWriteOffDate WriteOffDate Getter
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetWriteOffDate() string {
	return r._writeOffDate
}

// SetReceiveCode is ReceiveCode Setter
// 收货码
func (r *AlibabaascplogisticssellerordersgetAPIRequest) SetReceiveCode(_receiveCode string) error {
	r._receiveCode = _receiveCode
	r.Set("receive_code", _receiveCode)
	return nil
}

// GetReceiveCode ReceiveCode Getter
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetReceiveCode() string {
	return r._receiveCode
}

// SetTid is Tid Setter
// 淘系交易id
func (r *AlibabaascplogisticssellerordersgetAPIRequest) SetTid(_tid string) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetTid() string {
	return r._tid
}

// SetWriteOffStatus is WriteOffStatus Setter
// 1代表未核销，2代表已核销
func (r *AlibabaascplogisticssellerordersgetAPIRequest) SetWriteOffStatus(_writeOffStatus string) error {
	r._writeOffStatus = _writeOffStatus
	r.Set("write_off_status", _writeOffStatus)
	return nil
}

// GetWriteOffStatus WriteOffStatus Getter
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetWriteOffStatus() string {
	return r._writeOffStatus
}

// SetPageIndex is PageIndex Setter
// 分页索引
func (r *AlibabaascplogisticssellerordersgetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *AlibabaascplogisticssellerordersgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaascplogisticssellerordersgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

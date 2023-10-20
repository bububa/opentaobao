package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdscommissionorderlistbyindexAPIRequest 联盟订单分页查询 API请求
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
type AliexpressdscommissionorderlistbyindexAPIRequest struct {
	model.Params
	// Query index start value: if not passed, You can only check the first page
	_startQueryIndexId string
	// Start time, PST time
	_startTime string
	// End time, PST time
	_endTime string
	// Order status: Payment Completed(Buyer paid successfully), Buyer Confirmed Receipt(This status only change when :Buyer confirms receipt and settlement task begins which is manually executed by our operation team)
	_status string
	// record count of each page, 1 - 50
	_pageSize int64
	// page number
	_pageNo int64
}

// NewAliexpressdscommissionorderlistbyindexRequest 初始化AliexpressdscommissionorderlistbyindexAPIRequest对象
func NewAliexpressdscommissionorderlistbyindexRequest() *AliexpressdscommissionorderlistbyindexAPIRequest {
	return &AliexpressdscommissionorderlistbyindexAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.commissionorder.listbyindex"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartQueryIndexId is StartQueryIndexId Setter
// Query index start value: if not passed, You can only check the first page
func (r *AliexpressdscommissionorderlistbyindexAPIRequest) SetStartQueryIndexId(_startQueryIndexId string) error {
	r._startQueryIndexId = _startQueryIndexId
	r.Set("start_query_index_id", _startQueryIndexId)
	return nil
}

// GetStartQueryIndexId StartQueryIndexId Getter
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetStartQueryIndexId() string {
	return r._startQueryIndexId
}

// SetStartTime is StartTime Setter
// Start time, PST time
func (r *AliexpressdscommissionorderlistbyindexAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// End time, PST time
func (r *AliexpressdscommissionorderlistbyindexAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// Order status: Payment Completed(Buyer paid successfully), Buyer Confirmed Receipt(This status only change when :Buyer confirms receipt and settlement task begins which is manually executed by our operation team)
func (r *AliexpressdscommissionorderlistbyindexAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetStatus() string {
	return r._status
}

// SetPageSize is PageSize Setter
// record count of each page, 1 - 50
func (r *AliexpressdscommissionorderlistbyindexAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// page number
func (r *AliexpressdscommissionorderlistbyindexAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressdscommissionorderlistbyindexAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsCommissionorderListbyindexAPIRequest 联盟订单分页查询 API请求
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
type AliexpressDsCommissionorderListbyindexAPIRequest struct {
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

// NewAliexpressDsCommissionorderListbyindexRequest 初始化AliexpressDsCommissionorderListbyindexAPIRequest对象
func NewAliexpressDsCommissionorderListbyindexRequest() *AliexpressDsCommissionorderListbyindexAPIRequest {
	return &AliexpressDsCommissionorderListbyindexAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) Reset() {
	r._startQueryIndexId = ""
	r._startTime = ""
	r._endTime = ""
	r._status = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.commissionorder.listbyindex"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartQueryIndexId is StartQueryIndexId Setter
// Query index start value: if not passed, You can only check the first page
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) SetStartQueryIndexId(_startQueryIndexId string) error {
	r._startQueryIndexId = _startQueryIndexId
	r.Set("start_query_index_id", _startQueryIndexId)
	return nil
}

// GetStartQueryIndexId StartQueryIndexId Getter
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetStartQueryIndexId() string {
	return r._startQueryIndexId
}

// SetStartTime is StartTime Setter
// Start time, PST time
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// End time, PST time
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// Order status: Payment Completed(Buyer paid successfully), Buyer Confirmed Receipt(This status only change when :Buyer confirms receipt and settlement task begins which is manually executed by our operation team)
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetStatus() string {
	return r._status
}

// SetPageSize is PageSize Setter
// record count of each page, 1 - 50
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// page number
func (r *AliexpressDsCommissionorderListbyindexAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressDsCommissionorderListbyindexAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolAliexpressDsCommissionorderListbyindexAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressDsCommissionorderListbyindexRequest()
	},
}

// GetAliexpressDsCommissionorderListbyindexRequest 从 sync.Pool 获取 AliexpressDsCommissionorderListbyindexAPIRequest
func GetAliexpressDsCommissionorderListbyindexAPIRequest() *AliexpressDsCommissionorderListbyindexAPIRequest {
	return poolAliexpressDsCommissionorderListbyindexAPIRequest.Get().(*AliexpressDsCommissionorderListbyindexAPIRequest)
}

// ReleaseAliexpressDsCommissionorderListbyindexAPIRequest 将 AliexpressDsCommissionorderListbyindexAPIRequest 放入 sync.Pool
func ReleaseAliexpressDsCommissionorderListbyindexAPIRequest(v *AliexpressDsCommissionorderListbyindexAPIRequest) {
	v.Reset()
	poolAliexpressDsCommissionorderListbyindexAPIRequest.Put(v)
}

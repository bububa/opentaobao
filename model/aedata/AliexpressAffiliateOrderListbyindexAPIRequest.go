package aedata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateOrderListbyindexAPIRequest AE联盟推广者订单查询接口-按游标索引查询 API请求
// aliexpress.affiliate.order.listbyindex
//
// AE联盟推广者订单按游标查询接口
type AliexpressAffiliateOrderListbyindexAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 查询索引开始值：若不传，则只能查第一页
	_startQueryIndexId string
	// 结束时间
	_endTime string
	// 订单状态:Payment Completed,Buyer Confirmed Receipt
	_status string
	// 返回的字段信息
	_fields string
	// 安全签名
	_appSignature string
	// 每页记录数
	_pageSize int64
}

// NewAliexpressAffiliateOrderListbyindexRequest 初始化AliexpressAffiliateOrderListbyindexAPIRequest对象
func NewAliexpressAffiliateOrderListbyindexRequest() *AliexpressAffiliateOrderListbyindexAPIRequest {
	return &AliexpressAffiliateOrderListbyindexAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) Reset() {
	r._startTime = ""
	r._startQueryIndexId = ""
	r._endTime = ""
	r._status = ""
	r._fields = ""
	r._appSignature = ""
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.order.listbyindex"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetStartQueryIndexId is StartQueryIndexId Setter
// 查询索引开始值：若不传，则只能查第一页
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetStartQueryIndexId(_startQueryIndexId string) error {
	r._startQueryIndexId = _startQueryIndexId
	r.Set("start_query_index_id", _startQueryIndexId)
	return nil
}

// GetStartQueryIndexId StartQueryIndexId Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetStartQueryIndexId() string {
	return r._startQueryIndexId
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// 订单状态:Payment Completed,Buyer Confirmed Receipt
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetStatus() string {
	return r._status
}

// SetFields is Fields Setter
// 返回的字段信息
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetFields() string {
	return r._fields
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateOrderListbyindexAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressAffiliateOrderListbyindexAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAliexpressAffiliateOrderListbyindexAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAffiliateOrderListbyindexRequest()
	},
}

// GetAliexpressAffiliateOrderListbyindexRequest 从 sync.Pool 获取 AliexpressAffiliateOrderListbyindexAPIRequest
func GetAliexpressAffiliateOrderListbyindexAPIRequest() *AliexpressAffiliateOrderListbyindexAPIRequest {
	return poolAliexpressAffiliateOrderListbyindexAPIRequest.Get().(*AliexpressAffiliateOrderListbyindexAPIRequest)
}

// ReleaseAliexpressAffiliateOrderListbyindexAPIRequest 将 AliexpressAffiliateOrderListbyindexAPIRequest 放入 sync.Pool
func ReleaseAliexpressAffiliateOrderListbyindexAPIRequest(v *AliexpressAffiliateOrderListbyindexAPIRequest) {
	v.Reset()
	poolAliexpressAffiliateOrderListbyindexAPIRequest.Put(v)
}

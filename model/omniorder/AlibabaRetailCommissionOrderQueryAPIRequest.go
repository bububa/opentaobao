package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionOrderQueryAPIRequest 分销订单查询 API请求
// alibaba.retail.commission.order.query
//
// 查询商家的分销订单
type AlibabaRetailCommissionOrderQueryAPIRequest struct {
	model.Params
	// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
	_endPayTime string
	// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
	_startPayTime string
	// 页码，默认第一页
	_pageNo int64
	// 页大小，默认每页十条
	_pageSize int64
}

// NewAlibabaRetailCommissionOrderQueryRequest 初始化AlibabaRetailCommissionOrderQueryAPIRequest对象
func NewAlibabaRetailCommissionOrderQueryRequest() *AlibabaRetailCommissionOrderQueryAPIRequest {
	return &AlibabaRetailCommissionOrderQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailCommissionOrderQueryAPIRequest) Reset() {
	r._endPayTime = ""
	r._startPayTime = ""
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndPayTime is EndPayTime Setter
// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaRetailCommissionOrderQueryAPIRequest) SetEndPayTime(_endPayTime string) error {
	r._endPayTime = _endPayTime
	r.Set("end_pay_time", _endPayTime)
	return nil
}

// GetEndPayTime EndPayTime Getter
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetEndPayTime() string {
	return r._endPayTime
}

// SetStartPayTime is StartPayTime Setter
// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
func (r *AlibabaRetailCommissionOrderQueryAPIRequest) SetStartPayTime(_startPayTime string) error {
	r._startPayTime = _startPayTime
	r.Set("start_pay_time", _startPayTime)
	return nil
}

// GetStartPayTime StartPayTime Getter
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetStartPayTime() string {
	return r._startPayTime
}

// SetPageNo is PageNo Setter
// 页码，默认第一页
func (r *AlibabaRetailCommissionOrderQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小，默认每页十条
func (r *AlibabaRetailCommissionOrderQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaRetailCommissionOrderQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaRetailCommissionOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailCommissionOrderQueryRequest()
	},
}

// GetAlibabaRetailCommissionOrderQueryRequest 从 sync.Pool 获取 AlibabaRetailCommissionOrderQueryAPIRequest
func GetAlibabaRetailCommissionOrderQueryAPIRequest() *AlibabaRetailCommissionOrderQueryAPIRequest {
	return poolAlibabaRetailCommissionOrderQueryAPIRequest.Get().(*AlibabaRetailCommissionOrderQueryAPIRequest)
}

// ReleaseAlibabaRetailCommissionOrderQueryAPIRequest 将 AlibabaRetailCommissionOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailCommissionOrderQueryAPIRequest(v *AlibabaRetailCommissionOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaRetailCommissionOrderQueryAPIRequest.Put(v)
}

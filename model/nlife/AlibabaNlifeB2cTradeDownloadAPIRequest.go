package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ctradedownloadAPIRequest b2c下载订单 API请求
// alibaba.nlife.b2c.trade.download
//
// 下载零售商在零售+平台创建的订单
type Alibabanlifeb2ctradedownloadAPIRequest struct {
	model.Params
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
	// 零售门店在零售+平台对应的ID
	_storeId string
	// 页码
	_pageNo int64
	// 分页大小
	_pageSize int64
}

// NewAlibabanlifeb2ctradedownloadRequest 初始化Alibabanlifeb2ctradedownloadAPIRequest对象
func NewAlibabanlifeb2ctradedownloadRequest() *Alibabanlifeb2ctradedownloadAPIRequest {
	return &Alibabanlifeb2ctradedownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *Alibabanlifeb2ctradedownloadAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *Alibabanlifeb2ctradedownloadAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStoreId is StoreId Setter
// 零售门店在零售+平台对应的ID
func (r *Alibabanlifeb2ctradedownloadAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetPageNo is PageNo Setter
// 页码
func (r *Alibabanlifeb2ctradedownloadAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *Alibabanlifeb2ctradedownloadAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r Alibabanlifeb2ctradedownloadAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

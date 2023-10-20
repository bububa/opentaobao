package aedata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliateorderlistAPIRequest AE推广者订单批量获取接口 API请求
// aliexpress.affiliate.order.list
//
// AE联盟推广者订单分页查询接口
type AliexpressaffiliateorderlistAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 订单状态:Payment Completed,Buyer Confirmed Receipt
	_status string
	// 站点信息：global、ru_site、es_site、it_site
	_localeSite string
	// 返回的字段信息
	_fields string
	// 安全签名
	_appSignature string
	// 页数
	_pageNo int64
	// 每页记录数
	_pageSize int64
}

// NewAliexpressaffiliateorderlistRequest 初始化AliexpressaffiliateorderlistAPIRequest对象
func NewAliexpressaffiliateorderlistRequest() *AliexpressaffiliateorderlistAPIRequest {
	return &AliexpressaffiliateorderlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliateorderlistAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.order.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliateorderlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliateorderlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *AliexpressaffiliateorderlistAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *AliexpressaffiliateorderlistAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// 订单状态:Payment Completed,Buyer Confirmed Receipt
func (r *AliexpressaffiliateorderlistAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetStatus() string {
	return r._status
}

// SetLocaleSite is LocaleSite Setter
// 站点信息：global、ru_site、es_site、it_site
func (r *AliexpressaffiliateorderlistAPIRequest) SetLocaleSite(_localeSite string) error {
	r._localeSite = _localeSite
	r.Set("locale_site", _localeSite)
	return nil
}

// GetLocaleSite LocaleSite Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetLocaleSite() string {
	return r._localeSite
}

// SetFields is Fields Setter
// 返回的字段信息
func (r *AliexpressaffiliateorderlistAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetFields() string {
	return r._fields
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressaffiliateorderlistAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetPageNo is PageNo Setter
// 页数
func (r *AliexpressaffiliateorderlistAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *AliexpressaffiliateorderlistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressaffiliateorderlistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

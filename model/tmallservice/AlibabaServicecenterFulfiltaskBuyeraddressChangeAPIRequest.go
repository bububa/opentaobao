package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest 修改消费者服务地址 API请求
// alibaba.servicecenter.fulfiltask.buyeraddress.change
//
// 当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子
type AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest struct {
	model.Params
	// 详细地址
	_addressDetail string
	// 核销单id
	_fulfilTaskId int64
	// 地址编码
	_location int64
}

// NewAlibabaservicecenterfulfiltaskbuyeraddresschangeRequest 初始化AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest对象
func NewAlibabaservicecenterfulfiltaskbuyeraddresschangeRequest() *AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest {
	return &AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.fulfiltask.buyeraddress.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddressDetail is AddressDetail Setter
// 详细地址
func (r *AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) SetAddressDetail(_addressDetail string) error {
	r._addressDetail = _addressDetail
	r.Set("address_detail", _addressDetail)
	return nil
}

// GetAddressDetail AddressDetail Getter
func (r AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) GetAddressDetail() string {
	return r._addressDetail
}

// SetFulfilTaskId is FulfilTaskId Setter
// 核销单id
func (r *AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}

// SetLocation is Location Setter
// 地址编码
func (r *AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) SetLocation(_location int64) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// GetLocation Location Getter
func (r AlibabaservicecenterfulfiltaskbuyeraddresschangeAPIRequest) GetLocation() int64 {
	return r._location
}

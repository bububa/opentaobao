package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest 修改消费者服务地址 API请求
// alibaba.servicecenter.fulfiltask.buyeraddress.change
//
// 当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子
type AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest struct {
	model.Params
	// 核销单id
	_fulfilTaskId int64
	// 详细地址
	_addressDetail string
	// 地址编码
	_location int64
}

// NewAlibabaServicecenterFulfiltaskBuyeraddressChangeRequest 初始化AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest对象
func NewAlibabaServicecenterFulfiltaskBuyeraddressChangeRequest() *AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest {
	return &AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.fulfiltask.buyeraddress.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFulfilTaskId is FulfilTaskId Setter
// 核销单id
func (r *AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}

// SetAddressDetail is AddressDetail Setter
// 详细地址
func (r *AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) SetAddressDetail(_addressDetail string) error {
	r._addressDetail = _addressDetail
	r.Set("address_detail", _addressDetail)
	return nil
}

// GetAddressDetail AddressDetail Getter
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) GetAddressDetail() string {
	return r._addressDetail
}

// SetLocation is Location Setter
// 地址编码
func (r *AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) SetLocation(_location int64) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// GetLocation Location Getter
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest) GetLocation() int64 {
	return r._location
}

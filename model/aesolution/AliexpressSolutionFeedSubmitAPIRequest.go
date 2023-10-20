package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionfeedsubmitAPIRequest aliexpress.solution.feed.submit API请求
// aliexpress.solution.feed.submit
//
// API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
type AliexpresssolutionfeedsubmitAPIRequest struct {
	model.Params
	// item list, maximum size: 2000.
	_itemList []SingleItemRequestDto
	// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
	_operationType string
	// More information of the request.
	_developerFeatures string
}

// NewAliexpresssolutionfeedsubmitRequest 初始化AliexpresssolutionfeedsubmitAPIRequest对象
func NewAliexpresssolutionfeedsubmitRequest() *AliexpresssolutionfeedsubmitAPIRequest {
	return &AliexpresssolutionfeedsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionfeedsubmitAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionfeedsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionfeedsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// item list, maximum size: 2000.
func (r *AliexpresssolutionfeedsubmitAPIRequest) SetItemList(_itemList []SingleItemRequestDto) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r AliexpresssolutionfeedsubmitAPIRequest) GetItemList() []SingleItemRequestDto {
	return r._itemList
}

// SetOperationType is OperationType Setter
// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
func (r *AliexpresssolutionfeedsubmitAPIRequest) SetOperationType(_operationType string) error {
	r._operationType = _operationType
	r.Set("operation_type", _operationType)
	return nil
}

// GetOperationType OperationType Getter
func (r AliexpresssolutionfeedsubmitAPIRequest) GetOperationType() string {
	return r._operationType
}

// SetDeveloperFeatures is DeveloperFeatures Setter
// More information of the request.
func (r *AliexpresssolutionfeedsubmitAPIRequest) SetDeveloperFeatures(_developerFeatures string) error {
	r._developerFeatures = _developerFeatures
	r.Set("developer_features", _developerFeatures)
	return nil
}

// GetDeveloperFeatures DeveloperFeatures Getter
func (r AliexpresssolutionfeedsubmitAPIRequest) GetDeveloperFeatures() string {
	return r._developerFeatures
}

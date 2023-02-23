package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedSubmitAPIRequest aliexpress.solution.feed.submit API请求
// aliexpress.solution.feed.submit
//
// API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
type AliexpressSolutionFeedSubmitAPIRequest struct {
	model.Params
	// item list, maximum size: 2000.
	_itemList []SingleItemRequestDto
	// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
	_operationType string
	// More information of the request.
	_developerFeatures string
}

// NewAliexpressSolutionFeedSubmitRequest 初始化AliexpressSolutionFeedSubmitAPIRequest对象
func NewAliexpressSolutionFeedSubmitRequest() *AliexpressSolutionFeedSubmitAPIRequest {
	return &AliexpressSolutionFeedSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedSubmitAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionFeedSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemList is ItemList Setter
// item list, maximum size: 2000.
func (r *AliexpressSolutionFeedSubmitAPIRequest) SetItemList(_itemList []SingleItemRequestDto) error {
	r._itemList = _itemList
	r.Set("item_list", _itemList)
	return nil
}

// GetItemList ItemList Getter
func (r AliexpressSolutionFeedSubmitAPIRequest) GetItemList() []SingleItemRequestDto {
	return r._itemList
}

// SetOperationType is OperationType Setter
// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
func (r *AliexpressSolutionFeedSubmitAPIRequest) SetOperationType(_operationType string) error {
	r._operationType = _operationType
	r.Set("operation_type", _operationType)
	return nil
}

// GetOperationType OperationType Getter
func (r AliexpressSolutionFeedSubmitAPIRequest) GetOperationType() string {
	return r._operationType
}

// SetDeveloperFeatures is DeveloperFeatures Setter
// More information of the request.
func (r *AliexpressSolutionFeedSubmitAPIRequest) SetDeveloperFeatures(_developerFeatures string) error {
	r._developerFeatures = _developerFeatures
	r.Set("developer_features", _developerFeatures)
	return nil
}

// GetDeveloperFeatures DeveloperFeatures Getter
func (r AliexpressSolutionFeedSubmitAPIRequest) GetDeveloperFeatures() string {
	return r._developerFeatures
}

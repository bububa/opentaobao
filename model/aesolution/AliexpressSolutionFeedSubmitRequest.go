package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.submit API请求
aliexpress.solution.feed.submit

API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
*/
type AliexpressSolutionFeedSubmitAPIRequest struct {
    model.Params
    // Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
    _operationType   string
    // item list, maximum size: 2000.
    _itemList   []SingleItemRequestDTO
}

// 初始化AliexpressSolutionFeedSubmitAPIRequest对象
func NewAliexpressSolutionFeedSubmitRequest() *AliexpressSolutionFeedSubmitAPIRequest{
    return &AliexpressSolutionFeedSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedSubmitAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.submit"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperationType Setter
// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
func (r *AliexpressSolutionFeedSubmitAPIRequest) SetOperationType(_operationType string) error {
    r._operationType = _operationType
    r.Set("operation_type", _operationType)
    return nil
}

// OperationType Getter
func (r AliexpressSolutionFeedSubmitAPIRequest) GetOperationType() string {
    return r._operationType
}
// ItemList Setter
// item list, maximum size: 2000.
func (r *AliexpressSolutionFeedSubmitAPIRequest) SetItemList(_itemList []SingleItemRequestDTO) error {
    r._itemList = _itemList
    r.Set("item_list", _itemList)
    return nil
}

// ItemList Getter
func (r AliexpressSolutionFeedSubmitAPIRequest) GetItemList() []SingleItemRequestDTO {
    return r._itemList
}

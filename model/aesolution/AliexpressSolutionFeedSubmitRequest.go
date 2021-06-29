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
type AliexpressSolutionFeedSubmitRequest struct {
    model.Params
    // Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
    _operationType   string
    // item list, maximum size: 2000.
    _itemList   []SingleItemRequestDto
}

// 初始化AliexpressSolutionFeedSubmitRequest对象
func NewAliexpressSolutionFeedSubmitRequest() *AliexpressSolutionFeedSubmitRequest{
    return &AliexpressSolutionFeedSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedSubmitRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.submit"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperationType Setter
// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
func (r *AliexpressSolutionFeedSubmitRequest) SetOperationType(_operationType string) error {
    r._operationType = _operationType
    r.Set("operation_type", _operationType)
    return nil
}

// OperationType Getter
func (r AliexpressSolutionFeedSubmitRequest) GetOperationType() string {
    return r._operationType
}
// ItemList Setter
// item list, maximum size: 2000.
func (r *AliexpressSolutionFeedSubmitRequest) SetItemList(_itemList []SingleItemRequestDto) error {
    r._itemList = _itemList
    r.Set("item_list", _itemList)
    return nil
}

// ItemList Getter
func (r AliexpressSolutionFeedSubmitRequest) GetItemList() []SingleItemRequestDto {
    return r._itemList
}

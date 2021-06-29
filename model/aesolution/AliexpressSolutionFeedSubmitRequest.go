package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.submit APIRequest
aliexpress.solution.feed.submit

API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
*/
type AliexpressSolutionFeedSubmitRequest struct {
    model.Params

    // Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
    operationType   string 

    // item list, maximum size: 2000.
    itemList   []SingleItemRequestDto 

}

func NewAliexpressSolutionFeedSubmitRequest() *AliexpressSolutionFeedSubmitRequest{
    return &AliexpressSolutionFeedSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionFeedSubmitRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.submit"
}

func (r AliexpressSolutionFeedSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionFeedSubmitRequest) SetOperationType(operationType string) error {
    r.operationType = operationType
    r.Set("operation_type", operationType)
    return nil
}

func (r AliexpressSolutionFeedSubmitRequest) GetOperationType() string {
    return r.operationType
}

func (r *AliexpressSolutionFeedSubmitRequest) SetItemList(itemList []SingleItemRequestDto) error {
    r.itemList = itemList
    r.Set("item_list", itemList)
    return nil
}

func (r AliexpressSolutionFeedSubmitRequest) GetItemList() []SingleItemRequestDto {
    return r.itemList
}


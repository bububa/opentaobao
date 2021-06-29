package aetask

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE互动任务投放 APIRequest
aliexpress.interactive.task.delivery.query

将内部配置好的任务，如浏览商品，店铺投放给外部ISV
*/
type AliexpressInteractiveTaskDeliveryQueryRequest struct {
    model.Params

    // 返回结果
    requestDto   *QueryDeliveryRequestDto 

}

func NewAliexpressInteractiveTaskDeliveryQueryRequest() *AliexpressInteractiveTaskDeliveryQueryRequest{
    return &AliexpressInteractiveTaskDeliveryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressInteractiveTaskDeliveryQueryRequest) GetApiMethodName() string {
    return "aliexpress.interactive.task.delivery.query"
}

func (r AliexpressInteractiveTaskDeliveryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressInteractiveTaskDeliveryQueryRequest) SetRequestDto(requestDto *QueryDeliveryRequestDto) error {
    r.requestDto = requestDto
    r.Set("request_dto", requestDto)
    return nil
}

func (r AliexpressInteractiveTaskDeliveryQueryRequest) GetRequestDto() *QueryDeliveryRequestDto {
    return r.requestDto
}

